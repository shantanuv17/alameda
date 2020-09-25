package types

import (
	ApiPlannings "prophetstor.com/api/datahub/plannings"
)

// ContainerOperation defines container measurement operation of recommendation database
type ContainerPlanningsDAO interface {
	AddPodPlannings(in *ApiPlannings.CreatePodPlanningsRequest) error
	ListPodPlannings(in *ApiPlannings.ListPodPlanningsRequest) ([]*ApiPlannings.PodPlanning, error)
}
