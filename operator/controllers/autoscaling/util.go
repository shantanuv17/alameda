package autoscaling

import (
	operatorutils "github.com/containers-ai/alameda/operator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
)

func GetTotalResourceFromContainers(containers []corev1.Container) corev1.ResourceRequirements {
	return operatorutils.GetTotalResourceFromContainers(containers)
}
