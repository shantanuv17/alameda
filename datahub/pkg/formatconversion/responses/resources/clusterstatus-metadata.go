package resources

import (
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	"prophetstor.com/api/datahub/resources"
)

func NewObjectMeta(objectMeta *metadata.ObjectMeta) *resources.ObjectMeta {
	meta := resources.ObjectMeta{
		Name:        objectMeta.Name,
		Namespace:   objectMeta.Namespace,
		NodeName:    objectMeta.NodeName,
		ClusterName: objectMeta.ClusterName,
		Uid:         objectMeta.Uid,
	}
	return &meta
}
