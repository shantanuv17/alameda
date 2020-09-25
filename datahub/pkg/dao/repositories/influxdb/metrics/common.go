package metrics

import (
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	"prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	"prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	"prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/alameda/pkg/database/influxdb"
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
)

var aggregateFuncToInfluxDBFunc = map[common.AggregateFunction]influxdb.Functions{
	common.None:        influxdb.Last,
	common.MaxOverTime: influxdb.Max,
	common.AvgOverTime: influxdb.Mean,
}

var metricTypeMapTable = map[enumconv.MetricType]schemas.MetricType{
	enumconv.MetricTypeCPUUsageSecondsPercentage: schemas.CPUMilliCoresUsage,
	enumconv.MetricTypeMemoryUsageBytes:          schemas.MemoryBytesUsage,
}

func ListPodsByController(config influxdb.Config, controller *types.Controller) ([]*types.Pod, error) {
	objectMeta := metadata.ObjectMeta{
		Namespace:   controller.ObjectMeta.Namespace,
		ClusterName: controller.ObjectMeta.ClusterName,
	}

	request := types.NewListPodsRequest()
	request.ObjectMeta = append(request.ObjectMeta, &objectMeta)
	request.TopControllerName = controller.ObjectMeta.Name
	request.AlamedaScalerName = controller.AlamedaControllerSpec.AlamedaScaler.Name
	request.Kind = controller.Kind
	request.ScalingTool = controller.AlamedaControllerSpec.ScalingTool

	podRepo := clusterstatus.NewPodRepository(config)
	return podRepo.ListPods(request)
}

func ListPodsByApplication(config influxdb.Config, application *types.Application) ([]*types.Pod, error) {
	objectMeta := metadata.ObjectMeta{
		ClusterName: application.ObjectMeta.ClusterName,
	}

	request := types.NewListPodsRequest()
	request.ObjectMeta = append(request.ObjectMeta, &objectMeta)
	request.AlamedaScalerName = application.ObjectMeta.Name
	request.AlamedaScalerNamespace = application.ObjectMeta.Namespace

	podRepo := clusterstatus.NewPodRepository(config)
	return podRepo.ListPods(request)
}

func ListNodesByCluster(config influxdb.Config, cluster *types.Cluster) ([]*types.Node, error) {
	objectMeta := metadata.ObjectMeta{
		ClusterName: cluster.ObjectMeta.Name,
	}

	request := types.NewListNodesRequest()
	request.ObjectMeta = append(request.ObjectMeta, &objectMeta)

	nodeRepo := clusterstatus.NewNodeRepository(config)
	return nodeRepo.ListNodes(request)
}
