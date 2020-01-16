package controllers

import (
	"context"

	"github.com/pkg/errors"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var listCandidatesFunctions = []listCandidates{
	listCandidatesDefaultAlamedaScaler,
	listCandidatesKafkaAlamedaScaler,
}

type listCandidates = func(context.Context, client.Client, metav1.ObjectMeta) ([]autoscalingv1alpha1.AlamedaScaler, error)

var listCandidatesDefaultAlamedaScaler = func(
	ctx context.Context,
	k8sClient client.Client,
	objectMeta metav1.ObjectMeta,
) ([]autoscalingv1alpha1.AlamedaScaler, error) {

	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err := k8sClient.List(ctx, &alamedaScalerList)
	if err != nil {
		return nil, errors.Wrap(err, "list AlamedaScalers failed")
	}

	candidates := make([]autoscalingv1alpha1.AlamedaScaler, 0)
	for _, alamedaScaler := range alamedaScalerList.Items {
		if !(alamedaScaler.GetType() == autoscalingv1alpha1.AlamedaScalerTypeNotDefine ||
			alamedaScaler.GetType() == autoscalingv1alpha1.AlamedaScalerTypeDefault) {
			continue
		}
		if alamedaScaler.Spec.Selector == nil {
			continue
		}
		if ok := isLabelsSelectedBySelector(*alamedaScaler.Spec.Selector, objectMeta.GetLabels()); ok {
			candidates = append(candidates, alamedaScaler)
		}
	}
	return candidates, nil
}

var listCandidatesKafkaAlamedaScaler = func(
	ctx context.Context,
	k8sClient client.Client,
	objectMeta metav1.ObjectMeta,
) ([]autoscalingv1alpha1.AlamedaScaler, error) {

	alamedaScalerList := autoscalingv1alpha1.AlamedaScalerList{}
	err := k8sClient.List(ctx, &alamedaScalerList)
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
