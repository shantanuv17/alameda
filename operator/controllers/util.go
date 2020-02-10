package controllers

import (
	"context"
	"time"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	alamedaScalerNameAnnotationKey = "alamedascalers.autoscaling.containers.ai/name"
	alamedaScalerTypeAnnotationKey = "alamedascalers.autoscaling.containers.ai/type"
)

// setLastMonitorAlamedaScaler sets the last AlamedaScaler's name into the object's annotation
func setLastMonitorAlamedaScaler(obj metav1.Object, alamedaScalerName string) {

	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[alamedaScalerNameAnnotationKey] = alamedaScalerName

	obj.SetAnnotations(annotations)
}

// setLastMonitorAlamedaScalerType sets the last AlamedaScaler's type into the object's annotation.
func setLastMonitorAlamedaScalerType(obj metav1.Object, alamedaScalerType string) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[alamedaScalerTypeAnnotationKey] = alamedaScalerType

	obj.SetAnnotations(annotations)
}

func getCandidateAlamedaScaler(ctx context.Context, k8sClient client.Client, objectMeta metav1.ObjectMeta) (autoscalingv1alpha1.AlamedaScaler, error) {
	candidates := make([]metav1.ObjectMeta, 0)
	for _, f := range listCandidatesFunctions {
		alamedaScalers, err := f(ctx, k8sClient, objectMeta)
		if err != nil {
			return autoscalingv1alpha1.AlamedaScaler{}, err
		}
		for _, alamedaScaler := range alamedaScalers {
			candidates = append(candidates, alamedaScaler.ObjectMeta)
		}
	}
	if len(candidates) == 0 {
		return autoscalingv1alpha1.AlamedaScaler{}, nil
	}

	oldestObj := getFirstCreatedObjectMeta(candidates)
	alamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Namespace: oldestObj.GetNamespace(), Name: oldestObj.GetName()}, &alamedaScaler); err != nil {
		return autoscalingv1alpha1.AlamedaScaler{}, err
	}
	return alamedaScaler, nil
}

// isMonitoredByAlamedaScalerType returns if the last monitored AlamedaScaler's type to object if alamedaScalerType.
func isMonitoredByAlamedaScalerType(obj metav1.ObjectMeta, alamedaScalerType string) bool {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return false
	}
	if annotations[alamedaScalerTypeAnnotationKey] == alamedaScalerType {
		return true
	}
	return false
}

func isLabelsSelectedBySelector(selector metav1.LabelSelector, labels map[string]string) bool {
	for k, v := range selector.MatchLabels {
		v2, exist := labels[k]
		if v != v2 || !exist {
			return false
		}
	}

	return true
}

func getFirstCreatedObjectMeta(objs []metav1.ObjectMeta) metav1.ObjectMeta {
	t := time.Now()
	firstCreatedObe := metav1.ObjectMeta{}
	for _, obj := range objs {
		if obj.GetCreationTimestamp().UnixNano() < t.UnixNano() {
			firstCreatedObe = obj
			t = obj.GetCreationTimestamp().Time
		}
	}
	return firstCreatedObe
}

func getTotalResourceFromContainers(containers []corev1.Container) corev1.ResourceRequirements {
	total := corev1.ResourceRequirements{
		Limits:   corev1.ResourceList{},
		Requests: corev1.ResourceList{},
	}
	for _, c := range containers {
		for resourceName, quantity := range c.Resources.Limits {
			q := total.Limits[resourceName]
			q.Add(quantity)
			total.Limits[resourceName] = q
		}
		for resourceName, quantity := range c.Resources.Requests {
			q := total.Requests[resourceName]
			q.Add(quantity)
			total.Requests[resourceName] = q
		}
	}

	return total
}

func getFirstTime(times []time.Time)time.Time{
	min := time.Now()
	for _,t:=range times{
		if min.After(t){
			min=t
		}
	}
	
	return min
}