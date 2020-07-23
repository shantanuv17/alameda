package influxdb

import (
	DaoClusterTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	RepoInfluxCluster "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	"github.com/containers-ai/alameda/pkg/database/common"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	Log "github.com/containers-ai/alameda/pkg/utils/log"
)

var (
	scope = Log.RegisterScope("dao_influxdb_metric_implement", "dao implement", 0)
)

func ListPods(config InfluxDB.Config, request *DaoClusterTypes.ListPodsRequest) ([]*DaoClusterTypes.Pod, error) {
	podRepo := RepoInfluxCluster.NewPodRepository(config)
	return podRepo.ListPods(request)
}

func ListControllers(config InfluxDB.Config, request *DaoClusterTypes.ListControllersRequest) ([]*DaoClusterTypes.Controller, error) {
	controllerRepo := RepoInfluxCluster.NewControllerRepository(config)
	return controllerRepo.ListControllers(request)
}

func NewFunction(functionType common.FunctionType) *common.Function {
	function := common.Function{
		Type:   functionType,
		Target: "value",
	}
	return &function
}
