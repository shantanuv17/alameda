package pod

import (
	CoreV1 "k8s.io/api/core/v1"
	AlamedaUtils "prophetstor.com/alameda/operator/pkg/utils/resources"
	AlamedaConsts "prophetstor.com/alameda/pkg/consts"
	AlamedaEnum "prophetstor.com/alameda/pkg/utils/datahub/enumconv"
	AlamedaLog "prophetstor.com/alameda/pkg/utils/log"
	ApiResources "prophetstor.com/api/datahub/resources"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	scope = AlamedaLog.RegisterScope("datahubpodutils", "datahub pod utils", 0)
)

// NewStatus return pod status struct of datahub
func NewStatus(pod *CoreV1.Pod) *ApiResources.PodStatus {
	return &ApiResources.PodStatus{
		Message: pod.Status.Message,
		Reason:  pod.Status.Reason,
		Phase:   AlamedaEnum.PodPhaseEnumK8SToDatahub[pod.Status.Phase],
	}
}

// GetReplicasFromPod return number of replicas of pod
func GetReplicasFromPod(pod *CoreV1.Pod, client client.Client) int32 {
	getResource := AlamedaUtils.NewGetResource(client)

	for _, or := range pod.OwnerReferences {
		if or.Kind == AlamedaConsts.K8S_KIND_REPLICASET {
			rs, err := getResource.GetReplicaSet(pod.GetNamespace(), or.Name)
			if err == nil {
				return rs.Status.Replicas
			} else {
				scope.Errorf("Get replicaset for number of replicas failed due to %s", err.Error())
			}
		} else if or.Kind == AlamedaConsts.K8S_KIND_REPLICATIONCONTROLLER {
			rc, err := getResource.GetReplicationController(pod.GetNamespace(), or.Name)
			if err == nil {
				return rc.Status.Replicas
			} else {
				scope.Errorf("Get replicationcontroller for number of replicas failed due to %s", err.Error())
			}
		} else if or.Kind == AlamedaConsts.K8S_KIND_STATEFULSET {
			sts, err := getResource.GetStatefulSet(pod.GetNamespace(), or.Name)
			if err == nil {
				return sts.Status.Replicas
			} else {
				scope.Errorf("Get StatefulSet for number of replicas failed due to %s", err.Error())
			}
		}
	}
	return int32(-1)
}
