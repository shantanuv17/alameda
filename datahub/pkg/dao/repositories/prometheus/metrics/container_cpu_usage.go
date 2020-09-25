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

type ContainerCpuUsageRepository struct {
	PrometheusConfig Prometheus.Config
}

func NewContainerCpuUsageRepositoryWithConfig(cfg Prometheus.Config) ContainerCpuUsageRepository {
	return ContainerCpuUsageRepository{PrometheusConfig: cfg}
}

func (c ContainerCpuUsageRepository) ListContainerCPUUsageMillicoresEntitiesByNamespaceAndPodNames(ctx context.Context, namespace string, podNames []string, options ...DBCommon.Option) ([]EntityPromthMetric.ContainerCPUUsageMillicoresEntity, error) {

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
		queryLabelsString += fmt.Sprintf(`%s != "",`, AvailableContainerCpuUsagePercentageLabelPodName[labelType])
		queryLabelsString += fmt.Sprintf(`%s != "POD",`, AvailableContainerCpuUsagePercentageLabelContainerName[labelType])
		queryLabelsString += fmt.Sprintf(`%s = "%s",`, ContainerCpuUsagePercentageLabelNamespace, namespace)
		names := ""
		for _, podName := range podNames {
			names += fmt.Sprintf("%s|", podName)
		}
		if names != "" {
			names = strings.TrimSuffix(names, "|")
			queryLabelsString += fmt.Sprintf(`%s =~ "%s",`, AvailableContainerCpuUsagePercentageLabelPodName[labelType], names)
		}

		queryLabelsString = strings.TrimSuffix(queryLabelsString, ",")
		queryExpression := fmt.Sprintf("%s{%s}", ContainerCpuUsagePercentageMetricName, queryLabelsString)
		stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
		queryExpression, err = Prometheus.WrapQueryExpression(queryExpression, opt.AggregateOverTimeFunc, stepTimeInSeconds)
		if err != nil {
			return nil, errors.Wrap(err, "list pod container cpu usage metric by namespaced name failed")
		}
		queryExpression = fmt.Sprintf(`1000 * %s`, queryExpression)
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
		return nil, errors.Wrap(err, "get prometheus entities")
	}
	cpuUsageEntities := make([]EntityPromthMetric.ContainerCPUUsageMillicoresEntity, len(entities))
	for i, e := range entities {
		cpuUsageEntities[i] = c.newContainerCPUUsageMillicoresEntity(e, labelType)
	}

	return cpuUsageEntities, nil
}

func (c ContainerCpuUsageRepository) newContainerCPUUsageMillicoresEntity(e Prometheus.Entity, labelType int) EntityPromthMetric.ContainerCPUUsageMillicoresEntity {

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

	return EntityPromthMetric.ContainerCPUUsageMillicoresEntity{
		PrometheusEntity: e,
		Namespace:        e.Labels[ContainerCpuUsagePercentageLabelNamespace],
		PodName:          e.Labels[AvailableContainerCpuUsagePercentageLabelPodName[labelType]],
		ContainerName:    e.Labels[AvailableContainerCpuUsagePercentageLabelContainerName[labelType]],
		Samples:          samples,
	}
}
