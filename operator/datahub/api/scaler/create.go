package scaler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"
	autoscalingv1alpha2 "github.com/containers-ai/alameda/operator/api/v1alpha2"
	operatorutils "github.com/containers-ai/alameda/operator/pkg/utils"
	alamedaconsts "github.com/containers-ai/alameda/pkg/consts"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	openshiftappsv1 "github.com/openshift/api/apps/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func CreateV1Alpha2Scaler(
	datahubClient *datahubpkg.Client, k8sClient client.Client,
	kafkaClient kafka.Client, scaler *autoscalingv1alpha2.AlamedaScaler,
	enabledDA, isOpenshift bool) error {

	err := datahubClient.Create(&[]entities.ResourceClusterStatusApplication{
		{
			ClusterName: scaler.Spec.ClusterName,
			Namespace:   scaler.Namespace,
			Name:        scaler.Name,
		},
	})
	if err != nil {
		return err
	}

	err = datahubClient.Create(&[]entities.TargetClusterStatusCluster{
		{
			Name:                   scaler.Spec.ClusterName,
			AlamedaScalerNamespace: scaler.Namespace,
			AlamedaScalerName:      scaler.Name,
			RawSpec:                scaler.Spec.ClusterName,
		},
	})
	if err != nil {
		return err
	}

	targetCtls := []entities.TargetClusterStatusController{}
	targetKafkaCgs := []entities.TargetKafkaConsumerGroup{}
	for _, controller := range scaler.Spec.Controllers {
		enableExecution := false
		if controller.EnableExecution != nil {
			enableExecution = *controller.EnableExecution
		}
		if controller.Type == autoscalingv1alpha2.GenericTarget && controller.Generic != nil {
			genericRawSpec, _ := json.Marshal(controller)
			targetCtlEntity := entities.TargetClusterStatusController{
				Name:                     controller.Generic.Target.Name,
				Namespace:                controller.Generic.Target.Namespace,
				ClusterName:              scaler.Spec.ClusterName,
				Kind:                     autoscalingv1alpha2.ControllerKindMap[controller.Generic.Target.Kind],
				EnableExecution:          enableExecution,
				AlamedaScalerName:        scaler.Name,
				AlamedaScalerNamespace:   scaler.Namespace,
				AlamedaScalerScalingTool: autoscalingv1alpha2.ScalingTypeMap[controller.Scaling],
				RawSpec:                  string(genericRawSpec),
			}

			if controller.Scaling == autoscalingv1alpha2.HPAScaling && controller.Generic.HpaParameters != nil {
				if controller.Generic.HpaParameters.MinReplicas != nil {
					targetCtlEntity.MinReplicas = *controller.Generic.HpaParameters.MinReplicas
				} else {
					targetCtlEntity.MinReplicas = 1
				}
				targetCtlEntity.MaxReplicas = controller.Generic.HpaParameters.MaxReplicas
			}
			targetCtls = append(targetCtls, targetCtlEntity)
		}

		if controller.Type == autoscalingv1alpha2.KafkaTarget && controller.Kafka != nil {
			cgRawSpec, _ := json.Marshal(controller)
			targetCgEntity := entities.TargetKafkaConsumerGroup{
				ClusterName:              scaler.Spec.ClusterName,
				Name:                     controller.Kafka.ConsumerGroup.Name,
				ResourceK8sName:          controller.Kafka.ConsumerGroup.Name,
				ResourceK8sNamespace:     controller.Kafka.ConsumerGroup.Namespace,
				ResourceK8sKind:          autoscalingv1alpha2.ControllerKindMap[controller.Kafka.ConsumerGroup.Kind],
				AlamedaScalerName:        scaler.Name,
				AlamedaScalerNamespace:   scaler.Namespace,
				ExporterNamespace:        controller.Kafka.ExporterNamespce,
				TopicName:                controller.Kafka.ConsumerGroup.Topic,
				AlamedaScalerScalingTool: autoscalingv1alpha2.ScalingTypeMap[controller.Scaling],
				EnableExecution:          enableExecution,
				RawSpec:                  string(cgRawSpec),
			}
			if controller.Kafka.ConsumerGroup.GroupId != nil {
				targetCgEntity.GroupId = *controller.Kafka.ConsumerGroup.GroupId
			}
			if controller.Scaling == autoscalingv1alpha2.HPAScaling && controller.Kafka.HpaParameters != nil {
				if controller.Kafka.HpaParameters.MinReplicas != nil {
					targetCgEntity.ResourceK8sMinReplicas = *controller.Kafka.HpaParameters.MinReplicas
				} else {
					targetCgEntity.ResourceK8sMinReplicas = 1
				}
				targetCgEntity.ResourceK8sMaxReplicas = controller.Kafka.HpaParameters.MaxReplicas
			}
			targetKafkaCgs = append(targetKafkaCgs, targetCgEntity)
			targetCtlEntity := entities.TargetClusterStatusController{
				Name:                     controller.Kafka.ConsumerGroup.Name,
				Namespace:                controller.Kafka.ConsumerGroup.Namespace,
				ClusterName:              scaler.Spec.ClusterName,
				Kind:                     autoscalingv1alpha2.ControllerKindMap[controller.Kafka.ConsumerGroup.Kind],
				EnableExecution:          false,
				AlamedaScalerName:        scaler.Name,
				AlamedaScalerNamespace:   scaler.Namespace,
				AlamedaScalerScalingTool: autoscalingv1alpha2.ScalingTypeMap[autoscalingv1alpha2.NonScaling],
				RawSpec:                  string(cgRawSpec),
			}
			targetCtls = append(targetCtls, targetCtlEntity)
		}
	}

	err = datahubClient.Create(&targetCtls)
	if err != nil {
		return err
	}

	err = datahubClient.Create(&targetKafkaCgs)
	if err != nil {
		return err
	}

	// write final meta info if no data adapter available
	resourceContainers := []entities.ResourceClusterStatusContainer{}
	resourceControllers := []entities.ResourceClusterStatusController{}
	resourcePods := []entities.ResourceClusterStatusPod{}
	appKafkaTopics := []entities.ApplicationKafkaTopic{}
	appKafkaCgs := []entities.ApplicationKafkaConsumerGroup{}
	ctx := context.TODO()
	if !enabledDA {
		if uid, err := k8sutils.GetClusterUID(k8sClient); err != nil {
			return err
		} else if uid == "" {
			return errors.New("get empty cluster uid")
		} else if uid != scaler.Spec.ClusterName {
			return fmt.Errorf("local cluster id %s is not matched with scaler defined cluster name %s",
				uid, scaler.Spec.ClusterName)
		}

		if openErr := kafkaClient.Open(); openErr != nil {
			return openErr
		}
		for _, targetKafkaCg := range targetKafkaCgs {
			consumedTopics, err := kafkaClient.ListConsumeTopics(ctx, targetKafkaCg.GroupId)
			if err != nil {
				return err
			}
			topicMatched := false
			for _, consumedTopic := range consumedTopics {
				if consumedTopic == targetKafkaCg.TopicName {
					topicMatched = true
				}
			}
			if !topicMatched {
				continue
			}

			if targetKafkaCg.ResourceK8sKind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.DeploymentController] {
				deployIns := appsv1.Deployment{}
				err := k8sClient.Get(ctx, types.NamespacedName{
					Namespace: targetKafkaCg.ResourceK8sNamespace,
					Name:      targetKafkaCg.ResourceK8sName,
				}, &deployIns)
				if err != nil && k8serrors.IsNotFound(err) {
					continue
				} else if err != nil {
					return err
				}

				podInsList := corev1.PodList{}
				labelMap, err := metav1.LabelSelectorAsMap(deployIns.Spec.Selector)
				if err != nil {
					return err
				}
				err = k8sClient.List(ctx, &podInsList, &client.ListOptions{
					Namespace:     deployIns.Namespace,
					LabelSelector: labels.SelectorFromSet(labelMap),
				})
				if err != nil {
					return err
				}

				appCgEntity := entities.ApplicationKafkaConsumerGroup{
					Name:                     targetKafkaCg.GroupId,
					Namespace:                targetKafkaCg.ExporterNamespace,
					ClusterName:              targetKafkaCg.ClusterName,
					TopicName:                targetKafkaCg.TopicName,
					AlamedaScalerName:        targetKafkaCg.AlamedaScalerName,
					AlamedaScalerNamespace:   targetKafkaCg.AlamedaScalerNamespace,
					AlamedaScalerScalingTool: targetKafkaCg.AlamedaScalerScalingTool,
					ResourceK8sName:          targetKafkaCg.ResourceK8sName,
					ResourceK8sNamespace:     targetKafkaCg.ResourceK8sNamespace,
					ResourceK8sKind:          targetKafkaCg.ResourceK8sKind,
					ResourceK8sSpecReplicas:  *deployIns.Spec.Replicas,
					ResourceK8sReplicas:      deployIns.Status.AvailableReplicas,
					ResourceK8sMinReplicas:   targetKafkaCg.ResourceK8sMinReplicas,
					ResourceK8sMaxReplicas:   targetKafkaCg.ResourceK8sMaxReplicas,
					Policy:                   targetKafkaCg.Policy,
					EnableExecution:          targetKafkaCg.EnableExecution,
				}
				resource := operatorutils.GetTotalResourceFromContainers(
					deployIns.Spec.Template.Spec.Containers)
				appCgEntity.ResourceCPULimit =
					strconv.FormatInt(resource.Limits.Cpu().MilliValue(), 10)
				appCgEntity.ResourceCPURequest =
					strconv.FormatInt(resource.Requests.Cpu().MilliValue(), 10)
				appCgEntity.ResourceMemoryLimit =
					strconv.FormatInt(resource.Limits.Memory().Value(), 10)
				appCgEntity.ResourceMemoryRequest =
					strconv.FormatInt(resource.Requests.Memory().Value(), 10)
				volumeCapacity := operatorutils.VolumeCapacity{}
				for _, pod := range podInsList.Items {
					v, err := operatorutils.GetVolumeCapacityUsedByPod(ctx, k8sClient, pod)
					if err != nil {
						return err
					}
					volumeCapacity.Add(v)
				}
				appCgEntity.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
				appCgEntity.VolumesPvcSize = strconv.FormatInt(volumeCapacity.PVC, 10)
				appKafkaCgs = append(appKafkaCgs, appCgEntity)
				appKafkaTopics = append(appKafkaTopics, entities.ApplicationKafkaTopic{
					Name:                   appCgEntity.TopicName,
					Namespace:              appCgEntity.Namespace,
					ClusterName:            appCgEntity.ClusterName,
					AlamedaScalerName:      appCgEntity.AlamedaScalerName,
					AlamedaScalerNamespace: appCgEntity.AlamedaScalerNamespace,
				})
			}
			if targetKafkaCg.ResourceK8sKind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.StatefulSetController] {
				stsIns := appsv1.StatefulSet{}
				err := k8sClient.Get(ctx, types.NamespacedName{
					Namespace: targetKafkaCg.ResourceK8sNamespace, Name: targetKafkaCg.ResourceK8sName,
				}, &stsIns)
				if err != nil && k8serrors.IsNotFound(err) {
					continue
				} else if err != nil {
					return err
				}
				podInsList := corev1.PodList{}
				labelMap, err := metav1.LabelSelectorAsMap(stsIns.Spec.Selector)
				if err != nil {
					return err
				}
				err = k8sClient.List(ctx, &podInsList, &client.ListOptions{
					Namespace:     stsIns.Namespace,
					LabelSelector: labels.SelectorFromSet(labelMap),
				})
				if err != nil {
					return err
				}

				appCgEntity := entities.ApplicationKafkaConsumerGroup{
					Name:                     targetKafkaCg.GroupId,
					Namespace:                targetKafkaCg.ExporterNamespace,
					ClusterName:              targetKafkaCg.ClusterName,
					TopicName:                targetKafkaCg.TopicName,
					AlamedaScalerName:        targetKafkaCg.AlamedaScalerName,
					AlamedaScalerNamespace:   targetKafkaCg.AlamedaScalerNamespace,
					AlamedaScalerScalingTool: targetKafkaCg.AlamedaScalerScalingTool,
					ResourceK8sName:          targetKafkaCg.ResourceK8sName,
					ResourceK8sNamespace:     targetKafkaCg.ResourceK8sNamespace,
					ResourceK8sKind:          targetKafkaCg.ResourceK8sKind,
					ResourceK8sSpecReplicas:  *stsIns.Spec.Replicas,
					ResourceK8sReplicas:      stsIns.Status.CurrentReplicas,
					ResourceK8sMinReplicas:   targetKafkaCg.ResourceK8sMinReplicas,
					ResourceK8sMaxReplicas:   targetKafkaCg.ResourceK8sMaxReplicas,
					Policy:                   targetKafkaCg.Policy,
					EnableExecution:          targetKafkaCg.EnableExecution,
				}
				resource := operatorutils.GetTotalResourceFromContainers(
					stsIns.Spec.Template.Spec.Containers)
				appCgEntity.ResourceCPULimit =
					strconv.FormatInt(resource.Limits.Cpu().MilliValue(), 10)
				appCgEntity.ResourceCPURequest =
					strconv.FormatInt(resource.Requests.Cpu().MilliValue(), 10)
				appCgEntity.ResourceMemoryLimit =
					strconv.FormatInt(resource.Limits.Memory().Value(), 10)
				appCgEntity.ResourceMemoryRequest =
					strconv.FormatInt(resource.Requests.Memory().Value(), 10)
				volumeCapacity := operatorutils.VolumeCapacity{}
				for _, pod := range podInsList.Items {
					v, err := operatorutils.GetVolumeCapacityUsedByPod(ctx, k8sClient, pod)
					if err != nil {
						return err
					}
					volumeCapacity.Add(v)
				}
				appCgEntity.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
				appCgEntity.VolumesPvcSize = strconv.FormatInt(volumeCapacity.PVC, 10)
				appKafkaCgs = append(appKafkaCgs, appCgEntity)
				appKafkaTopics = append(appKafkaTopics, entities.ApplicationKafkaTopic{
					Name:                   appCgEntity.TopicName,
					Namespace:              appCgEntity.Namespace,
					ClusterName:            appCgEntity.ClusterName,
					AlamedaScalerName:      appCgEntity.AlamedaScalerName,
					AlamedaScalerNamespace: appCgEntity.AlamedaScalerNamespace,
				})
			}
			if isOpenshift && targetKafkaCg.ResourceK8sKind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.DeploymentConfigController] {
				dcIns := openshiftappsv1.DeploymentConfig{}
				err := k8sClient.Get(ctx, types.NamespacedName{
					Namespace: targetKafkaCg.ResourceK8sNamespace, Name: targetKafkaCg.ResourceK8sName,
				}, &dcIns)
				if err != nil && k8serrors.IsNotFound(err) {
					continue
				} else if err != nil {
					return err
				}
				podInsList := corev1.PodList{}
				err = k8sClient.List(ctx, &podInsList, &client.ListOptions{
					Namespace:     dcIns.Namespace,
					LabelSelector: labels.SelectorFromSet(labels.Set(dcIns.Spec.Selector)),
				})
				if err != nil {
					return err
				}
				appCgEntity := entities.ApplicationKafkaConsumerGroup{
					Name:                     targetKafkaCg.GroupId,
					Namespace:                targetKafkaCg.ExporterNamespace,
					ClusterName:              targetKafkaCg.ClusterName,
					TopicName:                targetKafkaCg.TopicName,
					AlamedaScalerName:        targetKafkaCg.AlamedaScalerName,
					AlamedaScalerNamespace:   targetKafkaCg.AlamedaScalerNamespace,
					AlamedaScalerScalingTool: targetKafkaCg.AlamedaScalerScalingTool,
					ResourceK8sName:          targetKafkaCg.ResourceK8sName,
					ResourceK8sNamespace:     targetKafkaCg.ResourceK8sNamespace,
					ResourceK8sKind:          targetKafkaCg.ResourceK8sKind,
					ResourceK8sSpecReplicas:  dcIns.Spec.Replicas,
					ResourceK8sReplicas:      dcIns.Status.AvailableReplicas,
					ResourceK8sMinReplicas:   targetKafkaCg.ResourceK8sMinReplicas,
					ResourceK8sMaxReplicas:   targetKafkaCg.ResourceK8sMaxReplicas,
					Policy:                   targetKafkaCg.Policy,
					EnableExecution:          targetKafkaCg.EnableExecution,
				}
				resource := operatorutils.GetTotalResourceFromContainers(
					dcIns.Spec.Template.Spec.Containers)
				appCgEntity.ResourceCPULimit =
					strconv.FormatInt(resource.Limits.Cpu().MilliValue(), 10)
				appCgEntity.ResourceCPURequest =
					strconv.FormatInt(resource.Requests.Cpu().MilliValue(), 10)
				appCgEntity.ResourceMemoryLimit =
					strconv.FormatInt(resource.Limits.Memory().Value(), 10)
				appCgEntity.ResourceMemoryRequest =
					strconv.FormatInt(resource.Requests.Memory().Value(), 10)
				volumeCapacity := operatorutils.VolumeCapacity{}
				for _, pod := range podInsList.Items {
					v, err := operatorutils.GetVolumeCapacityUsedByPod(ctx, k8sClient, pod)
					if err != nil {
						return err
					}
					volumeCapacity.Add(v)
				}
				appCgEntity.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
				appCgEntity.VolumesPvcSize = strconv.FormatInt(volumeCapacity.PVC, 10)
				appKafkaCgs = append(appKafkaCgs, appCgEntity)
				appKafkaTopics = append(appKafkaTopics, entities.ApplicationKafkaTopic{
					Name:                   appCgEntity.TopicName,
					Namespace:              appCgEntity.Namespace,
					ClusterName:            appCgEntity.ClusterName,
					AlamedaScalerName:      appCgEntity.AlamedaScalerName,
					AlamedaScalerNamespace: appCgEntity.AlamedaScalerNamespace,
				})
			}
		}

		for _, targetCtl := range targetCtls {
			enableExecution := targetCtl.EnableExecution
			resourceK8sMinReplicas := targetCtl.MinReplicas
			resourceK8sMaxReplicas := targetCtl.MaxReplicas
			resourceCtlEntity := entities.ResourceClusterStatusController{
				Name:                     targetCtl.Name,
				Namespace:                targetCtl.Namespace,
				ClusterName:              targetCtl.ClusterName,
				Kind:                     targetCtl.Kind,
				AlamedaScalerName:        targetCtl.AlamedaScalerName,
				AlamedaScalerNamespace:   targetCtl.AlamedaScalerNamespace,
				AlamedaScalerScalingTool: targetCtl.AlamedaScalerScalingTool,
				ResourceK8sMinReplicas:   resourceK8sMinReplicas,
				ResourceK8sMaxReplicas:   resourceK8sMaxReplicas,
				EnableExecution:          enableExecution,
			}
			if targetCtl.Kind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.DeploymentController] {
				deployIns := appsv1.Deployment{}
				err := k8sClient.Get(ctx, types.NamespacedName{Namespace: targetCtl.Namespace, Name: targetCtl.Name}, &deployIns)
				if err != nil && k8serrors.IsNotFound(err) {
					continue
				} else if err != nil {
					return err
				}
				resourceCtlEntity.SpecReplicas = *deployIns.Spec.Replicas
				resourceCtlEntity.Replicas = deployIns.Status.AvailableReplicas
			}
			if targetCtl.Kind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.StatefulSetController] {
				stsIns := appsv1.StatefulSet{}
				err := k8sClient.Get(ctx, types.NamespacedName{Namespace: targetCtl.Namespace, Name: targetCtl.Name}, &stsIns)
				if err != nil && k8serrors.IsNotFound(err) {
					continue
				} else if err != nil {
					return err
				}
				resourceCtlEntity.SpecReplicas = *stsIns.Spec.Replicas
				resourceCtlEntity.Replicas = stsIns.Status.CurrentReplicas
			}
			if isOpenshift && targetCtl.Kind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.DeploymentConfigController] {
				dcIns := openshiftappsv1.DeploymentConfig{}
				err := k8sClient.Get(ctx, types.NamespacedName{Namespace: targetCtl.Namespace, Name: targetCtl.Name}, &dcIns)
				if err != nil && k8serrors.IsNotFound(err) {
					continue
				} else if err != nil {
					return err
				}
				resourceCtlEntity.SpecReplicas = dcIns.Spec.Replicas
				resourceCtlEntity.Replicas = dcIns.Status.AvailableReplicas
			}
			resourceControllers = append(resourceControllers, resourceCtlEntity)
		}
	}

	finalResourceControllers := []entities.ResourceClusterStatusController{}
	for _, resourceController := range resourceControllers {
		isValidCtl := true
		for _, targetKafkaCg := range targetKafkaCgs {
			if targetKafkaCg.ClusterName == resourceController.ClusterName &&
				targetKafkaCg.ResourceK8sNamespace == resourceController.Namespace &&
				targetKafkaCg.ResourceK8sName == resourceController.Name {
				isValidCtl = false
				for _, appKafkaCg := range appKafkaCgs {
					if targetKafkaCg.ClusterName == appKafkaCg.ClusterName &&
						targetKafkaCg.ResourceK8sNamespace == appKafkaCg.ResourceK8sNamespace &&
						targetKafkaCg.ResourceK8sName == appKafkaCg.ResourceK8sName {
						isValidCtl = true
						break
					}
				}
				if isValidCtl {
					break
				}
			}
		}
		if isValidCtl {
			finalResourceControllers = append(finalResourceControllers, resourceController)
		}
	}

	for _, finalResourceController := range finalResourceControllers {
		if finalResourceController.Kind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.DeploymentController] {
			deployIns := appsv1.Deployment{}
			err := k8sClient.Get(ctx, types.NamespacedName{Namespace: finalResourceController.Namespace, Name: finalResourceController.Name}, &deployIns)
			if err != nil && k8serrors.IsNotFound(err) {
				continue
			} else if err != nil {
				return err
			}

			podInsList := corev1.PodList{}
			labelMap, err := metav1.LabelSelectorAsMap(deployIns.Spec.Selector)
			if err != nil {
				return err
			}
			err = k8sClient.List(ctx, &podInsList, &client.ListOptions{
				Namespace:     deployIns.Namespace,
				LabelSelector: labels.SelectorFromSet(labelMap),
			})
			if err != nil {
				return err
			}
			for _, pod := range podInsList.Items {
				podNamespace := pod.GetNamespace()
				podName := pod.GetName()
				nodeName := pod.Spec.NodeName
				deployName := deployIns.GetName()
				podNamePattern := fmt.Sprintf(alamedaconsts.DeploymentPodFormat, deployName)
				regExp := regexp.MustCompile(podNamePattern)
				res := regExp.FindAllStringSubmatch(pod.GetName(), -1)
				if len(res) > 0 {
					appName := fmt.Sprintf("%s-%s", scaler.Namespace, scaler.Name)
					if _, exist := scaler.Labels["app.federator.ai/name"]; exist {
						appName = scaler.Labels["app.federator.ai/name"]
					}
					appPartOf := appName
					if _, exist := scaler.Labels["app.federator.ai/part-of"]; exist {
						appPartOf = scaler.Labels["app.federator.ai/part-of"]
					}

					resourcePods = append(resourcePods, entities.ResourceClusterStatusPod{
						Name:                     podName,
						Namespace:                podNamespace,
						NodeName:                 nodeName,
						ClusterName:              finalResourceController.ClusterName,
						TopControllerName:        finalResourceController.Name,
						TopControllerKind:        finalResourceController.Kind,
						AlamedaScalerName:        finalResourceController.AlamedaScalerName,
						AlamedaScalerNamespace:   finalResourceController.AlamedaScalerNamespace,
						AlamedaScalerScalingTool: finalResourceController.AlamedaScalerScalingTool,
						AppName:                  appName,
						AppPartOf:                appPartOf,
						PodCreateTime:            pod.GetCreationTimestamp().Time.Unix(),
						TopControllerReplicas:    finalResourceController.Replicas,
					})
					for _, container := range pod.Spec.Containers {
						resourceContainerEntity := entities.ResourceClusterStatusContainer{
							Name:                     container.Name,
							Namespace:                podNamespace,
							NodeName:                 nodeName,
							ClusterName:              finalResourceController.ClusterName,
							PodName:                  podName,
							TopControllerName:        finalResourceController.Name,
							TopControllerKind:        finalResourceController.Kind,
							AlamedaScalerName:        finalResourceController.AlamedaScalerName,
							AlamedaScalerNamespace:   finalResourceController.AlamedaScalerNamespace,
							AlamedaScalerScalingTool: finalResourceController.AlamedaScalerScalingTool,
						}
						if container.Resources.Limits.Cpu() != nil {
							resourceContainerEntity.ResourceLimitCpu =
								strconv.FormatInt(container.Resources.Limits.Cpu().MilliValue(), 10)
						}
						if container.Resources.Limits.Memory() != nil {
							resourceContainerEntity.ResourceLimitMemory =
								strconv.FormatInt(container.Resources.Limits.Memory().Value(), 10)
						}
						if container.Resources.Requests.Cpu() != nil {
							resourceContainerEntity.ResourceRequestCPU =
								strconv.FormatInt(container.Resources.Requests.Cpu().MilliValue(), 10)
						}
						if container.Resources.Requests.Memory() != nil {
							resourceContainerEntity.ResourceRequestMemory =
								strconv.FormatInt(container.Resources.Requests.Memory().Value(), 10)
						}
						resourceContainers = append(resourceContainers, resourceContainerEntity)
					}
				}
			}
		}
		if finalResourceController.Kind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.StatefulSetController] {
			stsIns := appsv1.StatefulSet{}
			err := k8sClient.Get(ctx, types.NamespacedName{Namespace: finalResourceController.Namespace, Name: finalResourceController.Name}, &stsIns)
			if err != nil && k8serrors.IsNotFound(err) {
				continue
			} else if err != nil {
				return err
			}

			podInsList := corev1.PodList{}
			labelMap, err := metav1.LabelSelectorAsMap(stsIns.Spec.Selector)
			if err != nil {
				return err
			}
			err = k8sClient.List(ctx, &podInsList, &client.ListOptions{
				Namespace:     stsIns.Namespace,
				LabelSelector: labels.SelectorFromSet(labelMap),
			})
			if err != nil {
				return err
			}
			for _, pod := range podInsList.Items {
				podNamespace := pod.GetNamespace()
				podName := pod.GetName()
				nodeName := pod.Spec.NodeName
				stsName := stsIns.GetName()
				podNamePattern := fmt.Sprintf(alamedaconsts.StatefulSetPodFormat, stsName)
				regExp := regexp.MustCompile(podNamePattern)
				res := regExp.FindAllStringSubmatch(pod.GetName(), -1)
				if len(res) > 0 {
					appName := fmt.Sprintf("%s-%s", scaler.Namespace, scaler.Name)
					if _, exist := scaler.Labels["app.federator.ai/name"]; exist {
						appName = scaler.Labels["app.federator.ai/name"]
					}
					appPartOf := appName
					if _, exist := scaler.Labels["app.federator.ai/part-of"]; exist {
						appPartOf = scaler.Labels["app.federator.ai/part-of"]
					}

					resourcePods = append(resourcePods, entities.ResourceClusterStatusPod{
						Name:                     podName,
						Namespace:                podNamespace,
						NodeName:                 nodeName,
						ClusterName:              finalResourceController.ClusterName,
						TopControllerName:        finalResourceController.Name,
						TopControllerKind:        finalResourceController.Kind,
						AlamedaScalerName:        finalResourceController.AlamedaScalerName,
						AlamedaScalerScalingTool: finalResourceController.AlamedaScalerScalingTool,
						AppName:                  appName,
						AppPartOf:                appPartOf,
						PodCreateTime:            pod.GetCreationTimestamp().Time.Unix(),
						TopControllerReplicas:    finalResourceController.Replicas,
					})
					for _, container := range pod.Spec.Containers {
						resourceContainerEntity := entities.ResourceClusterStatusContainer{
							Name:                     container.Name,
							Namespace:                podNamespace,
							NodeName:                 nodeName,
							ClusterName:              finalResourceController.ClusterName,
							PodName:                  podName,
							TopControllerName:        finalResourceController.Name,
							TopControllerKind:        finalResourceController.Kind,
							AlamedaScalerName:        finalResourceController.AlamedaScalerName,
							AlamedaScalerScalingTool: finalResourceController.AlamedaScalerScalingTool,
						}
						if container.Resources.Limits.Cpu() != nil {
							resourceContainerEntity.ResourceLimitCpu =
								strconv.FormatInt(container.Resources.Limits.Cpu().MilliValue(), 10)
						}
						if container.Resources.Limits.Memory() != nil {
							resourceContainerEntity.ResourceLimitMemory =
								strconv.FormatInt(container.Resources.Limits.Memory().Value(), 10)
						}
						if container.Resources.Requests.Cpu() != nil {
							resourceContainerEntity.ResourceRequestCPU =
								strconv.FormatInt(container.Resources.Requests.Cpu().MilliValue(), 10)
						}
						if container.Resources.Requests.Memory() != nil {
							resourceContainerEntity.ResourceRequestMemory =
								strconv.FormatInt(container.Resources.Requests.Memory().Value(), 10)
						}
						resourceContainers = append(resourceContainers, resourceContainerEntity)
					}
				}
			}

		}
		if isOpenshift && finalResourceController.Kind == autoscalingv1alpha2.ControllerKindMap[autoscalingv1alpha2.DeploymentConfigController] {
			dcIns := openshiftappsv1.DeploymentConfig{}
			err := k8sClient.Get(ctx, types.NamespacedName{Namespace: finalResourceController.Namespace, Name: finalResourceController.Name}, &dcIns)
			if err != nil && k8serrors.IsNotFound(err) {
				continue
			} else if err != nil {
				return err
			}

			podInsList := corev1.PodList{}
			err = k8sClient.List(ctx, &podInsList, &client.ListOptions{
				Namespace:     dcIns.Namespace,
				LabelSelector: labels.SelectorFromSet(labels.Set(dcIns.Spec.Selector)),
			})
			if err != nil {
				return err
			}
			for _, pod := range podInsList.Items {
				podNamespace := pod.GetNamespace()
				podName := pod.GetName()
				nodeName := pod.Spec.NodeName
				dcName := dcIns.GetName()
				podNamePattern := fmt.Sprintf(alamedaconsts.DeploymentConfigPodFormat, dcName)
				regExp := regexp.MustCompile(podNamePattern)
				res := regExp.FindAllStringSubmatch(pod.GetName(), -1)
				if len(res) > 0 {
					appName := fmt.Sprintf("%s-%s", scaler.Namespace, scaler.Name)
					if _, exist := scaler.Labels["app.federator.ai/name"]; exist {
						appName = scaler.Labels["app.federator.ai/name"]
					}
					appPartOf := appName
					if _, exist := scaler.Labels["app.federator.ai/part-of"]; exist {
						appPartOf = scaler.Labels["app.federator.ai/part-of"]
					}

					resourcePods = append(resourcePods, entities.ResourceClusterStatusPod{
						Name:                     podName,
						Namespace:                podNamespace,
						NodeName:                 nodeName,
						ClusterName:              finalResourceController.ClusterName,
						TopControllerName:        finalResourceController.Name,
						TopControllerKind:        finalResourceController.Kind,
						AlamedaScalerName:        finalResourceController.AlamedaScalerName,
						AlamedaScalerScalingTool: finalResourceController.AlamedaScalerScalingTool,
						AppName:                  appName,
						AppPartOf:                appPartOf,
						PodCreateTime:            pod.GetCreationTimestamp().Time.Unix(),
						TopControllerReplicas:    finalResourceController.Replicas,
					})
					for _, container := range pod.Spec.Containers {
						resourceContainerEntity := entities.ResourceClusterStatusContainer{
							Name:                     container.Name,
							Namespace:                podNamespace,
							NodeName:                 nodeName,
							ClusterName:              finalResourceController.ClusterName,
							PodName:                  podName,
							TopControllerName:        finalResourceController.Name,
							TopControllerKind:        finalResourceController.Kind,
							AlamedaScalerName:        finalResourceController.AlamedaScalerName,
							AlamedaScalerScalingTool: finalResourceController.AlamedaScalerScalingTool,
						}
						if container.Resources.Limits.Cpu() != nil {
							resourceContainerEntity.ResourceLimitCpu =
								strconv.FormatInt(container.Resources.Limits.Cpu().MilliValue(), 10)
						}
						if container.Resources.Limits.Memory() != nil {
							resourceContainerEntity.ResourceLimitMemory =
								strconv.FormatInt(container.Resources.Limits.Memory().Value(), 10)
						}
						if container.Resources.Requests.Cpu() != nil {
							resourceContainerEntity.ResourceRequestCPU =
								strconv.FormatInt(container.Resources.Requests.Cpu().MilliValue(), 10)
						}
						if container.Resources.Requests.Memory() != nil {
							resourceContainerEntity.ResourceRequestMemory =
								strconv.FormatInt(container.Resources.Requests.Memory().Value(), 10)
						}
						resourceContainers = append(resourceContainers, resourceContainerEntity)
					}
				}
			}
		}
	}

	err = datahubClient.Create(&resourceContainers)
	if err != nil {
		return err
	}

	err = datahubClient.Create(&finalResourceControllers)
	if err != nil {
		return err
	}

	err = datahubClient.Create(&resourcePods)
	if err != nil {
		return err
	}

	err = datahubClient.Create(&appKafkaTopics)
	if err != nil {
		return err
	}

	err = datahubClient.Create(&appKafkaCgs)
	if err != nil {
		return err
	}
	return nil
}
