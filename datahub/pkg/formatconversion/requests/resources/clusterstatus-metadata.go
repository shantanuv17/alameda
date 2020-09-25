package resources

import (
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	"prophetstor.com/api/datahub/resources"
)

func NewObjectMeta(objectMeta *resources.ObjectMeta) metadata.ObjectMeta {
	meta := metadata.ObjectMeta{}
	if objectMeta != nil {
		meta.Name = objectMeta.GetName()
		meta.Namespace = objectMeta.GetNamespace()
		meta.NodeName = objectMeta.GetNodeName()
		meta.ClusterName = objectMeta.GetClusterName()
		meta.Uid = objectMeta.GetUid()
	}
	return meta
}
