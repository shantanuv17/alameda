package ca

import (
	"context"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	mahcinev1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SendExecutionTime(clusterUID string, executionTime *time.Time,
	k8sClient client.Client, datahubClient *datahubpkg.Client,
	machinesetNamespace, machinesetName, machineName string,
	isScalingUp bool) error {
	machineSetIns := mahcinev1beta1.MachineSet{}
	err := k8sClient.Get(context.Background(), client.ObjectKey{
		Namespace: machinesetNamespace,
		Name:      machinesetName,
	}, &machineSetIns)
	if err != nil {
		return err
	}

	machineSetEntities := []entities.ResourceClusterAutoscalerMachineset{}
	err = datahubClient.List(&machineSetEntities, datahubpkg.Option{
		Entity: entities.ResourceClusterAutoscalerMachineset{
			ClusterName: clusterUID,
			Namespace:   machinesetNamespace,
			Name:        machinesetName,
		},
		Fields: []string{"ClusterName", "Namespace", "Name"},
	})
	if err != nil {
		return err
	}

	if len(machineSetEntities) == 0 {
		return fmt.Errorf("machineset (%s/%s) not found in datahub",
			machinesetNamespace, machinesetName)
	}
	machineSetEntity := machineSetEntities[0]
	replicasTo := *machineSetIns.Spec.Replicas
	replicasFrom := replicasTo - 1
	if !isScalingUp {
		replicasFrom = replicasTo + 1
	}
	entities := []entities.ExecutionClusterAutoscalerMachineset{
		{
			Time:                  executionTime,
			ClusterName:           clusterUID,
			Namespace:             machinesetNamespace,
			Name:                  machinesetName,
			MachinegroupName:      machineSetEntity.MachinegroupName,
			MachinegroupNamespace: machineSetEntity.MachinegroupNamespace,
			ReplicasFrom:          replicasFrom,
			ReplicasTo:            replicasTo,
			NodeName:              machineName,
		},
	}

	return datahubClient.Create(&entities)
}
