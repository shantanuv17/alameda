package types

import (
	"prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	"prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/alameda/pkg/database/influxdb"
)

type ControllerDAO interface {
	CreateControllers([]*Controller) error
	ListControllers(*ListControllersRequest) ([]*Controller, error)
	DeleteControllers(*DeleteControllersRequest) error
}

type Controller struct {
	ObjectMeta            *metadata.ObjectMeta
	Kind                  string
	Replicas              int32
	SpecReplicas          int32
	AlamedaControllerSpec *AlamedaControllerSpec
}

type ListControllersRequest struct {
	common.QueryCondition
	ControllerObjectMeta []*ControllerObjectMeta
}

type DeleteControllersRequest struct {
	ControllerObjectMeta []*ControllerObjectMeta
}

type ControllerObjectMeta struct {
	ObjectMeta    *metadata.ObjectMeta
	AlamedaScaler *metadata.ObjectMeta
	Kind          string // Valid values: DEPLOYMENT, DEPLOYMENTCONFIG, STATEFULSET
	ScalingTool   string // Valid values: NONE, VPA, HPA
}

type AlamedaControllerSpec struct {
	AlamedaScaler   *metadata.ObjectMeta
	ScalingTool     string
	Policy          string
	MinReplicas     int32
	MaxReplicas     int32
	EnableExecution bool
}

func NewController(entity *clusterstatus.ControllerEntity) *Controller {
	controller := Controller{}
	controller.ObjectMeta = &metadata.ObjectMeta{}
	controller.ObjectMeta.Name = entity.Name
	controller.ObjectMeta.Namespace = entity.Namespace
	controller.ObjectMeta.ClusterName = entity.ClusterName
	controller.ObjectMeta.Uid = entity.Uid
	controller.Kind = entity.Kind
	controller.Replicas = entity.Replicas
	controller.SpecReplicas = entity.SpecReplicas
	controller.AlamedaControllerSpec = NewAlamedaControllerSpec(entity)
	return &controller
}

func NewListControllersRequest() *ListControllersRequest {
	request := ListControllersRequest{}
	request.ControllerObjectMeta = make([]*ControllerObjectMeta, 0)
	return &request
}

func NewDeleteControllersRequest() *DeleteControllersRequest {
	request := DeleteControllersRequest{}
	request.ControllerObjectMeta = make([]*ControllerObjectMeta, 0)
	return &request
}

func NewControllerObjectMeta(objectMeta, alamedaScaler *metadata.ObjectMeta, kind, scalingTool string) *ControllerObjectMeta {
	controllerObjectMeta := ControllerObjectMeta{}
	controllerObjectMeta.ObjectMeta = objectMeta
	controllerObjectMeta.AlamedaScaler = alamedaScaler
	controllerObjectMeta.Kind = kind
	controllerObjectMeta.ScalingTool = scalingTool
	return &controllerObjectMeta
}

func NewAlamedaControllerSpec(entity *clusterstatus.ControllerEntity) *AlamedaControllerSpec {
	spec := AlamedaControllerSpec{}
	spec.AlamedaScaler = &metadata.ObjectMeta{}
	spec.AlamedaScaler.Name = entity.AlamedaSpecScalerName
	spec.AlamedaScaler.Namespace = entity.AlamedaSpecScalerNamespace
	spec.ScalingTool = entity.AlamedaSpecScalingTool
	spec.Policy = entity.AlamedaSpecPolicy
	spec.MinReplicas = entity.MinReplicas
	spec.MaxReplicas = entity.MaxReplicas
	spec.EnableExecution = entity.AlamedaSpecEnableExecution
	return &spec
}

func (p *Controller) BuildEntity() *clusterstatus.ControllerEntity {
	entity := clusterstatus.ControllerEntity{}

	entity.Time = influxdb.ZeroTime
	entity.Kind = p.Kind
	entity.Replicas = p.Replicas
	entity.SpecReplicas = p.SpecReplicas

	if p.ObjectMeta != nil {
		entity.Name = p.ObjectMeta.Name
		entity.Namespace = p.ObjectMeta.Namespace
		entity.ClusterName = p.ObjectMeta.ClusterName
		entity.Uid = p.ObjectMeta.Uid
	}

	if p.AlamedaControllerSpec != nil {
		if p.AlamedaControllerSpec.AlamedaScaler != nil {
			entity.AlamedaSpecScalerName = p.AlamedaControllerSpec.AlamedaScaler.Name
			entity.AlamedaSpecScalerNamespace = p.AlamedaControllerSpec.AlamedaScaler.Namespace
		}
		entity.AlamedaSpecScalingTool = p.AlamedaControllerSpec.ScalingTool
		entity.AlamedaSpecPolicy = p.AlamedaControllerSpec.Policy
		entity.MinReplicas = p.AlamedaControllerSpec.MinReplicas
		entity.MaxReplicas = p.AlamedaControllerSpec.MaxReplicas
		entity.AlamedaSpecEnableExecution = p.AlamedaControllerSpec.EnableExecution
	}

	return &entity
}
