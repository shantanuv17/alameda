package controllers

import (
	"time"

	operatorutils "github.com/containers-ai/alameda/operator/pkg/utils"
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

func GetTotalResourceFromContainers(containers []corev1.Container) corev1.ResourceRequirements {
	return operatorutils.GetTotalResourceFromContainers(containers)
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
