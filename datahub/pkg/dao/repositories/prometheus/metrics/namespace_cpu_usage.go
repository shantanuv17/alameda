package metrics

import (
	"context"
	"fmt"
	EntityPromthMetric "github.com/containers-ai/alameda/datahub/pkg/dao/entities/prometheus/metrics"
	FormatTypes "github.com/containers-ai/alameda/datahub/pkg/formatconversion/types"
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	Prometheus "github.com/containers-ai/alameda/pkg/database/prometheus"
	"github.com/pkg/errors"
	"strings"
	"time"
)

// NamespaceCPUUsageRepository Repository to access metric namespace_pod_name_container_name:container_cpu_usage_seconds_total:sum_rate from prometheus
type NamespaceCPUUsageRepository struct {
	PrometheusConfig Prometheus.Config
}

// NewNamespaceCPUUsageRepositoryWithConfig New namespace cpu usage millicores repository with prometheus configuration
func NewNamespaceCPUUsageRepositoryWithConfig(cfg Prometheus.Config) NamespaceCPUUsageRepository {
	return NamespaceCPUUsageRepository{PrometheusConfig: cfg}
}

func (c NamespaceCPUUsageRepository) ListNamespaceCPUUsageMillicoresEntitiesByNamespaceNames(ctx context.Context, namespaceNames []string, options ...DBCommon.Option) ([]EntityPromthMetric.NamespaceCPUUsageMillicoresEntity, error) {
	// Example of expression to query prometheus
	// 1000 * sum(namespace_pod_name_container_name:container_cpu_usage_seconds_total:sum_rate{pod_name!="",container_name!="POD",namespace=~"@n1"}) by (namespace)

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
		queryLabelsString := c.buildQueryLabelsStringByNamespaceNames(namespaceNames, labelType)
		queryExpression := fmt.Sprintf("%s{%s}", ContainerCpuUsagePercentageMetricName, queryLabelsString)
		stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
		queryExpression, err = Prometheus.WrapQueryExpression(queryExpression, opt.AggregateOverTimeFunc, stepTimeInSeconds)
		if err != nil {
			return nil, errors.Wrap(err, "wrap query expression failed")
		}
		queryExpression = fmt.Sprintf(`1000 * sum(%s) by (%s)`, queryExpression, ContainerCpuUsagePercentageLabelNamespace)

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
	foundMap := map[string]bool{}
	for _, name := range namespaceNames {
		foundMap[name] = false
	}
	namespaceCPUUsageMillicoresEntities := make([]EntityPromthMetric.NamespaceCPUUsageMillicoresEntity, len(entities))
	for i, entity := range entities {
		e := c.newNamespaceCPUUsageMillicoresEntity(entity)
		namespaceCPUUsageMillicoresEntities[i] = e
		foundMap[e.NamespaceName] = true
	}
	for name, exist := range foundMap {
		if !exist {
			namespaceCPUUsageMillicoresEntities = append(namespaceCPUUsageMillicoresEntities, EntityPromthMetric.NamespaceCPUUsageMillicoresEntity{
				NamespaceName: name,
			})
		}
	}

	return namespaceCPUUsageMillicoresEntities, nil
}

func (c NamespaceCPUUsageRepository) buildDefaultQueryLabelsString(labelType int) string {
	// sum(namespace_pod_name_container_name:container_cpu_usage_seconds_total:sum_rate{pod_name!="",container_name!="POD",namespace="@n1"})

	var queryLabelsString = ""

	queryLabelsString += fmt.Sprintf(`%s != "",`, AvailableContainerCpuUsagePercentageLabelPodName[labelType])
	queryLabelsString += fmt.Sprintf(`%s != "POD"`, AvailableContainerCpuUsagePercentageLabelContainerName[labelType])

	return queryLabelsString
}

func (c NamespaceCPUUsageRepository) buildQueryLabelsStringByNamespaceNames(namespaceNames []string, labelType int) string {
	var (
		queryLabelsString = c.buildDefaultQueryLabelsString(labelType)
	)

	names := ""
	for _, name := range namespaceNames {
		names += fmt.Sprintf("%s|", name)
	}
	if names != "" {
		names = strings.TrimSuffix(names, "|")
		queryLabelsString += fmt.Sprintf(`,%s =~ "%s"`, ContainerCpuUsagePercentageLabelNamespace, names)
	}

	return queryLabelsString
}

func (c NamespaceCPUUsageRepository) newNamespaceCPUUsageMillicoresEntity(e Prometheus.Entity) EntityPromthMetric.NamespaceCPUUsageMillicoresEntity {

	samples := make([]FormatTypes.Sample, len(e.Values))
	for i, value := range e.Values {
		samples[i] = FormatTypes.Sample{
			Timestamp: value.UnixTime,
			Value:     value.SampleValue,
		}
	}
	return EntityPromthMetric.NamespaceCPUUsageMillicoresEntity{
		PrometheusEntity: e,
		NamespaceName:    e.Labels[ContainerCpuUsagePercentageLabelNamespace],
		Samples:          samples,
	}
}
