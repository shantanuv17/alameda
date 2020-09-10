package keycodes

import (
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	"github.com/containers-ai/alameda/datahub/pkg/kubernetes/metadata"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
)

func getClustersFromDB(dao types.ClusterDAO) ([]*types.Cluster, error) {
	req := types.ListClustersRequest{}
	clsts, err := dao.ListClusters(&req)
	return clsts, err
}

func getClusterNodesFromDB(dao types.NodeDAO, clusterName string) ([]*types.Node, error) {
	objmeta := make([]*metadata.ObjectMeta, 0)
	objmeta = append(objmeta, &metadata.ObjectMeta{ClusterName: clusterName})
	req := types.ListNodesRequest{ObjectMeta: objmeta}
	ns, err := dao.ListNodes(&req)
	return ns, err
}

func GetAlamedaClusterCPUs(influxCfg *InfluxDB.Config) (int, error) {
	numCPU := 0
	clusterName := ""

	dbNode := influxdb.NewNodeWithConfig(*influxCfg)

	nodes, err := getClusterNodesFromDB(dbNode, clusterName)
	if err != nil {
		scope.Errorf("Failed to get cluster [%s] node info: %s", clusterName, err.Error())
		return numCPU, err
	} else {
		scope.Infof("Number of node: %d", len(nodes))
		for _, node := range nodes {
			if node != nil {
				scope.Infof("  node: %s (cluster: %s), CPU Cores: %d", node.ObjectMeta.Name,
					node.ObjectMeta.ClusterName, node.Capacity.CpuCores)
				numCPU += int(node.Capacity.CpuCores)
			}
		}
	}
	return numCPU, nil
}
