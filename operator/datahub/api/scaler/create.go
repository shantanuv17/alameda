package scaler

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	autoscalingv1alpha2 "github.com/containers-ai/alameda/operator/api/v1alpha2"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
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
	datahubClient *datahubpkg.Client,
	k8sClient client.Client, scaler *autoscalingv1alpha2.AlamedaScaler,
	enabledDA, isOpenshift bool) error {
	// finally delete application
	err := datahubClient.Create(&[]entities.ResourceClusterStatusApplication{
		{
			ClusterName: scaler.ClusterName,
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
				}
				if controller.Generic.HpaParameters.MaxReplicas != nil {
					targetCtlEntity.MaxReplicas = *controller.Generic.HpaParameters.MaxReplicas
				}
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
				}
				if controller.Kafka.HpaParameters.MaxReplicas != nil {
					targetCgEntity.ResourceK8sMaxReplicas = *controller.Kafka.HpaParameters.MaxReplicas
				}
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
					podNamePattern := fmt.Sprintf(`%s-([a-z0-9]+)-([a-z0-9]+)`, deployName)
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
							ClusterName:              targetCtl.ClusterName,
							TopControllerName:        targetCtl.Name,
							TopControllerKind:        targetCtl.Kind,
							AlamedaScalerName:        targetCtl.AlamedaScalerName,
							AlamedaScalerNamespace:   targetCtl.AlamedaScalerNamespace,
							AlamedaScalerScalingTool: targetCtl.AlamedaScalerScalingTool,
							AppName:                  appName,
							AppPartOf:                appPartOf,
							PodCreateTime:            pod.GetCreationTimestamp().Time.Unix(),
							TopControllerReplicas:    resourceCtlEntity.Replicas,
						})
						for _, container := range pod.Spec.Containers {
							resourceContainerEntity := entities.ResourceClusterStatusContainer{
								Name:                     container.Name,
								Namespace:                podNamespace,
								NodeName:                 nodeName,
								ClusterName:              targetCtl.ClusterName,
								PodName:                  podName,
								TopControllerName:        targetCtl.Name,
								TopControllerKind:        targetCtl.Kind,
								AlamedaScalerName:        targetCtl.AlamedaScalerName,
								AlamedaScalerNamespace:   targetCtl.AlamedaScalerNamespace,
								AlamedaScalerScalingTool: targetCtl.AlamedaScalerScalingTool,
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
					podNamePattern := fmt.Sprintf(`%s-([0-9]+)`, stsName)
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
							ClusterName:              targetCtl.ClusterName,
							TopControllerName:        targetCtl.Name,
							TopControllerKind:        targetCtl.Kind,
							AlamedaScalerName:        targetCtl.AlamedaScalerName,
							AlamedaScalerScalingTool: targetCtl.AlamedaScalerScalingTool,
							AppName:                  appName,
							AppPartOf:                appPartOf,
							PodCreateTime:            pod.GetCreationTimestamp().Time.Unix(),
							TopControllerReplicas:    resourceCtlEntity.Replicas,
						})
						for _, container := range pod.Spec.Containers {
							resourceContainerEntity := entities.ResourceClusterStatusContainer{
								Name:                     container.Name,
								Namespace:                podNamespace,
								NodeName:                 nodeName,
								ClusterName:              targetCtl.ClusterName,
								PodName:                  podName,
								TopControllerName:        targetCtl.Name,
								TopControllerKind:        targetCtl.Kind,
								AlamedaScalerName:        targetCtl.AlamedaScalerName,
								AlamedaScalerScalingTool: targetCtl.AlamedaScalerScalingTool,
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
					podNamePattern := fmt.Sprintf(`%s-([0-9]+)-([a-z0-9]+)`, dcName)
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
							ClusterName:              targetCtl.ClusterName,
							TopControllerName:        targetCtl.Name,
							TopControllerKind:        targetCtl.Kind,
							AlamedaScalerName:        targetCtl.AlamedaScalerName,
							AlamedaScalerScalingTool: targetCtl.AlamedaScalerScalingTool,
							AppName:                  appName,
							AppPartOf:                appPartOf,
							PodCreateTime:            pod.GetCreationTimestamp().Time.Unix(),
							TopControllerReplicas:    resourceCtlEntity.Replicas,
						})
						for _, container := range pod.Spec.Containers {
							resourceContainerEntity := entities.ResourceClusterStatusContainer{
								Name:                     container.Name,
								Namespace:                podNamespace,
								NodeName:                 nodeName,
								ClusterName:              targetCtl.ClusterName,
								PodName:                  podName,
								TopControllerName:        targetCtl.Name,
								TopControllerKind:        targetCtl.Kind,
								AlamedaScalerName:        targetCtl.AlamedaScalerName,
								AlamedaScalerScalingTool: targetCtl.AlamedaScalerScalingTool,
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
			resourceControllers = append(resourceControllers, resourceCtlEntity)
		}

		//for _, targetKafkaCg := range targetKafkaCgs {
		//}
	}
	err = datahubClient.Create(&resourceContainers)
	if err != nil {
		return err
	}
	err = datahubClient.Create(&resourceControllers)
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
