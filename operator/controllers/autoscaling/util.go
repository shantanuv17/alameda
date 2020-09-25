package autoscaling

import (
	corev1 "k8s.io/api/core/v1"
	operatorutils "prophetstor.com/alameda/operator/pkg/utils"
)

func GetTotalResourceFromContainers(containers []corev1.Container) corev1.ResourceRequirements {
	return operatorutils.GetTotalResourceFromContainers(containers)
}
