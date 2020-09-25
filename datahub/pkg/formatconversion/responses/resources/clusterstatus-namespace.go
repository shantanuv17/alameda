package resources

import (
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	"prophetstor.com/api/datahub/resources"
)

type NamespaceExtended struct {
	*types.Namespace
}

func (p *NamespaceExtended) ProduceNamespace() *resources.Namespace {
	namespace := &resources.Namespace{
		ObjectMeta: NewObjectMeta(p.ObjectMeta),
	}
	return namespace
}
