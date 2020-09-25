package resources

import (
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	"prophetstor.com/api/datahub/resources"
)

type NodeExtended struct {
	*types.Node
}

func (p *NodeExtended) ProduceNode() *resources.Node {
	node := resources.Node{}
	node.ObjectMeta = NewObjectMeta(p.ObjectMeta)
	node.StartTime = p.CreateTime
	node.Capacity = NewCapacity(p.Capacity)
	node.AlamedaNodeSpec = NewAlamedaNodeSpec(p.AlamedaNodeSpec)
	return &node
}
