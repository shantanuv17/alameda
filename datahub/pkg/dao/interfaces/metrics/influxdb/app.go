package influxdb

import (
	"context"

	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	DaoClusterStatusTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/types"
	DaoMetricTypes "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/metrics/types"
	RepoInfluxMetric "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/metrics"
	FormatEnum "github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	Utils "github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	"github.com/pkg/errors"
	"time"
)

type AppMetrics struct {
	InfluxDBConfig  InternalInflux.Config
	ApplicationDAO  DaoClusterStatusTypes.ApplicationDAO
	ApplicationRepo *RepoInfluxMetric.Application
}

func NewAppMetricsWithConfig(config InternalInflux.Config) DaoMetricTypes.AppMetricsDAO {
	appMetrics := AppMetrics{
		InfluxDBConfig:  config,
		ApplicationDAO:  influxdb.NewApplicationWithConfig(config),
		ApplicationRepo: RepoInfluxMetric.NewApplicationWithConfig(config),
	}
	return appMetrics
}

func (n AppMetrics) CreateMetrics(ctx context.Context, m DaoMetricTypes.AppMetricMap) error {
	// Write app cpu metrics
	appCPURepo := RepoInfluxMetric.NewApplicationCPURepositoryWithConfig(n.InfluxDBConfig)
	err := appCPURepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeCPUUsageSecondsPercentage))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create application cpu metrics failed")
	}

	// Write app memory metrics
	appMemoryRepo := RepoInfluxMetric.NewApplicationMemoryRepositoryWithConfig(n.InfluxDBConfig)
	err = appMemoryRepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeMemoryUsageBytes))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create application memory metrics failed")
	}
	return nil
}

func (n AppMetrics) ListMetrics(ctx context.Context, req DaoMetricTypes.ListAppMetricsRequest) (DaoMetricTypes.AppMetricMap, error) {
	metricMap := DaoMetricTypes.NewAppMetricMap()

	n.rebuildRequest(&req)

	applications, err := n.listApplications(req)
	if err != nil {
		return DaoMetricTypes.AppMetricMap{}, errors.Wrap(err, "list application metadatas from request failed")
	}

	// Read app cpu metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeCPUUsageSecondsPercentage) {
		cpuMetricMap, err := n.ApplicationRepo.GetMetricMap(FormatEnum.MetricTypeCPUUsageSecondsPercentage, applications, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get application cpu usage metric map failed")
		}
		for _, m := range cpuMetricMap.MetricMap {
			copyM := m
			metricMap.AddAppMetric(copyM)
		}

		/*appCPURepo := RepoInfluxMetric.NewApplicationCPURepositoryWithConfig(n.InfluxDBConfig)
		cpuMetricMap, err := appCPURepo.GetApplicationMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get application cpu usage metric map failed")
		}
		for _, m := range cpuMetricMap.MetricMap {
			copyM := m
			metricMap.AddAppMetric(copyM)
		}*/
	}

	// Read app memory metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeMemoryUsageBytes) {
		memoryMetricMap, err := n.ApplicationRepo.GetMetricMap(FormatEnum.MetricTypeMemoryUsageBytes, applications, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get application memory usage metric map failed")
		}
		for _, m := range memoryMetricMap.MetricMap {
			copyM := m
			metricMap.AddAppMetric(copyM)
		}

		/*appMemoryRepo := RepoInfluxMetric.NewApplicationMemoryRepositoryWithConfig(n.InfluxDBConfig)
		memoryMetricMap, err := appMemoryRepo.GetApplicationMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get application memory usage metric map failed")
		}
		for _, m := range memoryMetricMap.MetricMap {
			copyM := m
			metricMap.AddAppMetric(copyM)
		}*/
	}

	return metricMap, nil
}

func (n AppMetrics) listApplications(req DaoMetricTypes.ListAppMetricsRequest) ([]*DaoClusterStatusTypes.Application, error) {
	// Generate list resource application request
	listApplicationsReq := DaoClusterStatusTypes.NewListApplicationsRequest()
	for index := range req.ObjectMetas {
		applicationObjectMeta := DaoClusterStatusTypes.NewApplicationObjectMeta(&req.ObjectMetas[index], "")
		listApplicationsReq.ApplicationObjectMeta = append(listApplicationsReq.ApplicationObjectMeta, applicationObjectMeta)
	}

	apps, err := n.ApplicationDAO.ListApplications(listApplicationsReq)
	if err != nil {
		return nil, errors.Wrap(err, "list application metadatas failed")
	}

	return apps, nil
}

func (n AppMetrics) rebuildRequest(req *DaoMetricTypes.ListAppMetricsRequest) {
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
