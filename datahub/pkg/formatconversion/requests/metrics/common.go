package metrics

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"prophetstor.com/alameda/datahub/pkg/apis"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	FormatCommon "prophetstor.com/alameda/datahub/pkg/formatconversion/requests/common"
	Log "prophetstor.com/alameda/pkg/utils/log"
	ApiCommon "prophetstor.com/api/datahub/common"
)

type MetricsDBType = string

const (
	MetricsDBTypePromethues MetricsDBType = "prometheus"
	MetricsDBTypeInfluxdb   MetricsDBType = "influxdb"
)

var (
	scope = Log.RegisterScope("request-extend-applications", "datahub(request-extend-applications) log", 0)
)

var MetricTypeNameMap = map[ApiCommon.MetricType]FormatEnum.MetricType{
	ApiCommon.MetricType_CPU_MILLICORES_USAGE: FormatEnum.MetricTypeCPUUsageSecondsPercentage,
	ApiCommon.MetricType_MEMORY_BYTES_USAGE:   FormatEnum.MetricTypeMemoryUsageBytes,
	ApiCommon.MetricType_POWER_USAGE_WATTS:    FormatEnum.MetricTypePowerUsageWatts,
	ApiCommon.MetricType_TEMPERATURE_CELSIUS:  FormatEnum.MetricTypeTemperatureCelsius,
	ApiCommon.MetricType_DUTY_CYCLE:           FormatEnum.MetricTypeDutyCycle,
}

func normalizeListMetricsRequestQueryConditionWthMetricsDBType(q *ApiCommon.QueryCondition, dbType MetricsDBType) *ApiCommon.QueryCondition {

	t := q.TimeRange
	if t == nil {
		t = &ApiCommon.TimeRange{}
	}
	normalizeT := normalizeListMetricsRequestTimeRangeByMetricsDBType(t, dbType)
	q.TimeRange = normalizeT

	return q
}

func normalizeListMetricsRequestTimeRange(t *ApiCommon.TimeRange) *ApiCommon.TimeRange {

	defaultStartTime := timestamp.Timestamp{}
	defaultEndTime := *ptypes.TimestampNow()
	defaultStep := duration.Duration{
		Seconds: 30,
	}

	if t.StartTime == nil {
		t.StartTime = &defaultStartTime
	}
	if t.EndTime == nil {
		t.EndTime = &defaultEndTime
	}
	if t.Step == nil {
		t.Step = &defaultStep
	}

	return t
}

func normalizeListMetricsRequestTimeRangeByMetricsDBType(t *ApiCommon.TimeRange, metricsDBType MetricsDBType) *ApiCommon.TimeRange {

	t = normalizeListMetricsRequestTimeRange(t)

	t.StartTime = &timestamp.Timestamp{
		Seconds: t.StartTime.Seconds - t.StartTime.Seconds%t.Step.Seconds,
	}

	switch metricsDBType {
	case MetricsDBTypePromethues:
		t.EndTime = &timestamp.Timestamp{
			Seconds: t.EndTime.Seconds - t.EndTime.Seconds%t.Step.Seconds,
		}
	case MetricsDBTypeInfluxdb:
		// If start time and end time are at the same HOUR frame
		if (t.StartTime.Seconds / t.Step.Seconds) == (t.EndTime.Seconds / t.Step.Seconds) {
			t.EndTime = &timestamp.Timestamp{
				Seconds: t.EndTime.Seconds + (t.Step.Seconds - (t.EndTime.Seconds % t.Step.Seconds) - 1),
			}
		} else {
			t.EndTime = &timestamp.Timestamp{
				Seconds: t.EndTime.Seconds - t.EndTime.Seconds%t.Step.Seconds - 1,
			}
		}
	}
	return t
}

func newFunction(config *apis.MetricsConfig) *ApiCommon.Function {
	f := ApiCommon.Function{}
	f.Type = FormatCommon.FunctionValueMap[config.RollupFunction.Function]
	f.Number = config.RollupFunction.Number
	return &f
}
