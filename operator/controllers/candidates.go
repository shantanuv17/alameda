package controllers

import (
	"context"
	"sync"

	autoscalingv1alpha1 "github.com/containers-ai/alameda/operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	lock sync.Mutex

	listCandidatesFunctionMap = make(map[string]ListMonitoringAlamedaScaler)
)

type ListMonitoringAlamedaScaler = func(context.Context, client.Client, metav1.ObjectMeta) ([]autoscalingv1alpha1.AlamedaScaler, error)

// RegisterAlamedaScalerController registers the controller name to listCandidatesFunctionMap with function provieded in arg 2.
func RegisterAlamedaScalerController(name string, f ListMonitoringAlamedaScaler) {
	lock.Lock()
	defer lock.Unlock()
	listCandidatesFunctionMap[name] = f
}

func GetAlamedaScalerControllerName(ctx context.Context, k8sClient client.Client, objectMeta metav1.ObjectMeta) (string, error) {
	uidToControllerNameMap := make(map[types.UID]string)
	candidates := make([]metav1.ObjectMeta, 0)
	for t, f := range listCandidatesFunctionMap {
		alamedaScalers, err := f(ctx, k8sClient, objectMeta)
		if err != nil {
			return "", err
		}
		for _, alamedaScaler := range alamedaScalers {
			candidates = append(candidates, alamedaScaler.ObjectMeta)
			uidToControllerNameMap[alamedaScaler.GetUID()] = t
		}
	}
	if len(candidates) == 0 {
		return "", nil
	}

	oldestObj := getFirstCreatedObjectMeta(candidates)
	alamedaScaler := autoscalingv1alpha1.AlamedaScaler{}
	if err := k8sClient.Get(ctx, client.ObjectKey{Namespace: oldestObj.GetNamespace(), Name: oldestObj.GetName()}, &alamedaScaler); err != nil {
		return "", err
	}

	return uidToControllerNameMap[alamedaScaler.GetUID()], nil
}
