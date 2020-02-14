package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"
	kafkarepository "github.com/containers-ai/alameda/operator/datahub/client/kafka"
	kafkamodel "github.com/containers-ai/alameda/operator/pkg/kafka"
	utilsresources "github.com/containers-ai/alameda/operator/pkg/utils/resources"
	"github.com/containers-ai/alameda/pkg/utils/log"

	datahubresources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	datahubschemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"

	openshiftapiappsv1 "github.com/openshift/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// chooseTopic chooses and returns topics name
func chooseTopic(majorTopic string, topicsWantToConsume, currentConsumeTopics []string) string {
	var empty struct{}
	currentConsumeTopicMap := make(map[string]struct{})
	for _, topic := range currentConsumeTopics {
		currentConsumeTopicMap[topic] = empty
	}

	if _, isMajorTopicBeingConsumed := currentConsumeTopicMap[majorTopic]; isMajorTopicBeingConsumed {
		return majorTopic
	}

	for _, topic := range topicsWantToConsume {
		if _, isTopicBeingConsumed := currentConsumeTopicMap[topic]; isTopicBeingConsumed {
			return topic
		}
	}

	return ""
}

// AlamedaScalerKafkaReconciler reconciles AlamedaScaler with Spec.Type in "kafka".
type AlamedaScalerKafkaReconciler struct {
	ClusterUID            string
	HasOpenShiftAPIAppsv1 bool

	K8SClient client.Client
	Scheme    *runtime.Scheme

	KafkaRepository                                 kafkarepository.KafkaRepository
	DatahubApplicationKafkaTopicSchema              datahubschemas.Schema
	DatahubApplicationKafkaTopicMeasurement         datahubschemas.Measurement
	DatahubApplicationKafkaConsumerGroupSchema      datahubschemas.Schema
	DatahubApplicationKafkaConsumerGroupMeasurement datahubschemas.Measurement

	KafkaClient      kafka.Client
	PrometheusClient prometheus.Prometheus

	ReconcileTimeout time.Duration

	Logger *log.Scope

	NeededMetrics []string
}

func (r AlamedaScalerKafkaReconciler) isMetricsExist(ctx context.Context, metrics []string) (bool, error) {
	ok, nonExistMetrics, err := r.PrometheusClient.IsMetricsExist(ctx, r.NeededMetrics)
	if err != nil {
		return false, err
	} else if len(nonExistMetrics) > 0 {
		return false, errors.Errorf("metrics not eixst: metrics: %+v", metrics)
	}
	return ok, nil
}

func (r AlamedaScalerKafkaReconciler) openClient() error {
	if err := r.KafkaClient.Open(); err != nil {
		return errors.Wrap(err, "open kafka client failed")
	}
	return nil
}

func (r *AlamedaScalerKafkaReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {

	ctx, cancel := context.WithTimeout(context.Background(), r.ReconcileTimeout)
	defer cancel()
	cachedAlamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	err := r.K8SClient.Get(ctx, client.ObjectKey{Namespace: req.Namespace, Name: req.Name}, &cachedAlamedaScaler)
	if err != nil && k8serrors.IsNotFound(err) {
		r.Logger.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleDeletion(ctx, req.Namespace, req.Name); err != nil {
			r.Logger.Warnf("Handle deleteion of AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		r.Logger.Infof("Handle deletion of AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	}
	alamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	cachedAlamedaScaler.DeepCopyInto(&alamedaScaler)
	if alamedaScaler.Status.Kafka == nil {
		alamedaScaler.Status.Kafka = &autoscalingv1alpha1.KafkaStatus{}
	}

	if ok, err := r.isAlamedaScalerNeedToBeReconciled(ctx, alamedaScaler); err != nil {
		r.Logger.Infof("Check if AlamedaScaler(%s/%s) needs to be reconciled failed, retry reconciling: %s", req.Namespace, req.Name, err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	} else if !ok {
		r.Logger.Infof("AlamedaScaler(%s/%s) type(%s), skip reconciling.", req.Namespace, req.Name, alamedaScaler.GetType())
		alamedaScaler.Status.Kafka.Effective = false
		if err != nil {
			alamedaScaler.Status.Kafka.Message = err.Error()
		}
		err := r.updateAlamedaScaler(ctx, &alamedaScaler)
		if err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		return ctrl.Result{Requeue: false}, nil
	}

	if cachedAlamedaScaler.GetDeletionTimestamp() != nil {
		r.Logger.Infof("Handling deletion of AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
		if err := r.handleDeletion(ctx, req.Namespace, req.Name); err != nil {
			r.Logger.Warnf("Handle deleteion of AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		r.Logger.Infof("Handle deletion of AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
		return ctrl.Result{Requeue: false}, nil
	}

	ok, err := r.isMetricsExist(ctx, r.NeededMetrics)
	if err != nil {
		r.Logger.Warnf("Check if metrics exist in Prometheus failed: metrics: %+v: err: %+v", r.NeededMetrics, err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	if !ok {
		alamedaScaler.Status.Kafka.Effective = false
		alamedaScaler.Status.Kafka.Message = "Needed metrics not exist in Prometheus."
		err := r.updateAlamedaScaler(ctx, &alamedaScaler)
		if err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
		return ctrl.Result{Requeue: false}, nil
	}

	if err := r.openClient(); err != nil {
		r.Logger.Warnf("Open AlamedaScaler(%s/%s) clients' connection failed: err: %+v", req.Namespace, req.Name, err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	r.Logger.Infof("Reconciling AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
	consumerGroupDetails, err := r.prepareConsumerGroupDetails(ctx, alamedaScaler)
	if err != nil {
		r.Logger.Warnf("Prepare consumerGroupDetails to synchornize with remote of AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
		alamedaScaler.Status.Kafka.Effective = false
		alamedaScaler.Status.Kafka.Message = err.Error()
		err := r.updateAlamedaScaler(ctx, &alamedaScaler)
		if err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	topics := r.prepareTopics(consumerGroupDetails)

	err = r.syncWithDatahub(ctx, alamedaScaler, topics, consumerGroupDetails)
	if err != nil {
		r.Logger.Warnf("Synchornize consumerGroupDetails with remote of AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
		alamedaScaler.Status.Kafka.Effective = false
		alamedaScaler.Status.Kafka.Message = err.Error()
		err := r.updateAlamedaScaler(ctx, &alamedaScaler)
		if err != nil {
			r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	kafkaStatus := r.getKafkaStatus(alamedaScaler, consumerGroupDetails)
	alamedaScaler.Status.Kafka = &kafkaStatus
	if err := r.updateAlamedaScaler(ctx, &alamedaScaler); err != nil {
		r.Logger.Warnf("Update AlamedaScaler(%s/%s) failed, retry reconciling: %s", req.Namespace, req.Name, err)
		return ctrl.Result{Requeue: false}, nil
	}

	r.Logger.Infof("Update AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
	return ctrl.Result{Requeue: false}, nil
}

func (r AlamedaScalerKafkaReconciler) handleDeletion(ctx context.Context, namespace, name string) error {
	wg := errgroup.Group{}
	wg.Go(func() error {
		opt := kafkarepository.DeleteTopicsOption{
			ClusterName:            r.ClusterUID,
			AlamedaScalerNamespace: namespace,
			AlamedaScalerName:      name,
		}
		if err := r.KafkaRepository.DeleteTopicsByOption(ctx, opt); err != nil {
			return errors.Wrap(err, "delete topics from Datahub failed")
		}
		return nil
	})

	wg.Go(func() error {
		opt := kafkarepository.DeleteConsumerGroupsOption{
			ClusterName:            r.ClusterUID,
			AlamedaScalerNamespace: namespace,
			AlamedaScalerName:      name,
		}
		if err := r.KafkaRepository.DeleteConsumerGroupsByOption(ctx, opt); err != nil {
			return errors.Wrap(err, "delete consumerGroups from Datahub failed")
		}
		return nil
	})
	return wg.Wait()
}

func (r AlamedaScalerKafkaReconciler) updateAlamedaScaler(ctx context.Context, alamedaScaler *autoscalingv1alpha1.AlamedaScaler) error {
	if err := r.K8SClient.Update(ctx, alamedaScaler); err != nil {
		return errors.Wrap(err, "update AlamedaScaler failed")
	}
	return nil
}

func (r AlamedaScalerKafkaReconciler) isAlamedaScalerNeedToBeReconciled(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) (bool, error) {
	if alamedaScaler.GetType() != autoscalingv1alpha1.AlamedaScalerTypeKafka {
		return false, nil
	}

	//Ensure that count of the AlamedaScaler with same value of AlamedaScaler.Spec.Kafka.ExporterNamespace per namespace at most one.
	lock, err := r.getOrCreateLock(ctx, alamedaScaler)
	if err != nil {
		return false, errors.Wrap(err, "get or create lock failed")
	}
	if !r.isLockOwnBy(&lock, &alamedaScaler) {
		return false, errors.Errorf("lock is owned by other AlamedaScaler(%s)", lock.GetOwnerReferences()[0].Name)
	}

	return true, nil
}

func (r AlamedaScalerKafkaReconciler) getOrCreateLock(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) (corev1.ConfigMap, error) {
	empty := corev1.ConfigMap{}

	namespace := alamedaScaler.GetNamespace()
	name := r.getLockName(alamedaScaler)
	lock := corev1.ConfigMap{}
	if err := r.K8SClient.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, &lock); err != nil && !k8serrors.IsNotFound(err) {
		return empty, errors.Wrap(err, "get lock failed")
	} else if k8serrors.IsNotFound(err) {
		lock.Namespace = namespace
		lock.Name = name
		if err := ctrl.SetControllerReference(&alamedaScaler, &lock, r.Scheme); err != nil {
			return empty, errors.Wrap(err, "set OwnerReference to lock failed")
		}
		if err := r.K8SClient.Create(ctx, &lock); err != nil {
			return empty, errors.Wrap(err, "create lock failed")
		}
	}
	return lock, nil
}

func (r AlamedaScalerKafkaReconciler) isLockOwnBy(lock metav1.ObjectMetaAccessor, owner metav1.ObjectMetaAccessor) bool {
	ownerReferences := lock.GetObjectMeta().GetOwnerReferences()
	if len(ownerReferences) == 1 {
		return ownerReferences[0].UID == owner.GetObjectMeta().GetUID()
	}
	return false
}

func (r AlamedaScalerKafkaReconciler) getLockName(alamedaScaler autoscalingv1alpha1.AlamedaScaler) string {
	return fmt.Sprintf(`alameda-scaler-kafka-lock-%s`, alamedaScaler.Spec.Kafka.ExporterNamespace)
}

func (r AlamedaScalerKafkaReconciler) prepareConsumerGroupDetails(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) ([]kafkamodel.ConsumerGroup, error) {
	// Get map from consumerGroupName to consuming topics from Kafka client.
	consumerGroupNames := make([]string, 0, len(alamedaScaler.Spec.Kafka.ConsumerGroups))
	for _, consumerGroupSpec := range alamedaScaler.Spec.Kafka.ConsumerGroups {
		consumerGroupNames = append(consumerGroupNames, consumerGroupSpec.Name)
	}
	consumerGroupToConsumeTopicsMap, err := r.getConsumerGroupToConsumeTopicsMap(ctx, consumerGroupNames)
	if err != nil {
		return nil, errors.Wrap(err, "get consumerGroup consume topics map failed")
	}

	consumerGroupDetails, err := r.listConsumerGroups(ctx, alamedaScaler, consumerGroupToConsumeTopicsMap)
	if err != nil {
		return nil, errors.Wrap(err, "list consumerGroups' detail failed")
	}

	return consumerGroupDetails, nil
}

func (r AlamedaScalerKafkaReconciler) prepareTopics(consumerGroups []kafkamodel.ConsumerGroup) []kafkamodel.Topic {
	topicSet := make(map[string]kafkamodel.Topic)
	for _, cg := range consumerGroups {
		id := fmt.Sprintf("%s/%s/%s", cg.ClusterName, cg.ExporterNamespace, cg.ConsumeTopic)

		topicSet[id] = kafkamodel.Topic{
			Name:                   cg.ConsumeTopic,
			ExporterNamespace:      cg.ExporterNamespace,
			ClusterName:            cg.ClusterName,
			AlamedaScalerName:      cg.AlamedaScalerName,
			AlamedaScalerNamespace: cg.AlamedaScalerNamespace,
		}
	}

	topics := make([]kafkamodel.Topic, 0, len(topicSet))
	for _, topic := range topicSet {
		topics = append(topics, topic)
	}
	return topics
}

// listConsumerGroups returns []kafkamodel.ConsumerGroup by AlamedaScaler.Spec.Kafka and current map recorded consumerGroup consuming topics.
func (r AlamedaScalerKafkaReconciler) listConsumerGroups(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, consumerGroupToConsumeTopicsMap map[string][]string) ([]kafkamodel.ConsumerGroup, error) {
	// Get topicToPartitionCountMap for later usage that setting min/max replicas of consumerGroup.
	empty := struct{}{}
	topicSet := make(map[string]struct{})
	for _, topics := range consumerGroupToConsumeTopicsMap {
		for _, topic := range topics {
			topicSet[topic] = empty
		}
	}
	topics := make([]string, 0, len(topicSet))
	for topic := range topicSet {
		topics = append(topics, topic)
	}
	topicToPartitionCountMap, err := r.KafkaClient.ListTopicsPartitionCounts(ctx, topics)
	if err != nil {
		return nil, errors.Wrap(err, "list topics partition counts failed")
	}

	policy := datahubresources.RecommendationPolicy_RECOMMENDATION_POLICY_UNDEFINED.String()
	switch alamedaScaler.Spec.Policy {
	case autoscalingv1alpha1.RecommendationPolicySTABLE:
		policy = datahubresources.RecommendationPolicy_STABLE.String()
	case autoscalingv1alpha1.RecommendationPolicyCOMPACT:
		policy = datahubresources.RecommendationPolicy_COMPACT.String()
	}

	consumerGroupSpecs := alamedaScaler.Spec.Kafka.ConsumerGroups
	consumerGroups := make([]kafkamodel.ConsumerGroup, 0, len(consumerGroupSpecs))
	for _, consumerGroupSpec := range consumerGroupSpecs {
		majorTopic := ""
		if consumerGroupSpec.MajorTopic != nil {
			majorTopic = *consumerGroupSpec.MajorTopic
		}
		topic := chooseTopic(majorTopic, alamedaScaler.Spec.Kafka.Topics, consumerGroupToConsumeTopicsMap[consumerGroupSpec.Name])
		if topic == "" {
			r.Logger.Infof("ConsumerGroup(%s) does not consumes Topic(%s).", consumerGroupSpec.Name, topic)
			continue
		}

		consumerGroup := kafkamodel.ConsumerGroup{
			Name:                   consumerGroupSpec.Name,
			ExporterNamespace:      alamedaScaler.Spec.Kafka.ExporterNamespace,
			ClusterName:            r.ClusterUID,
			AlamedaScalerName:      alamedaScaler.GetName(),
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			Policy:                 policy,
			EnableExecution:        alamedaScaler.IsEnableExecution(),
			ConsumeTopic:           topic,
			ResourceMeta: kafkamodel.ResourceMeta{
				CustomName: consumerGroupSpec.Resource.Custom,
			},
			MinReplicas: 1,
			MaxReplicas: int32(topicToPartitionCountMap[topic]),
		}

		if consumerGroupSpec.MinReplicas != nil {
			consumerGroup.MinReplicas = *consumerGroupSpec.MinReplicas
		}
		if consumerGroupSpec.MaxReplicas != nil {
			consumerGroup.MaxReplicas = *consumerGroupSpec.MaxReplicas
		}

		if consumerGroupSpec.Resource.Kubernetes != nil {
			metadata, err := r.getFirstCreatedMatchedKubernetesMetadata(ctx, alamedaScaler.GetNamespace(), *consumerGroupSpec.Resource.Kubernetes)
			if err != nil {
				return nil, errors.Wrapf(err, "get matched kubernetes resource failed: ConsumerGroup.Name: %s", consumerGroupSpec.Name)
			}
			if (metadata == kafkamodel.KubernetesMeta{}) {
				r.Logger.Warnf("No Kubernetes resource can map to ConsumerGroup: ConsumerGroup: %+v", consumerGroupSpec)
			}
			consumerGroup.KubernetesMeta = metadata
		}

		consumerGroups = append(consumerGroups, consumerGroup)
	}
	return consumerGroups, nil
}

func (r AlamedaScalerKafkaReconciler) getFirstCreatedMatchedKubernetesMetadata(ctx context.Context, namespace string, spec autoscalingv1alpha1.KubernetesResourceSpec) (kafkamodel.KubernetesMeta, error) {
	resourcesLister := utilsresources.NewListResources(r.K8SClient)

	deployments, err := resourcesLister.ListDeploymentsByNamespaceLabels(namespace, spec.Selector.MatchLabels)
	if err != nil {
		return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "list Deployments by namespace and labels failed")
	}
	deploymentConfigs := []openshiftapiappsv1.DeploymentConfig{}
	if r.HasOpenShiftAPIAppsv1 {
		deploymentConfigs, err = resourcesLister.ListDeploymentConfigsByNamespaceLabels(namespace, spec.Selector.MatchLabels)
		if err != nil {
			return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "list DeploymentConfigs by namespace and labels failed")
		}
	}
	statefulSets, err := resourcesLister.ListStatefulSetsByNamespaceLabels(namespace, spec.Selector.MatchLabels)
	if err != nil {
		return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "list StatefulSets by namespace and labels failed")
	}
	if len(deployments) == 0 && len(deploymentConfigs) == 0 && len(statefulSets) == 0 {
		return kafkamodel.KubernetesMeta{}, nil
	}

	kubernetesMetadata := kafkamodel.KubernetesMeta{}
	indexDeployment := 0
	indexStatefulSet := 0
	indexDeploymentConfig := 0
	earliestDeploymentCreationTimestamp := metav1.NewTime(time.Now())
	earliestStatefulSetCreationTimestamp := metav1.NewTime(time.Now())
	earliestDeploymentConfigCreationTimestamp := metav1.NewTime(time.Now())
	for i, r := range deployments {
		if !isMonitoredByAlamedaScalerType(r.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeKafka) {
			continue
		}
		if r.CreationTimestamp.Before(&earliestDeploymentCreationTimestamp) {
			indexDeployment = i
			earliestDeploymentCreationTimestamp = r.CreationTimestamp
		}
	}
	for i, r := range statefulSets {
		if !isMonitoredByAlamedaScalerType(r.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeKafka) {
			continue
		}
		if r.CreationTimestamp.Before(&earliestStatefulSetCreationTimestamp) {
			indexStatefulSet = i
			earliestStatefulSetCreationTimestamp = r.CreationTimestamp
		}
	}
	for i, r := range deploymentConfigs {
		if !isMonitoredByAlamedaScalerType(r.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeKafka) {
			continue
		}
		if r.CreationTimestamp.Before(&earliestDeploymentConfigCreationTimestamp) {
			indexDeploymentConfig = i
			earliestDeploymentConfigCreationTimestamp = r.CreationTimestamp
		}
	}

	switch getFirstTime([]time.Time{
		earliestDeploymentCreationTimestamp.Time,
		earliestStatefulSetCreationTimestamp.Time,
		earliestDeploymentConfigCreationTimestamp.Time,
	}) {
	case earliestDeploymentCreationTimestamp.Time:
		controller := deployments[indexDeployment]
		specReplicas := int32(0)
		if controller.Spec.Replicas != nil {
			specReplicas = *controller.Spec.Replicas
		}

		resource := getTotalResourceFromContainers(controller.Spec.Template.Spec.Containers)

		kubernetesMetadata.Kind = controller.GetObjectKind().GroupVersionKind().Kind
		kubernetesMetadata.Namespace = controller.GetNamespace()
		kubernetesMetadata.Name = controller.GetName()
		kubernetesMetadata.ReadyReplicas = controller.Status.ReadyReplicas
		kubernetesMetadata.SpecReplicas = specReplicas
		kubernetesMetadata.SetResourceRequirements(resource)

		pods, err := resourcesLister.ListPodsByDeployment(controller.GetNamespace(), controller.GetName())
		if err != nil {
			return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "list pods by Deployment failed")
		}
		volumeCapacity := VolumeCapacity{}
		for _, pod := range pods {
			v, err := getVolumeCapacityUsedByPod(ctx, r.K8SClient, pod)
			if err != nil {
				return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "get volume capacity used by Pod failed")
			}
			volumeCapacity.add(v)
		}
		kubernetesMetadata.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
		kubernetesMetadata.VolumesPVCSize = strconv.FormatInt(volumeCapacity.PVC, 10)
	case earliestStatefulSetCreationTimestamp.Time:
		controller := statefulSets[indexStatefulSet]
		specReplicas := int32(0)
		if controller.Spec.Replicas != nil {
			specReplicas = *controller.Spec.Replicas
		}

		resource := getTotalResourceFromContainers(controller.Spec.Template.Spec.Containers)

		kubernetesMetadata.Kind = controller.GetObjectKind().GroupVersionKind().Kind
		kubernetesMetadata.Namespace = controller.GetNamespace()
		kubernetesMetadata.Name = controller.GetName()
		kubernetesMetadata.ReadyReplicas = controller.Status.ReadyReplicas
		kubernetesMetadata.SpecReplicas = specReplicas
		kubernetesMetadata.SetResourceRequirements(resource)

		pods, err := resourcesLister.ListPodsByStatefulSet(controller.GetNamespace(), controller.GetName())
		if err != nil {
			return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "list pods by Deployment failed")
		}
		volumeCapacity := VolumeCapacity{}
		for _, pod := range pods {
			v, err := getVolumeCapacityUsedByPod(ctx, r.K8SClient, pod)
			if err != nil {
				return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "get volume capacity used by Pod failed")
			}
			volumeCapacity.add(v)
		}
		kubernetesMetadata.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
		kubernetesMetadata.VolumesPVCSize = strconv.FormatInt(volumeCapacity.PVC, 10)
	case earliestDeploymentConfigCreationTimestamp.Time:
		controller := deploymentConfigs[indexDeploymentConfig]
		specReplicas := controller.Spec.Replicas

		resource := getTotalResourceFromContainers(controller.Spec.Template.Spec.Containers)

		kubernetesMetadata.Kind = controller.GetObjectKind().GroupVersionKind().Kind
		kubernetesMetadata.Namespace = controller.GetNamespace()
		kubernetesMetadata.Name = controller.GetName()
		kubernetesMetadata.ReadyReplicas = controller.Status.ReadyReplicas
		kubernetesMetadata.SpecReplicas = specReplicas
		kubernetesMetadata.SetResourceRequirements(resource)

		pods, err := resourcesLister.ListPodsByDeploymentConfig(controller.GetNamespace(), controller.GetName())
		if err != nil {
			return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "list pods by Deployment failed")
		}
		volumeCapacity := VolumeCapacity{}
		for _, pod := range pods {
			v, err := getVolumeCapacityUsedByPod(ctx, r.K8SClient, pod)
			if err != nil {
				return kafkamodel.KubernetesMeta{}, errors.Wrap(err, "get volume capacity used by Pod failed")
			}
			volumeCapacity.add(v)
		}
		kubernetesMetadata.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
		kubernetesMetadata.VolumesPVCSize = strconv.FormatInt(volumeCapacity.PVC, 10)
	}

	return kubernetesMetadata, nil
}

// getConsumerGroupToConsumeTopicsMap returns map from consumerGroup to currently consumed topics that records in Kafka.
func (r AlamedaScalerKafkaReconciler) getConsumerGroupToConsumeTopicsMap(ctx context.Context, consumerGroups []string) (map[string][]string, error) {
	type consumeDetail struct {
		consumerGroup string
		consumeTopics []string
	}

	ch := make(chan consumeDetail)
	wg := errgroup.Group{}
	for _, consumerGroup := range consumerGroups {
		copyConsumerGroup := consumerGroup
		wg.Go(func() error {
			consumeTopics, err := r.KafkaClient.ListConsumeTopics(ctx, copyConsumerGroup)
			if err != nil {
				return errors.Wrap(err, "list consume topics failed")
			}

			ch <- consumeDetail{
				consumerGroup: copyConsumerGroup,
				consumeTopics: consumeTopics,
			}
			return nil
		})
	}

	consumerGroupToConsumeTopicsMap := make(map[string][]string)
	done := make(chan bool)
	go func() {
		for consumeDetail := range ch {
			consumerGroupToConsumeTopicsMap[consumeDetail.consumerGroup] = consumeDetail.consumeTopics
		}
		done <- true
	}()

	err := wg.Wait()
	close(ch)
	<-done
	if err != nil {
		return consumerGroupToConsumeTopicsMap, err
	}

	return consumerGroupToConsumeTopicsMap, nil
}

func (r AlamedaScalerKafkaReconciler) syncWithDatahub(ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler, topics []kafkamodel.Topic, consumerGroups []kafkamodel.ConsumerGroup) error {
	r.Logger.Debugf("Synchronize with Datahub. Topics: %+v, ConsumerGroups: %+v", topics, consumerGroups)

	alamedaScalerName := alamedaScaler.Name
	alamedaScalerNamespace := alamedaScaler.Namespace
	exporterNamespace := alamedaScaler.Spec.Kafka.ExporterNamespace

	wg := errgroup.Group{}
	wg.Go(func() error {
		// Create topics
		if err := r.KafkaRepository.CreateTopics(ctx, topics); err != nil {
			return errors.Wrap(err, "creae topics to Datahub failed")
		}
		// Delete topics
		topics, err := r.getTopicsToDelete(ctx, exporterNamespace, alamedaScalerNamespace, alamedaScalerName, topics)
		if err != nil {
			return err
		}
		if err := r.KafkaRepository.DeleteTopics(ctx, topics); err != nil {
			return errors.Wrap(err, "delete topics failed")
		}
		return nil
	})

	wg.Go(func() error {
		// Create consumerGroups
		if err := r.KafkaRepository.CreateConsumerGroups(ctx, consumerGroups); err != nil {
			return errors.Wrap(err, "create consumerGroupDetails to Datahub failed")
		}
		// Delete consumerGroups
		consumerGroups, err := r.getConsumerGroupsToDelete(ctx, exporterNamespace, alamedaScalerNamespace, alamedaScalerName, consumerGroups)
		if err != nil {
			return err
		}
		if err := r.KafkaRepository.DeleteConsumerGroups(ctx, consumerGroups); err != nil {
			return errors.Wrap(err, "delete consumerGroups failed")
		}
		return nil
	})

	return wg.Wait()
}

func (r AlamedaScalerKafkaReconciler) getTopicsToDelete(ctx context.Context, exporterNamespace string, alamedaScalerNamespace, alamedaScalerName string, topics []kafkamodel.Topic) ([]kafkamodel.Topic, error) {
	empty := struct{}{}
	topicsNeedExisting := make(map[kafkamodel.Topic]struct{}, len(topics))
	for _, topic := range topics {
		topicsNeedExisting[topic] = empty
	}

	topics, err := r.KafkaRepository.ListTopics(ctx, kafkarepository.ListTopicsOption{
		ClusterName:            r.ClusterUID,
		ExporterNamespace:      exporterNamespace,
		AlamedaScalerName:      alamedaScalerName,
		AlamedaScalerNamespace: alamedaScalerNamespace,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list topics from Datahub failed")
	}
	topicsToDelete := make([]kafkamodel.Topic, 0)
	for _, topic := range topics {
		if _, exist := topicsNeedExisting[topic]; !exist {
			topicsToDelete = append(topicsToDelete, topic)
		}
	}
	return topicsToDelete, nil
}

func (r AlamedaScalerKafkaReconciler) getConsumerGroupsToDelete(ctx context.Context, exporterNamespace string, alamedaScalerNamespace string, alamedaScalerName string, consumerGroups []kafkamodel.ConsumerGroup) ([]kafkamodel.ConsumerGroup, error) {
	empty := struct{}{}
	consumerGroupsNeedExisting := make(map[kafkamodel.ConsumerGroup]struct{}, len(consumerGroups))
	for _, consumerGroup := range consumerGroups {
		consumerGroupsNeedExisting[consumerGroup] = empty
	}

	consumerGroups, err := r.KafkaRepository.ListConsumerGroups(ctx, kafkarepository.ListConsumerGroupsOption{
		ClusterName:            r.ClusterUID,
		ExporterNamespace:      exporterNamespace,
		AlamedaScalerName:      alamedaScalerName,
		AlamedaScalerNamespace: alamedaScalerNamespace,
	})
	if err != nil {
		return nil, errors.Wrap(err, "list consumerGroups from Datahub failed")
	}
	consumerGroupToDelete := make([]kafkamodel.ConsumerGroup, 0)
	for _, consumerGroup := range consumerGroups {
		if _, exist := consumerGroupsNeedExisting[consumerGroup]; !exist {
			consumerGroupToDelete = append(consumerGroupToDelete, consumerGroup)
		}
	}
	return consumerGroupToDelete, nil
}

func (r AlamedaScalerKafkaReconciler) getKafkaStatus(alamedaScaler autoscalingv1alpha1.AlamedaScaler, details []kafkamodel.ConsumerGroup) autoscalingv1alpha1.KafkaStatus {
	kafkaStatus := autoscalingv1alpha1.KafkaStatus{}
	kafkaStatus.Effective = true
	kafkaStatus.Message = ""
	kafkaStatus.ExporterNamespace = alamedaScaler.Spec.Kafka.ExporterNamespace
	kafkaStatus.ConsumerGroups = r.getKafkaConsumerGroupStatusListFromConsumerGroupDetails(details)

	empty := struct{}{}
	topicSet := make(map[string]struct{})
	for _, cg := range kafkaStatus.ConsumerGroups {
		topicSet[cg.Topic] = empty
	}
	topics := make([]string, 0, len(topicSet))
	for topic := range topicSet {
		topics = append(topics, topic)
	}
	kafkaStatus.Topics = topics
	return kafkaStatus
}

func (r AlamedaScalerKafkaReconciler) getKafkaConsumerGroupStatusListFromConsumerGroupDetails(details []kafkamodel.ConsumerGroup) []autoscalingv1alpha1.KafkaConsumerGroupStatus {
	statusList := make([]autoscalingv1alpha1.KafkaConsumerGroupStatus, 0, len(details))
	for _, detail := range details {
		statusList = append(statusList, r.getKafkaConsumerGroupStatusFromConsumerGroupDetail(detail))
	}
	return statusList
}

func (r AlamedaScalerKafkaReconciler) getKafkaConsumerGroupStatusFromConsumerGroupDetail(detail kafkamodel.ConsumerGroup) autoscalingv1alpha1.KafkaConsumerGroupStatus {
	status := autoscalingv1alpha1.KafkaConsumerGroupStatus{
		Name:  detail.Name,
		Topic: detail.ConsumeTopic,
		Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
			CustomName: detail.CustomName,
		},
		MinReplicas: detail.MinReplicas,
		MaxReplicas: detail.MaxReplicas,
	}

	if !detail.ResourceMeta.KubernetesMeta.IsEmpty() {
		status.Resource.Kubernetes = &autoscalingv1alpha1.KubernetesObjectMetadata{
			Namespace: detail.ResourceMeta.KubernetesMeta.Namespace,
			Name:      detail.ResourceMeta.KubernetesMeta.Name,
			Kind:      detail.ResourceMeta.KubernetesMeta.Kind,
		}
	}

	return status
}

func (r *AlamedaScalerKafkaReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&autoscalingv1alpha1.AlamedaScaler{}).
		Complete(r)
}
