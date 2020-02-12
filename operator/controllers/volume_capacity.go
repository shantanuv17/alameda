package controllers

import (
	"context"
	"reflect"

	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type VolumeCapacity struct {
	Total int64
	PVC   int64
}

func getVolumeCapacityUsedByPod(ctx context.Context, k8sClient client.Client, pod corev1.Pod) (VolumeCapacity, error) {
	totalSize := int64(0)
	pvcsSize := int64(0)
	pvcsName := make([]string, 0)
	for _, volume := range pod.Spec.Volumes {
		if volume.PersistentVolumeClaim != nil {
			pvcsName = append(pvcsName, volume.PersistentVolumeClaim.ClaimName)
		}
	}

	namespace := pod.GetNamespace()
	for _, name := range pvcsName {
		pvc := corev1.PersistentVolumeClaim{}
		err := k8sClient.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, &pvc)
		if err != nil {
			return VolumeCapacity{}, errors.Wrap(err, "get pvc failed")
		}
		quantity := pvc.Status.Capacity[corev1.ResourceStorage]
		size := quantity.Value()
		totalSize += size
		pvcsSize += size
	}

	return VolumeCapacity{
		Total: totalSize,
		PVC:   pvcsSize,
	}, nil
}

func (v *VolumeCapacity) add(in VolumeCapacity) {
	inV := reflect.ValueOf(in)

	ptrV := reflect.ValueOf(v)
	vV := ptrV.Elem()
	for i := 0; i < inV.NumField(); i++ {
		vV.Field(i).SetInt(vV.Field(i).Int() + inV.Field(i).Int())
	}
}
