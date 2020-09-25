package types

import (
	"prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	"prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/alameda/pkg/database/influxdb"
)

type NamespaceDAO interface {
	CreateNamespaces([]*Namespace) error
	ListNamespaces(*ListNamespacesRequest) ([]*Namespace, error)
	DeleteNamespaces(*DeleteNamespacesRequest) error
}

type Namespace struct {
	ObjectMeta *metadata.ObjectMeta
	Value      string
}

type ListNamespacesRequest struct {
	common.QueryCondition
	ObjectMeta []*metadata.ObjectMeta
}

type DeleteNamespacesRequest struct {
	ObjectMeta []*metadata.ObjectMeta
}

func NewNamespace(entity *clusterstatus.NamespaceEntity) *Namespace {
	namespace := Namespace{}
	namespace.ObjectMeta = &metadata.ObjectMeta{}
	namespace.ObjectMeta.Name = entity.Name
	namespace.ObjectMeta.ClusterName = entity.ClusterName
	namespace.ObjectMeta.Uid = entity.Uid
	namespace.Value = entity.Value
	return &namespace
}

func NewListNamespacesRequest() *ListNamespacesRequest {
	request := ListNamespacesRequest{}
	request.ObjectMeta = make([]*metadata.ObjectMeta, 0)
	return &request
}

func NewDeleteNamespacesRequest() *DeleteNamespacesRequest {
	request := DeleteNamespacesRequest{}
	request.ObjectMeta = make([]*metadata.ObjectMeta, 0)
	return &request
}

func (p *Namespace) BuildEntity() *clusterstatus.NamespaceEntity {
	entity := clusterstatus.NamespaceEntity{}

	entity.Time = influxdb.ZeroTime
	entity.Value = p.Value

	if p.ObjectMeta != nil {
		entity.Name = p.ObjectMeta.Name
		entity.ClusterName = p.ObjectMeta.ClusterName
		entity.Uid = p.ObjectMeta.Uid
	}

	return &entity
}
