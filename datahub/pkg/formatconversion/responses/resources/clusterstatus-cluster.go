package resources

import (
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	"prophetstor.com/api/datahub/resources"
)

type ClusterExtended struct {
	*types.Cluster
}

func (p *ClusterExtended) ProduceCluster() *resources.Cluster {
	cluster := &resources.Cluster{
		ObjectMeta: NewObjectMeta(p.ObjectMeta),
	}
	return cluster
}
