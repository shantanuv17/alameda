package controllers

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	alamedaScalerTypeAnnotationKey = "alamedascalers.autoscaling.containers.ai/type"
)

// setAlamedaScalerControllerName sets the AlamedaScaler controller's type into the object's annotation.
func setAlamedaScalerControllerName(obj metav1.Object, name string) {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[alamedaScalerTypeAnnotationKey] = name

	obj.SetAnnotations(annotations)
}

// IsMonitoredByAlamedaScalerController returns if the object is monitored by the provided name of AlamdeScaler controller.
func IsMonitoredByAlamedaScalerController(obj metav1.ObjectMeta, name string) bool {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		return false
	}
	if annotations[alamedaScalerTypeAnnotationKey] == name {
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

func getFirstTime(times []time.Time) time.Time {
	min := time.Now()
	for _, t := range times {
		if min.After(t) {
			min = t
		}
	}

	return min
}
