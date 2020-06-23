package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	"github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"
	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/autoscaling/v1alpha1"
	utilsresources "github.com/containers-ai/alameda/operator/pkg/utils/resources"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	datahubresources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"

	openshiftapiappsv1 "github.com/openshift/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var listCandidatesKafkaAlamedaScaler = func(
	ctx context.Context,
	k8sClient client.Client,
	objectMeta metav1.ObjectMeta,
) ([]autoscalingv1alpha1.AlamedaScaler, error) {

	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err := k8sClient.List(ctx, &alamedaScalerList, &client.ListOptions{Namespace: objectMeta.Namespace})
	if err != nil {
		return nil, errors.Wrap(err, "list AlamedaScalers failed")
	}

	candidates := make([]autoscalingv1alpha1.AlamedaScaler, 0)
	for _, alamedaScaler := range alamedaScalerList.Items {
		if alamedaScaler.GetType() != autoscalingv1alpha1.AlamedaScalerTypeKafka {
			continue
		}
		if alamedaScaler.Spec.Kafka == nil {
			continue
		}
		for _, consumerGroupSpec := range alamedaScaler.Spec.Kafka.ConsumerGroups {
			if consumerGroupSpec.Resource.Kubernetes == nil || consumerGroupSpec.Resource.Kubernetes.Selector == nil {
				continue
			}
			if ok := isLabelsSelectedBySelector(*consumerGroupSpec.Resource.Kubernetes.Selector, objectMeta.GetLabels()); ok {
				candidates = append(candidates, alamedaScaler)
			}
		}
	}
	return candidates, nil
}

func init() {
	RegisterAlamedaScalerController(autoscalingv1alpha1.AlamedaScalerTypeKafka, listCandidatesKafkaAlamedaScaler)
}

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

	K8SClient        client.Client
	Scheme           *runtime.Scheme
	KafkaClient      kafka.Client
	PrometheusClient prometheus.Prometheus
	DatahubClient    *datahubpkg.Client
	ReconcileTimeout time.Duration

	Logger *log.Scope

	NeededMetrics []string
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

	if !r.isAlamedaScalerTypeNeedToBeReconciled(alamedaScaler) {
		return ctrl.Result{Requeue: false}, nil
	}

	if ok, err := r.isCreateOrOwnLock(ctx, alamedaScaler); err != nil {
		r.Logger.Infof("Check if AlamedaScaler(%s/%s) needs to be reconciled failed, retry reconciling: %s", req.Namespace, req.Name, err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	} else if !ok {
		alamedaScaler.Status.Kafka.Effective = false
		alamedaScaler.Status.Kafka.Message = "Other AlamedaScaler is effective."
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

	if err := r.openClient(); err != nil {
		r.Logger.Warnf("Open AlamedaScaler(%s/%s) clients' connection failed: err: %+v", req.Namespace, req.Name, err)
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	r.Logger.Infof("Reconciling AlamedaScaler(%s/%s)...", req.Namespace, req.Name)
	consumerGroupDetails, err := r.prepareConsumerGroupDetails(ctx, alamedaScaler)
	r.Logger.Debugf("Consumer group details %+v", consumerGroupDetails)
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

	r.Logger.Infof("Reconcile AlamedaScaler(%s/%s) done.", req.Namespace, req.Name)
	return ctrl.Result{Requeue: false}, nil
}

func (r AlamedaScalerKafkaReconciler) handleDeletion(ctx context.Context, namespace, name string) error {
	wg := errgroup.Group{}
	wg.Go(func() error {
		delErr := r.DatahubClient.DeleteByOpts(
			&entities.ApplicationKafkaTopic{}, datahubpkg.Option{
				Entity: entities.ApplicationKafkaTopic{
					ClusterName:            r.ClusterUID,
					AlamedaScalerNamespace: namespace,
					AlamedaScalerName:      name,
				},
				Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
			})
		if delErr != nil {
			return fmt.Errorf(
				"delete topics for scaler (%s/%s) from Datahub failed: %s",
				namespace, name, delErr.Error())
		}
		return nil
	})

	wg.Go(func() error {
		delErr := r.DatahubClient.DeleteByOpts(
			&entities.ApplicationKafkaConsumerGroup{}, datahubpkg.Option{
				Entity: entities.ApplicationKafkaConsumerGroup{
					ClusterName:            r.ClusterUID,
					AlamedaScalerNamespace: namespace,
					AlamedaScalerName:      name,
				},
				Fields: []string{"ClusterName", "AlamedaScalerNamespace", "AlamedaScalerName"},
			})
		if delErr != nil {
			return fmt.Errorf(
				"delete consumergroups for scaler (%s/%s) from Datahub failed: %s",
				namespace, name, delErr.Error())
		}
		return nil
	})
	return wg.Wait()
}

func (r AlamedaScalerKafkaReconciler) updateAlamedaScaler(
	ctx context.Context, alamedaScaler *autoscalingv1alpha1.AlamedaScaler) error {
	if err := r.K8SClient.Update(ctx, alamedaScaler); err != nil {
		return errors.Wrap(err, "update AlamedaScaler failed")
	}
	return nil
}

func (r AlamedaScalerKafkaReconciler) isAlamedaScalerTypeNeedToBeReconciled(
	alamedaScaler autoscalingv1alpha1.AlamedaScaler) bool {
	if alamedaScaler.GetType() != autoscalingv1alpha1.AlamedaScalerTypeKafka {
		return false
	}
	return true
}

func (r AlamedaScalerKafkaReconciler) isCreateOrOwnLock(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) (
	bool, error) {
	//Ensure that count of the AlamedaScaler with same value of AlamedaScaler.Spec.Kafka.ExporterNamespace per namespace at most one.
	lock, err := r.getOrCreateLock(ctx, alamedaScaler)
	if err != nil {
		return false, errors.Wrap(err, "get or create lock failed")
	}
	if !r.isLockOwnBy(&lock, &alamedaScaler) {
		return false, nil
	}

	return true, nil
}

func (r AlamedaScalerKafkaReconciler) getOrCreateLock(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) (corev1.ConfigMap, error) {
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

func (r AlamedaScalerKafkaReconciler) prepareConsumerGroupDetails(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler) (
	[]entities.ApplicationKafkaConsumerGroup, error) {
	// Get map from consumerGroupName to consuming topics from Kafka client.
	consumerGroupNames := make([]string, 0, len(alamedaScaler.Spec.Kafka.ConsumerGroups))
	for _, consumerGroupSpec := range alamedaScaler.Spec.Kafka.ConsumerGroups {
		consumerGroupNames = append(consumerGroupNames, consumerGroupSpec.Name)
	}
	consumerGroupToConsumeTopicsMap, err := r.getConsumerGroupToConsumeTopicsMap(ctx, consumerGroupNames)
	if err != nil {
		return nil, errors.Wrap(err, "get consumerGroup consume topics map failed")
	}
	r.Logger.Debugf("Consumer group consumed topics mapping %+v", consumerGroupToConsumeTopicsMap)

	consumerGroupDetails, err := r.listConsumerGroups(ctx, alamedaScaler, consumerGroupToConsumeTopicsMap)
	if err != nil {
		return nil, errors.Wrap(err, "list consumerGroups' detail failed")
	}

	return consumerGroupDetails, nil
}

func (r AlamedaScalerKafkaReconciler) prepareTopics(consumerGroups []entities.ApplicationKafkaConsumerGroup) []entities.ApplicationKafkaTopic {
	topicSet := make(map[string]entities.ApplicationKafkaTopic)
	for _, cg := range consumerGroups {
		id := fmt.Sprintf("%s/%s/%s", cg.ClusterName, cg.Namespace, cg.TopicName)

		topicSet[id] = entities.ApplicationKafkaTopic{
			Name:                   cg.TopicName,
			Namespace:              cg.Namespace,
			ClusterName:            cg.ClusterName,
			AlamedaScalerName:      cg.AlamedaScalerName,
			AlamedaScalerNamespace: cg.AlamedaScalerNamespace,
		}
	}

	topics := make([]entities.ApplicationKafkaTopic, 0, len(topicSet))
	for _, topic := range topicSet {
		topics = append(topics, topic)
	}
	return topics
}

// listConsumerGroups returns []kafkamodel.ConsumerGroup by AlamedaScaler.Spec.Kafka and current map recorded consumerGroup consuming topics.
func (r AlamedaScalerKafkaReconciler) listConsumerGroups(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler,
	consumerGroupToConsumeTopicsMap map[string][]string) (
	[]entities.ApplicationKafkaConsumerGroup, error) {
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
	consumerGroups := make([]entities.ApplicationKafkaConsumerGroup, 0, len(consumerGroupSpecs))
	for _, consumerGroupSpec := range consumerGroupSpecs {
		topic := chooseTopic(consumerGroupSpec.MajorTopic, alamedaScaler.Spec.Kafka.Topics,
			consumerGroupToConsumeTopicsMap[consumerGroupSpec.Name])
		r.Logger.Debugf("Consumer group %s is consuming topic %s", consumerGroupSpec.Name, topic)
		if topic == "" {
			r.Logger.Infof("ConsumerGroup(%s) does not consumes Topic(%s).", consumerGroupSpec.Name, topic)
			continue
		}

		consumerGroup := entities.ApplicationKafkaConsumerGroup{
			Name:                   consumerGroupSpec.Name,
			Namespace:              alamedaScaler.Spec.Kafka.ExporterNamespace,
			ClusterName:            r.ClusterUID,
			AlamedaScalerName:      alamedaScaler.GetName(),
			AlamedaScalerNamespace: alamedaScaler.GetNamespace(),
			Policy:                 policy,
			EnableExecution:        alamedaScaler.IsEnableExecution(),
			TopicName:              topic,
			ResourceCustomName:     consumerGroupSpec.Resource.Custom,
			ResourceK8sMinReplicas: 1,
			ResourceK8sMaxReplicas: int32(topicToPartitionCountMap[topic]),
		}

		if consumerGroupSpec.MinReplicas != nil {
			consumerGroup.ResourceK8sMinReplicas = *consumerGroupSpec.MinReplicas
		}
		if consumerGroupSpec.MaxReplicas != nil {
			consumerGroup.ResourceK8sMaxReplicas = *consumerGroupSpec.MaxReplicas
		}

		if consumerGroupSpec.Resource.Kubernetes != nil {
			metadata, err := r.getFirstCreatedMatchedKubernetesMetadata(
				ctx, alamedaScaler.GetNamespace(), *consumerGroupSpec.Resource.Kubernetes)
			if err != nil {
				return nil,
					errors.Wrapf(err, "get matched kubernetes resource failed: ConsumerGroup.Name: %s",
						consumerGroupSpec.Name)
			}
			if r.isConsumerGroupEmpty(metadata) {
				r.Logger.Warnf(
					"No Kubernetes resource can map to ConsumerGroup: ConsumerGroup: %+v",
					consumerGroupSpec)
			}
			consumerGroup = metadata
		}

		consumerGroups = append(consumerGroups, consumerGroup)
	}
	return consumerGroups, nil
}

func (r AlamedaScalerKafkaReconciler) getFirstCreatedMatchedKubernetesMetadata(
	ctx context.Context, namespace string, spec autoscalingv1alpha1.KubernetesResourceSpec) (
	entities.ApplicationKafkaConsumerGroup, error) {
	consumerGroupEntity := entities.ApplicationKafkaConsumerGroup{}
	resourcesLister := utilsresources.NewListResources(r.K8SClient)
	deployments, err := resourcesLister.ListDeploymentsByNamespaceLabels(
		namespace, spec.Selector.MatchLabels)
	if err != nil {
		return consumerGroupEntity,
			errors.Wrap(err, "list Deployments by namespace and labels failed")
	}
	r.Logger.Debugf(
		"List deployments by matched namespace %s and labels %+v and get result %+v",
		namespace, spec.Selector.MatchLabels, deployments)
	deploymentConfigs := []openshiftapiappsv1.DeploymentConfig{}
	if r.HasOpenShiftAPIAppsv1 {
		deploymentConfigs, err = resourcesLister.ListDeploymentConfigsByNamespaceLabels(
			namespace, spec.Selector.MatchLabels)
		if err != nil {
			return consumerGroupEntity,
				errors.Wrap(err, "list DeploymentConfigs by namespace and labels failed")
		}
	}
	r.Logger.Debugf("List deploymentconfigs by matched namespace %s and labels %+v and get result %+v",
		namespace, spec.Selector.MatchLabels, deploymentConfigs)
	statefulSets, err := resourcesLister.ListStatefulSetsByNamespaceLabels(
		namespace, spec.Selector.MatchLabels)
	if err != nil {
		return consumerGroupEntity,
			errors.Wrap(err, "list StatefulSets by namespace and labels failed")
	}
	r.Logger.Debugf("List statefulsets by matched namespace %s and labels %+v and get result %+v",
		namespace, spec.Selector.MatchLabels, statefulSets)
	if len(deployments) == 0 && len(deploymentConfigs) == 0 && len(statefulSets) == 0 {
		return consumerGroupEntity, nil
	}

	indexDeployment := 0
	indexStatefulSet := 0
	indexDeploymentConfig := 0
	earliestDeploymentCreationTimestamp := metav1.NewTime(time.Now())
	earliestStatefulSetCreationTimestamp := metav1.NewTime(time.Now())
	earliestDeploymentConfigCreationTimestamp := metav1.NewTime(time.Now())
	for i, r := range deployments {
		if !IsMonitoredByAlamedaScalerController(r.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeKafka) {
			continue
		}
		if r.CreationTimestamp.Before(&earliestDeploymentCreationTimestamp) {
			indexDeployment = i
			earliestDeploymentCreationTimestamp = r.CreationTimestamp
		}
	}
	for i, r := range statefulSets {
		if !IsMonitoredByAlamedaScalerController(r.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeKafka) {
			continue
		}
		if r.CreationTimestamp.Before(&earliestStatefulSetCreationTimestamp) {
			indexStatefulSet = i
			earliestStatefulSetCreationTimestamp = r.CreationTimestamp
		}
	}
	for i, r := range deploymentConfigs {
		if !IsMonitoredByAlamedaScalerController(r.ObjectMeta, autoscalingv1alpha1.AlamedaScalerTypeKafka) {
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

		consumerGroupEntity.ResourceK8sKind = controller.GetObjectKind().GroupVersionKind().Kind
		consumerGroupEntity.ResourceK8sNamespace = controller.GetNamespace()
		consumerGroupEntity.ResourceK8sName = controller.GetName()
		consumerGroupEntity.ResourceK8sReplicas = controller.Status.ReadyReplicas
		consumerGroupEntity.ResourceK8sSpecReplicas = specReplicas
		consumerGroupEntity.ResourceCPULimit =
			strconv.FormatInt(resource.Limits.Cpu().MilliValue(), 10)
		consumerGroupEntity.ResourceCPURequest =
			strconv.FormatInt(resource.Requests.Cpu().MilliValue(), 10)
		consumerGroupEntity.ResourceMemoryLimit =
			strconv.FormatInt(resource.Limits.Memory().Value(), 10)
		consumerGroupEntity.ResourceMemoryRequest =
			strconv.FormatInt(resource.Requests.Memory().Value(), 10)
		r.Logger.Debugf("deployment create timestamp compared get replicas: %v", controller.Status.ReadyReplicas)
		pods, err := resourcesLister.ListPodsByDeployment(controller.GetNamespace(), controller.GetName())
		if err != nil {
			return consumerGroupEntity, errors.Wrap(err, "list pods by Deployment failed")
		}
		volumeCapacity := VolumeCapacity{}
		for _, pod := range pods {
			v, err := getVolumeCapacityUsedByPod(ctx, r.K8SClient, pod)
			if err != nil {
				return consumerGroupEntity, errors.Wrap(err, "get volume capacity used by Pod failed")
			}
			volumeCapacity.add(v)
		}
		consumerGroupEntity.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
		consumerGroupEntity.VolumesPvcSize = strconv.FormatInt(volumeCapacity.PVC, 10)
	case earliestStatefulSetCreationTimestamp.Time:
		controller := statefulSets[indexStatefulSet]
		specReplicas := int32(0)
		if controller.Spec.Replicas != nil {
			specReplicas = *controller.Spec.Replicas
		}

		resource := getTotalResourceFromContainers(controller.Spec.Template.Spec.Containers)
		consumerGroupEntity.ResourceK8sKind = controller.GetObjectKind().GroupVersionKind().Kind
		consumerGroupEntity.ResourceK8sNamespace = controller.GetNamespace()
		consumerGroupEntity.ResourceK8sName = controller.GetName()
		consumerGroupEntity.ResourceK8sReplicas = controller.Status.ReadyReplicas
		consumerGroupEntity.ResourceK8sSpecReplicas = specReplicas
		consumerGroupEntity.ResourceCPULimit =
			strconv.FormatInt(resource.Limits.Cpu().MilliValue(), 10)
		consumerGroupEntity.ResourceCPURequest =
			strconv.FormatInt(resource.Requests.Cpu().MilliValue(), 10)
		consumerGroupEntity.ResourceMemoryLimit =
			strconv.FormatInt(resource.Limits.Memory().Value(), 10)
		consumerGroupEntity.ResourceMemoryRequest =
			strconv.FormatInt(resource.Requests.Memory().Value(), 10)
		r.Logger.Debugf("statefulset create timestamp compared get replicas: %v",
			controller.Status.ReadyReplicas)
		pods, err := resourcesLister.ListPodsByStatefulSet(
			controller.GetNamespace(), controller.GetName())
		if err != nil {
			return consumerGroupEntity, errors.Wrap(err, "list pods by Deployment failed")
		}
		volumeCapacity := VolumeCapacity{}
		for _, pod := range pods {
			v, err := getVolumeCapacityUsedByPod(ctx, r.K8SClient, pod)
			if err != nil {
				return consumerGroupEntity, errors.Wrap(err, "get volume capacity used by Pod failed")
			}
			volumeCapacity.add(v)
		}
		consumerGroupEntity.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
		consumerGroupEntity.VolumesPvcSize = strconv.FormatInt(volumeCapacity.PVC, 10)
	case earliestDeploymentConfigCreationTimestamp.Time:
		controller := deploymentConfigs[indexDeploymentConfig]
		specReplicas := controller.Spec.Replicas

		resource := getTotalResourceFromContainers(controller.Spec.Template.Spec.Containers)

		consumerGroupEntity.ResourceK8sKind = controller.GetObjectKind().GroupVersionKind().Kind
		consumerGroupEntity.ResourceK8sNamespace = controller.GetNamespace()
		consumerGroupEntity.ResourceK8sName = controller.GetName()
		consumerGroupEntity.ResourceK8sReplicas = controller.Status.ReadyReplicas
		consumerGroupEntity.ResourceK8sSpecReplicas = specReplicas
		consumerGroupEntity.ResourceCPULimit =
			strconv.FormatInt(resource.Limits.Cpu().MilliValue(), 10)
		consumerGroupEntity.ResourceCPURequest =
			strconv.FormatInt(resource.Requests.Cpu().MilliValue(), 10)
		consumerGroupEntity.ResourceMemoryLimit =
			strconv.FormatInt(resource.Limits.Memory().Value(), 10)
		consumerGroupEntity.ResourceMemoryRequest =
			strconv.FormatInt(resource.Requests.Memory().Value(), 10)
		r.Logger.Debugf("deploymentconfig create timestamp compared get replicas: %v", controller.Status.ReadyReplicas)
		pods, err := resourcesLister.ListPodsByDeploymentConfig(controller.GetNamespace(), controller.GetName())
		if err != nil {
			return consumerGroupEntity, errors.Wrap(err, "list pods by Deployment failed")
		}
		volumeCapacity := VolumeCapacity{}
		for _, pod := range pods {
			v, err := getVolumeCapacityUsedByPod(ctx, r.K8SClient, pod)
			if err != nil {
				return consumerGroupEntity, errors.Wrap(err, "get volume capacity used by Pod failed")
			}
			volumeCapacity.add(v)
		}
		consumerGroupEntity.VolumesSize = strconv.FormatInt(volumeCapacity.Total, 10)
		consumerGroupEntity.VolumesPvcSize = strconv.FormatInt(volumeCapacity.PVC, 10)
	}

	return consumerGroupEntity, nil
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
			r.Logger.Debugf("Consumer group %s has topics %+v", copyConsumerGroup, consumeTopics)
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
			r.Logger.Debugf("Consumer detail received %+v", consumeDetail)
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

func (r AlamedaScalerKafkaReconciler) syncWithDatahub(
	ctx context.Context, alamedaScaler autoscalingv1alpha1.AlamedaScaler,
	topics []entities.ApplicationKafkaTopic,
	consumerGroups []entities.ApplicationKafkaConsumerGroup) error {
	r.Logger.Debugf("Synchronize with Datahub. Topics: %+v, ConsumerGroups: %+v", topics, consumerGroups)

	alamedaScalerName := alamedaScaler.Name
	alamedaScalerNamespace := alamedaScaler.Namespace
	exporterNamespace := alamedaScaler.Spec.Kafka.ExporterNamespace

	wg := errgroup.Group{}
	wg.Go(func() error {
		// Create topics
		if err := r.DatahubClient.Create(&topics); err != nil {
			return errors.Wrap(err, "creae topics to Datahub failed")
		}
		// Delete topics
		topics, err := r.getTopicsToDelete(ctx, exporterNamespace, alamedaScalerNamespace, alamedaScalerName, topics)
		if err != nil {
			return err
		}
		if err := r.DatahubClient.Delete(&topics); err != nil {
			return errors.Wrap(err, "delete topics failed")
		}
		return nil
	})

	wg.Go(func() error {
		// Create consumerGroups
		if err := r.DatahubClient.Create(&consumerGroups); err != nil {
			return errors.Wrap(err, "create consumerGroupDetails to Datahub failed")
		}
		// Delete consumerGroups
		consumerGroups, err := r.getConsumerGroupsToDelete(ctx, exporterNamespace, alamedaScalerNamespace, alamedaScalerName, consumerGroups)
		if err != nil {
			return err
		}
		if err := r.DatahubClient.Delete(&consumerGroups); err != nil {
			return errors.Wrap(err, "delete consumerGroups failed")
		}
		return nil
	})

	return wg.Wait()
}

func (r AlamedaScalerKafkaReconciler) getTopicsToDelete(
	ctx context.Context, exporterNamespace string, alamedaScalerNamespace, alamedaScalerName string,
	topics []entities.ApplicationKafkaTopic) ([]entities.ApplicationKafkaTopic, error) {
	empty := struct{}{}
	topicsNeedExisting := make(map[string]struct{}, len(topics))
	for _, topic := range topics {
		topicsNeedExisting[r.getTopicId(topic)] = empty
	}

	existingTopics := []entities.ApplicationKafkaTopic{}
	err := r.DatahubClient.List(&existingTopics, datahubpkg.Option{
		Entity: entities.ApplicationKafkaTopic{
			ClusterName:            r.ClusterUID,
			Namespace:              exporterNamespace,
			AlamedaScalerName:      alamedaScalerName,
			AlamedaScalerNamespace: alamedaScalerNamespace,
		},
		Fields: []string{
			"ClusterName", "Namespace", "AlamedaScalerName", "AlamedaScalerNamespace"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "list topics from Datahub failed")
	}
	topicsToDelete := make([]entities.ApplicationKafkaTopic, 0)
	for _, topic := range existingTopics {
		if _, exist := topicsNeedExisting[r.getTopicId(topic)]; !exist {
			topicsToDelete = append(topicsToDelete, topic)
		}
	}
	return topicsToDelete, nil
}

func (r AlamedaScalerKafkaReconciler) getConsumerGroupsToDelete(
	ctx context.Context, exporterNamespace, alamedaScalerNamespace, alamedaScalerName string,
	consumerGroups []entities.ApplicationKafkaConsumerGroup) ([]entities.ApplicationKafkaConsumerGroup, error) {
	empty := struct{}{}
	consumerGroupsNeedExisting := make(map[string]struct{}, len(consumerGroups))
	for _, consumerGroup := range consumerGroups {
		consumerGroupsNeedExisting[r.getConsumerGroupId(consumerGroup)] = empty
	}

	existingConsumerGroups := []entities.ApplicationKafkaConsumerGroup{}
	err := r.DatahubClient.List(&existingConsumerGroups, datahubpkg.Option{
		Entity: entities.ApplicationKafkaConsumerGroup{
			ClusterName:            r.ClusterUID,
			Namespace:              exporterNamespace,
			AlamedaScalerName:      alamedaScalerName,
			AlamedaScalerNamespace: alamedaScalerNamespace,
		},
		Fields: []string{
			"ClusterName", "Namespace", "AlamedaScalerName", "AlamedaScalerNamespace"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "list consumerGroups from Datahub failed")
	}
	consumerGroupToDelete := make([]entities.ApplicationKafkaConsumerGroup, 0)
	for _, consumerGroup := range existingConsumerGroups {
		if _, exist := consumerGroupsNeedExisting[r.getConsumerGroupId(consumerGroup)]; !exist {
			consumerGroupToDelete = append(consumerGroupToDelete, consumerGroup)
		}
	}
	return consumerGroupToDelete, nil
}

func (r AlamedaScalerKafkaReconciler) getKafkaStatus(
	alamedaScaler autoscalingv1alpha1.AlamedaScaler, details []entities.ApplicationKafkaConsumerGroup) autoscalingv1alpha1.KafkaStatus {
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

func (r AlamedaScalerKafkaReconciler) getKafkaConsumerGroupStatusListFromConsumerGroupDetails(
	details []entities.ApplicationKafkaConsumerGroup) []autoscalingv1alpha1.KafkaConsumerGroupStatus {
	statusList := make([]autoscalingv1alpha1.KafkaConsumerGroupStatus, 0, len(details))
	for _, detail := range details {
		statusList = append(statusList, r.getKafkaConsumerGroupStatusFromConsumerGroupDetail(detail))
	}
	return statusList
}

func (r AlamedaScalerKafkaReconciler) getKafkaConsumerGroupStatusFromConsumerGroupDetail(
	detail entities.ApplicationKafkaConsumerGroup) autoscalingv1alpha1.KafkaConsumerGroupStatus {
	status := autoscalingv1alpha1.KafkaConsumerGroupStatus{
		Name:  detail.Name,
		Topic: detail.TopicName,
		Resource: autoscalingv1alpha1.KafkaConsumerGroupResourceMetadata{
			CustomName: detail.ResourceCustomName,
		},
		MinReplicas: detail.ResourceK8sMinReplicas,
		MaxReplicas: detail.ResourceK8sMaxReplicas,
	}

	if !r.isConsumerGroupEmpty(detail) {
		status.Resource.Kubernetes = &autoscalingv1alpha1.KubernetesObjectMetadata{
			Namespace: detail.ResourceK8sNamespace,
			Name:      detail.ResourceK8sName,
			Kind:      detail.ResourceK8sKind,
		}
	}

	return status
}

func (r *AlamedaScalerKafkaReconciler) isConsumerGroupEmpty(
	cg entities.ApplicationKafkaConsumerGroup) bool {
	return cg.Name == "" && cg.Namespace == "" && cg.ClusterName == "" &&
		cg.TopicName == "" && cg.AlamedaScalerName == "" &&
		cg.AlamedaScalerNamespace == "" && cg.ResourceK8sName == "" &&
		cg.ResourceK8sNamespace == "" && cg.ResourceK8sKind == ""
}

func (r AlamedaScalerKafkaReconciler) getTopicId(
	topic entities.ApplicationKafkaTopic) string {
	return fmt.Sprintf("%s%s%s%s%s", topic.ClusterName, topic.Name,
		topic.Namespace, topic.AlamedaScalerNamespace, topic.Name)
}

func (r AlamedaScalerKafkaReconciler) getConsumerGroupId(
	consumerGroup entities.ApplicationKafkaConsumerGroup) string {
	return fmt.Sprintf("%s%s%s%s%s%s%s%s%s", consumerGroup.Name,
		consumerGroup.Namespace, consumerGroup.ClusterName,
		consumerGroup.TopicName, consumerGroup.AlamedaScalerName,
		consumerGroup.AlamedaScalerNamespace, consumerGroup.ResourceK8sName,
		consumerGroup.ResourceK8sNamespace, consumerGroup.ResourceK8sKind)
}

func (r *AlamedaScalerKafkaReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&autoscalingv1alpha1.AlamedaScaler{}).
		Complete(r)
}
