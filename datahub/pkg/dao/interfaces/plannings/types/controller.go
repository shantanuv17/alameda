package types

import (
	ApiPlannings "prophetstor.com/api/datahub/plannings"
)

// ContainerOperation defines container measurement operation of recommendation database
type ControllerPlanningsDAO interface {
	AddControllerPlannings([]*ApiPlannings.ControllerPlanning) error
	ListControllerPlannings(in *ApiPlannings.ListControllerPlanningsRequest) ([]*ApiPlannings.ControllerPlanning, error)
}
