package metrics

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	EntityPromthMetric "prophetstor.com/alameda/datahub/pkg/dao/entities/prometheus/metrics"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	Prometheus "prophetstor.com/alameda/pkg/database/prometheus"
)

// PodCPUUsageRepository Repository to access metric namespace_pod_name_container_name:container_cpu_usage_seconds_total:sum_rate from prometheus
type PodCPUUsageRepository struct {
	PrometheusConfig Prometheus.Config
}

// NewPodCPUUsageRepositoryWithConfig New pod cpu usage millicores repository with prometheus configuration
func NewPodCPUUsageRepositoryWithConfig(cfg Prometheus.Config) PodCPUUsageRepository {
	return PodCPUUsageRepository{PrometheusConfig: cfg}
}

func (c PodCPUUsageRepository) ListPodCPUUsageMillicoresEntitiesBySummingPodMetrics(ctx context.Context, namespace string, podNames []string, options ...DBCommon.Option) ([]EntityPromthMetric.PodCPUUsageMillicoresEntity, error) {
	// Example of expression to query prometheus
	// 1000 * sum(namespace_pod_name_container_name:container_cpu_usage_seconds_total:sum_rate{pod_name!="",container_name!="POD",namespace="@n1",pod_name=~"@p1|@p2"})

	var (
		response  Prometheus.Response
		labelType = 0
		err       error
	)

	prometheusClient, err := Prometheus.NewClient(&c.PrometheusConfig)
	if err != nil {
		return nil, errors.Wrap(err, "new prometheus client failed")
	}

	opt := DBCommon.NewDefaultOptions()
	for _, option := range options {
		option(&opt)
	}

	for labelType = 0; labelType < AvailableLabelType; labelType++ {
		queryLabelsString := c.buildDefaultQueryLabelsString(labelType)
		queryLabelsString += fmt.Sprintf(`, %s = "%s"`, ContainerCpuUsagePercentageLabelNamespace, namespace)
		names := ""
		for _, name := range podNames {
			names += fmt.Sprintf("%s|", name)
		}
		if names != "" {
			names = strings.TrimSuffix(names, "|")
			queryLabelsString += fmt.Sprintf(`,%s =~ "%s"`, AvailableContainerCpuUsagePercentageLabelPodName[labelType], names)
		}
		queryExpression := fmt.Sprintf(`%s{%s}`, ContainerCpuUsagePercentageMetricName, queryLabelsString)
		stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
		queryExpression, err = Prometheus.WrapQueryExpression(queryExpression, opt.AggregateOverTimeFunc, stepTimeInSeconds)
		if err != nil {
			return nil, errors.Wrap(err, "wrap query expression failed")
		}
		queryExpression = fmt.Sprintf(`1000 * sum(%s)`, queryExpression)

		scope.Debugf("Query to prometheus: queryExpression: %+v, StartTime: %+v, EndTime: %+v, StepTime: %+v", queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
		response, err = prometheusClient.QueryRange(ctx, queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
		if err != nil {
			return nil, errors.Wrap(err, "query prometheus failed")
		} else if response.Status != Prometheus.StatusSuccess {
			return nil, errors.Errorf("query prometheus failed: receive error response from prometheus: %s", response.Error)
		}
		if len(response.Data.Result) != 0 {
			break
		}
	}

	entities, err := response.GetEntities()
	if err != nil {
		return nil, errors.Wrap(err, "get prometheus response entities failed")
	}
	podCPUUsageMillicoresEntities := make([]EntityPromthMetric.PodCPUUsageMillicoresEntity, len(entities))
	for i, entity := range entities {
		e := c.newPodCPUUsageMillicoresEntity(entity, labelType)
		podCPUUsageMillicoresEntities[i] = e
	}

	return podCPUUsageMillicoresEntities, nil
}

func (c PodCPUUsageRepository) buildDefaultQueryLabelsString(labelType int) string {
	// 1000 * sum( {pod_name!="",container_name!="POD",namespace="@n1",pod_name=~"@p1|@p2"})

	var queryLabelsString = ""
	queryLabelsString += fmt.Sprintf(`%s != "",`, AvailableContainerCpuUsagePercentageLabelPodName[labelType])
	queryLabelsString += fmt.Sprintf(`%s != "POD"`, AvailableContainerCpuUsagePercentageLabelContainerName[labelType])
	return queryLabelsString
}

func (c PodCPUUsageRepository) newPodCPUUsageMillicoresEntity(e Prometheus.Entity, labelType int) EntityPromthMetric.PodCPUUsageMillicoresEntity {

	samples := make([]FormatTypes.Sample, len(e.Values))
	for i, value := range e.Values {
		samples[i] = FormatTypes.Sample{
			Timestamp: value.UnixTime,
			Value:     value.SampleValue,
		}
	}
	return EntityPromthMetric.PodCPUUsageMillicoresEntity{
		NamespaceName: e.Labels[ContainerCpuUsagePercentageLabelNamespace],
		PodName:       e.Labels[AvailableContainerCpuUsagePercentageLabelPodName[labelType]],
		Samples:       samples,
	}
}
