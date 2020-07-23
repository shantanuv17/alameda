package metrics

import (
	"context"
	"fmt"
	"strings"
	"time"

	EntityPromthMetric "github.com/containers-ai/alameda/datahub/pkg/dao/entities/prometheus/metrics"
	FormatTypes "github.com/containers-ai/alameda/datahub/pkg/formatconversion/types"
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	Prometheus "github.com/containers-ai/alameda/pkg/database/prometheus"
	"github.com/pkg/errors"
)

// PodMemoryUsageRepository Repository to access metric container_memory_usage_bytes from prometheus
type PodMemoryUsageRepository struct {
	PrometheusConfig Prometheus.Config
}

// NewPodMemoryUsageRepositoryWithConfig New pod memory usage bytes repository with prometheus configuration
func NewPodMemoryUsageRepositoryWithConfig(cfg Prometheus.Config) PodMemoryUsageRepository {
	return PodMemoryUsageRepository{PrometheusConfig: cfg}
}

func (c PodMemoryUsageRepository) ListPodMemoryUsageBytesEntityBySummingPodMetrics(ctx context.Context, namespace string, podNames []string, options ...DBCommon.Option) ([]EntityPromthMetric.PodMemoryUsageBytesEntity, error) {
	// Example of expression to query prometheus
	// sum(container_memory_usage_bytes{pod_name!="",container_name!="",container_name!="POD",namespace="@n1",pod_name=~"@p1|@p2"})

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
		queryLabelsString += fmt.Sprintf(`, %s = "%s"`, ContainerMemoryUsageBytesLabelNamespace, namespace)
		names := ""
		for _, name := range podNames {
			names += fmt.Sprintf("%s|", name)
		}
		if names != "" {
			names = strings.TrimSuffix(names, "|")
			queryLabelsString += fmt.Sprintf(`,%s =~ "%s"`, AvailableContainerCpuUsagePercentageLabelPodName[labelType], names)
		}
		queryExpression := fmt.Sprintf(`%s{%s}`, ContainerMemoryUsageBytesMetricName, queryLabelsString)
		stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
		queryExpression, err = Prometheus.WrapQueryExpression(queryExpression, opt.AggregateOverTimeFunc, stepTimeInSeconds)
		if err != nil {
			return nil, errors.Wrap(err, "wrap query expression failed")
		}
		queryExpression = fmt.Sprintf(`sum(%s)`, queryExpression)

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
	podMemoryUsageBytesEntities := make([]EntityPromthMetric.PodMemoryUsageBytesEntity, len(entities))
	for i, entity := range entities {
		e := c.newPodMemoryUsageBytesEntity(entity, labelType)
		podMemoryUsageBytesEntities[i] = e
	}

	return podMemoryUsageBytesEntities, nil
}

func (c PodMemoryUsageRepository) buildDefaultQueryLabelsString(labelType int) string {
	// sum(container_memory_usage_bytes{pod_name!="",container_name!="",container_name!="POD",namespace="@n1",pod_name=~"@p1|@p2"})

	var queryLabelsString = ""
	queryLabelsString += fmt.Sprintf(`%s != "",`, AvailableContainerMemoryUsageBytesLabelPodName[labelType])
	queryLabelsString += fmt.Sprintf(`%s != "",`, AvailableContainerMemoryUsageBytesLabelContainerName[labelType])
	queryLabelsString += fmt.Sprintf(`%s != "POD"`, AvailableContainerMemoryUsageBytesLabelContainerName[labelType])
	return queryLabelsString
}

func (c PodMemoryUsageRepository) newPodMemoryUsageBytesEntity(e Prometheus.Entity, labelType int) EntityPromthMetric.PodMemoryUsageBytesEntity {

	samples := make([]FormatTypes.Sample, len(e.Values))
	for i, value := range e.Values {
		samples[i] = FormatTypes.Sample{
			Timestamp: value.UnixTime,
			Value:     value.SampleValue,
		}
	}
	return EntityPromthMetric.PodMemoryUsageBytesEntity{
		NamespaceName: e.Labels[ContainerCpuUsagePercentageLabelNamespace],
		PodName:       e.Labels[AvailableContainerCpuUsagePercentageLabelPodName[labelType]],
		Samples:       samples,
	}
}
