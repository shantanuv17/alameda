package metrics

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	EntityPromthMetric "prophetstor.com/alameda/datahub/pkg/dao/entities/prometheus/metrics"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	Prometheus "prophetstor.com/alameda/pkg/database/prometheus"
	"time"
)

// NodeCPUUsageRepository Repository to access metric node:node_cpu_utilisation:avg1m from prometheus
type NodeCPUUsageRepository struct {
	PrometheusConfig Prometheus.Config
}

// NewNodeCPUUsageRepositoryWithConfig New node cpu usage   repository with prometheus configuration
func NewNodeCPUUsageRepositoryWithConfig(cfg Prometheus.Config) NodeCPUUsageRepository {
	return NodeCPUUsageRepository{PrometheusConfig: cfg}
}

func (n NodeCPUUsageRepository) ListNodeCPUUsageMillicoresEntitiesByNodeNames(ctx context.Context, nodeNames []string, options ...DBCommon.Option) ([]EntityPromthMetric.NodeCPUUsageMillicoresEntity, error) {

	prometheusClient, err := Prometheus.NewClient(&n.PrometheusConfig)
	if err != nil {
		return nil, errors.Wrap(err, "new prometheus client failed")
	}

	opt := DBCommon.NewDefaultOptions()
	for _, option := range options {
		option(&opt)
	}

	names := ""
	for _, nodeName := range nodeNames {
		names += fmt.Sprintf("%s|", nodeName)
	}
	names = strings.TrimSuffix(names, "|")

	stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))

	queryLabelsStringSum := ""
	if names != "" {
		queryLabelsStringSum = fmt.Sprintf(`%s =~ "%s"`, NodeCpuUsagePercentageLabelNode, names)
	}
	queryExpressionSum := fmt.Sprintf("%s{%s}", NodeCpuUsagePercentageMetricNameSum, queryLabelsStringSum)
	queryExpressionSum, err = Prometheus.WrapQueryExpression(queryExpressionSum, opt.AggregateOverTimeFunc, stepTimeInSeconds)
	if err != nil {
		return nil, errors.Wrap(err, "wrap query expression failed")
	}

	queryLabelsStringAvg := ""
	if names != "" {
		queryLabelsStringAvg = fmt.Sprintf(`%s =~ "%s"`, NodeCpuUsagePercentageLabelNode, names)
	}
	queryExpressionAvg := fmt.Sprintf("%s{%s}", NodeCpuUsagePercentageMetricNameAvg, queryLabelsStringAvg)
	queryExpressionAvg, err = Prometheus.WrapQueryExpression(queryExpressionAvg, opt.AggregateOverTimeFunc, stepTimeInSeconds)
	if err != nil {
		return nil, errors.Wrap(err, "wrap query expression failed")
	}

	queryExpression := fmt.Sprintf("1000 * %s * %s", queryExpressionSum, queryExpressionAvg)
	scope.Debugf("Query to prometheus: queryExpression: %+v, StartTime: %+v, EndTime: %+v, StepTime: %+v", queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
	response, err := prometheusClient.QueryRange(ctx, queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
	if err != nil {
		return nil, errors.Wrap(err, "query prometheus failed")
	} else if response.Status != Prometheus.StatusSuccess {
		return nil, errors.Errorf("receive error response from prometheus: %s", response.Error)
	}

	entities, err := response.GetEntities()
	if err != nil {
		return nil, errors.Wrap(err, "get prometheus entities failed")
	}
	cpuUsageEntities := make([]EntityPromthMetric.NodeCPUUsageMillicoresEntity, len(entities))
	for i, e := range entities {
		cpuUsageEntities[i] = n.newNodeCPUUsageMillicoresEntity(e)
	}

	return cpuUsageEntities, nil
}

func (n NodeCPUUsageRepository) ListSumOfNodeCPUUsageMillicoresByNodeNames(ctx context.Context, nodeNames []string, options ...DBCommon.Option) ([]EntityPromthMetric.NodeCPUUsageMillicoresEntity, error) {
	// Example of expression to query prometheus
	// 1000 * sum(node:node_cpu_utilisation:avg1m{node=~"@n1|@n2"} * node:node_num_cpu:sum{node=~"@n1|@n2"})

	prometheusClient, err := Prometheus.NewClient(&n.PrometheusConfig)
	if err != nil {
		return nil, errors.Wrap(err, "new prometheus client failed")
	}

	opt := DBCommon.NewDefaultOptions()
	for _, option := range options {
		option(&opt)
	}

	names := ""
	for _, name := range nodeNames {
		names += fmt.Sprintf("%s|", name)
	}
	queryLabelsString := ""
	if names != "" {
		names = strings.TrimSuffix(names, "|")
		queryLabelsString = fmt.Sprintf(`%s=~"%s"`, NodeCpuUsagePercentageLabelNode, names)
	}

	//metricName = EntityPromthNodeCpu.MetricName
	metricNameSum := NodeCpuUsagePercentageMetricNameSum
	metricNameAvg := NodeCpuUsagePercentageMetricNameAvg
	queryExpressionSum := fmt.Sprintf("%s{%s}", metricNameSum, queryLabelsString)
	queryExpressionAvg := fmt.Sprintf("%s{%s}", metricNameAvg, queryLabelsString)
	stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
	queryExpressionSum, err = Prometheus.WrapQueryExpression(queryExpressionSum, opt.AggregateOverTimeFunc, stepTimeInSeconds)
	if err != nil {
		return nil, errors.Wrap(err, "wrap query expression failed")
	}
	queryExpressionAvg, err = Prometheus.WrapQueryExpression(queryExpressionAvg, opt.AggregateOverTimeFunc, stepTimeInSeconds)
	if err != nil {
		return nil, errors.Wrap(err, "wrap query expression failed")
	}

	queryExpression := fmt.Sprintf("1000 * sum(%s * %s)", queryExpressionSum, queryExpressionAvg)
	scope.Debugf("Query to prometheus: queryExpression: %+v, StartTime: %+v, EndTime: %+v, StepTime: %+v", queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
	response, err := prometheusClient.QueryRange(ctx, queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
	if err != nil {
		return nil, errors.Wrap(err, "query prometheus failed")
	} else if response.Status != Prometheus.StatusSuccess {
		return nil, errors.Errorf("query prometheus failed: receive error response from prometheus: %s", response.Error)
	}

	entities, err := response.GetEntities()
	if err != nil {
		return nil, errors.Wrap(err, "get prometheus entities failed")
	}
	nodeCPUUsageMillicoresEntities := make([]EntityPromthMetric.NodeCPUUsageMillicoresEntity, len(entities))
	for i, entity := range entities {
		e := n.newNodeCPUUsageMillicoresEntity(entity)
		nodeCPUUsageMillicoresEntities[i] = e
	}

	return nodeCPUUsageMillicoresEntities, nil
}

func (n NodeCPUUsageRepository) newNodeCPUUsageMillicoresEntity(e Prometheus.Entity) EntityPromthMetric.NodeCPUUsageMillicoresEntity {

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

	return EntityPromthMetric.NodeCPUUsageMillicoresEntity{
		NodeName: e.Labels[NodeCpuUsagePercentageLabelNode],
		Samples:  samples,
	}
}
