package node

import (
	"context"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entities"
	nodeinfo "github.com/containers-ai/alameda/operator/pkg/nodeinfo"
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	k8sutils "github.com/containers-ai/alameda/pkg/utils/kubernetes"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SyncWithDatahub(client client.Client, datahubClient *datahubpkg.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	nodeList := corev1.NodeList{}
	if err := client.List(ctx, &nodeList); err != nil {
		return errors.Errorf(
			"Sync nodes with datahub failed due to list nodes from cluster failed: %s", err.Error())
	}

	clusterUID, err := k8sutils.GetClusterUID(client)
	if err != nil {
		return errors.Wrap(err, "get cluster uid failed")
	}

	datahubNodeRepo := NewNodeRepository(datahubClient, clusterUID)
	nodes := make([]entities.ResourceClusterStatusNode, len(nodeList.Items))
	for i, node := range nodeList.Items {
		nodeInfo, err := nodeinfo.NewNodeInfo(node, clusterUID)
		if err != nil {
			return errors.Wrap(err, "new nodeInfo failed")
		}
		nodes[i] = nodeInfo
	}

	if err := datahubNodeRepo.CreateNodes(nodes); err != nil {
		return fmt.Errorf(
			"Sync nodes with datahub failed due to register node failed: %s", err.Error())
	}

	// Clean up unexisting nodes from Datahub
	existingNodeMap := make(map[string]bool)
	for _, node := range nodeList.Items {
		existingNodeMap[node.Name] = true
	}

	nodesFromDatahub, err := datahubNodeRepo.ListNodes()
	if err != nil {
		return fmt.Errorf(
			"Sync nodes with datahub failed due to list nodes from datahub failed: %s", err.Error())
	}
	nodesNeedDeleting := make([]*datahub_resources.Node, 0)
	for _, n := range nodesFromDatahub {
		if _, exist := existingNodeMap[n.ObjectMeta.GetName()]; exist {
			continue
		}
		nodesNeedDeleting = append(nodesNeedDeleting, n)
	}
	if len(nodesNeedDeleting) > 0 {
		err = datahubNodeRepo.DeleteNodes(nodesNeedDeleting)
		if err != nil {
			return errors.Wrap(err, "delete nodes from Datahub failed")
		}
	}

	return nil
}
