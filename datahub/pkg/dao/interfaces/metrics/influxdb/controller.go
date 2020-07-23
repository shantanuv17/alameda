package influxdb

import (
	"context"

	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	DaoClusterStatusTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	RepoInfluxMetric "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/metrics"
	FormatEnum "github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	Utils "github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/pkg/database/common"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type ControllerMetrics struct {
	InfluxDBConfig InfluxDB.Config
	ControllerDAO  DaoClusterStatusTypes.ControllerDAO
	ControllerRepo *RepoInfluxMetric.Controller
}

func NewControllerMetricsWithConfig(config InfluxDB.Config) DaoMetricTypes.ControllerMetricsDAO {
	controllerMetrics := ControllerMetrics{
		InfluxDBConfig: config,
		ControllerDAO:  influxdb.NewControllerWithConfig(config),
		ControllerRepo: RepoInfluxMetric.NewControllerWithConfig(config),
	}
	return controllerMetrics
}

func (n ControllerMetrics) CreateMetrics(ctx context.Context, m DaoMetricTypes.ControllerMetricMap) error {
	// Write controller cpu metrics
	controllerCPURepo := RepoInfluxMetric.NewControllerCPURepositoryWithConfig(n.InfluxDBConfig)
	err := controllerCPURepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeCPUUsageSecondsPercentage))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create controller cpu metrics failed")
	}

	// Write controller memory metrics
	memoryRepo := RepoInfluxMetric.NewControllerMemoryRepositoryWithConfig(n.InfluxDBConfig)
	err = memoryRepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeMemoryUsageBytes))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create controller memory metrics failed")
	}

	return nil
}

func (n ControllerMetrics) ListMetrics(ctx context.Context, req DaoMetricTypes.ListControllerMetricsRequest) (DaoMetricTypes.ControllerMetricMap, error) {
	metricMap := DaoMetricTypes.NewControllerMetricMap()

	n.rebuildRequest(&req)

	controllers, err := n.listControllers(req)
	if err != nil {
		return DaoMetricTypes.ControllerMetricMap{}, errors.Wrap(err, "list controller metadatas from request failed")
	}

	// Read controller cpu metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeCPUUsageSecondsPercentage) {
		cpuMetricMap, err := n.ControllerRepo.GetMetricMap(FormatEnum.MetricTypeCPUUsageSecondsPercentage, controllers, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get controller cpu usage metric map failed")
		}
		for _, m := range cpuMetricMap.MetricMap {
			copyM := m
			metricMap.AddControllerMetric(copyM)
		}

		/*controllerCPURepo := RepoInfluxMetric.NewControllerCPURepositoryWithConfig(n.InfluxDBConfig)
		cpuMetricMap, err := controllerCPURepo.GetControllerMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get controller cpu usage metric map failed")
		}
		for _, m := range cpuMetricMap.MetricMap {
			copyM := m
			metricMap.AddControllerMetric(copyM)
		}*/
	}

	// Read controller memory metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeMemoryUsageBytes) {
		memoryMetricMap, err := n.ControllerRepo.GetMetricMap(FormatEnum.MetricTypeMemoryUsageBytes, controllers, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get controller memory usage metric map failed")
		}
		for _, m := range memoryMetricMap.MetricMap {
			copyM := m
			metricMap.AddControllerMetric(copyM)
		}

		/*controllerMemoryRepo := RepoInfluxMetric.NewControllerMemoryRepositoryWithConfig(n.InfluxDBConfig)
		memoryMetricMap, err := controllerMemoryRepo.GetControllerMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get controller memory usage metric map failed")
		}
		for _, m := range memoryMetricMap.MetricMap {
			copyM := m
			metricMap.AddControllerMetric(copyM)
		}*/
	}

	return metricMap, nil
}

func (n ControllerMetrics) listControllers(req DaoMetricTypes.ListControllerMetricsRequest) ([]*DaoClusterStatusTypes.Controller, error) {
	// Generate list resource controllers request
	listControllersReq := DaoClusterStatusTypes.NewListControllersRequest()
	for _, objectMeta := range req.ObjectMetas {
		controllerObjectMeta := DaoClusterStatusTypes.NewControllerObjectMeta(&objectMeta, nil, strings.ToUpper(req.Kind), "")
		listControllersReq.ControllerObjectMeta = append(listControllersReq.ControllerObjectMeta, controllerObjectMeta)

	}

	if len(listControllersReq.ControllerObjectMeta) == 0 {
		controllerObjectMeta := DaoClusterStatusTypes.NewControllerObjectMeta(nil, nil, strings.ToUpper(req.Kind), "")
		listControllersReq.ControllerObjectMeta = append(listControllersReq.ControllerObjectMeta, controllerObjectMeta)
	}

	controllers, err := n.ControllerDAO.ListControllers(listControllersReq)
	if err != nil {
		return nil, errors.Wrap(err, "list controller metadatas failed")
	}

	return controllers, nil
}

func (n ControllerMetrics) rebuildRequest(req *DaoMetricTypes.ListControllerMetricsRequest) {
	if req.Kind == "KIND_UNDEFINED" {
		req.Kind = ""
	}

	req.Selects = append(req.Selects, "value")
	if req.AggregateOverTimeFunction != common.None {
		req.Function = NewFunction(common.FunctionTypeMap[req.AggregateOverTimeFunction])
		req.AggregateOverTimeFunction = common.None
	}

	step := time.Duration(15) * time.Second
	req.QueryCondition.SubQuery.StepTime = &step
	req.QueryCondition.SubQuery.Selects = append(req.QueryCondition.SubQuery.Selects, "value")
	req.QueryCondition.SubQuery.Groups = append(req.QueryCondition.SubQuery.Groups, "cluster_name")
	req.QueryCondition.SubQuery.AggregateOverTimeFunction = common.None
	req.QueryCondition.SubQuery.Function = NewFunction(common.FunctionSum)
}
