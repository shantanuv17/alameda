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

// ContainerMemoryUsageRepository Repository to access metric container_memory_usage_bytes from prometheus
type ContainerMemoryUsageRepository struct {
	PrometheusConfig Prometheus.Config
}

// NewContainerMemoryUsageRepositoryWithConfig New pod container memory usage bytes repository with prometheus configuration
func NewContainerMemoryUsageRepositoryWithConfig(cfg Prometheus.Config) ContainerMemoryUsageRepository {
	return ContainerMemoryUsageRepository{PrometheusConfig: cfg}
}

func (c ContainerMemoryUsageRepository) ListContainerMemoryUsageBytesEntitiesByNamespaceAndPodNames(ctx context.Context, namespace string, podNames []string, options ...DBCommon.Option) ([]EntityPromthMetric.ContainerMemoryUsageBytesEntity, error) {

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
		queryLabelsString := ""
		queryLabelsString += fmt.Sprintf(`%s != "",`, AvailableContainerMemoryUsageBytesLabelPodName[labelType])
		queryLabelsString += fmt.Sprintf(`%s != "",`, AvailableContainerMemoryUsageBytesLabelContainerName[labelType])
		queryLabelsString += fmt.Sprintf(`%s != "POD",`, AvailableContainerMemoryUsageBytesLabelContainerName[labelType])
		queryLabelsString += fmt.Sprintf(`%s = "%s",`, ContainerMemoryUsageBytesLabelNamespace, namespace)
		names := ""
		for _, podName := range podNames {
			names += fmt.Sprintf("%s|", podName)
		}
		if names != "" {
			names = strings.TrimSuffix(names, "|")
			queryLabelsString += fmt.Sprintf(`%s =~ "%s",`, AvailableContainerMemoryUsageBytesLabelPodName[labelType], names)
		}

		queryLabelsString = strings.TrimSuffix(queryLabelsString, ",")
		queryExpression := fmt.Sprintf("%s{%s}", ContainerMemoryUsageBytesMetricName, queryLabelsString)
		stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
		queryExpression, err = Prometheus.WrapQueryExpression(queryExpression, opt.AggregateOverTimeFunc, stepTimeInSeconds)
		if err != nil {
			return nil, errors.Wrap(err, "wrap query expression failed")
		}
		scope.Debugf("Query to prometheus: queryExpression: %+v, StartTime: %+v, EndTime: %+v, StepTime: %+v", queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
		response, err = prometheusClient.QueryRange(ctx, queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
		if err != nil {
			return nil, errors.Wrap(err, "query prometheus failed")
		} else if response.Status != Prometheus.StatusSuccess {
			return nil, errors.Errorf("receive error response from prometheus: %s", response.Error)
		}
		if len(response.Data.Result) != 0 {
			break
		}
	}

	entities, err := response.GetEntities()
	if err != nil {
		return nil, errors.Wrap(err, "get prometheus entities failed")
	}
	memoryUsageEntities := make([]EntityPromthMetric.ContainerMemoryUsageBytesEntity, len(entities))
	for i, e := range entities {
		memoryUsageEntities[i] = c.newContainerMemoryUsageBytesEntity(e, labelType)
	}

	return memoryUsageEntities, nil
}

func (c ContainerMemoryUsageRepository) newContainerMemoryUsageBytesEntity(e Prometheus.Entity, labelType int) EntityPromthMetric.ContainerMemoryUsageBytesEntity {

	var (
		samples []FormatTypes.Sample
	)

	samples = make([]FormatTypes.Sample, 0)

	for _, value := range e.Values {
		sample := FormatTypes.Sample{
			Timestamp: value.UnixTime,
			Value:     value.SampleValue,
		}
		samples = append(samples, sample)
	}

	return EntityPromthMetric.ContainerMemoryUsageBytesEntity{
		PrometheusEntity: e,
		Namespace:        e.Labels[ContainerMemoryUsageBytesLabelNamespace],
		PodName:          e.Labels[AvailableContainerMemoryUsageBytesLabelPodName[labelType]],
		ContainerName:    e.Labels[AvailableContainerMemoryUsageBytesLabelContainerName[labelType]],
		Samples:          samples,
	}
}
