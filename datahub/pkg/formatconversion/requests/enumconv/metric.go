package enumconv

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

var MetricTypeNameMap = map[ApiCommon.MetricType]schemas.MetricType{
	ApiCommon.MetricType_METRICS_TYPE_UNDEFINED:       schemas.MetricTypeUndefined,
	ApiCommon.MetricType_CPU_USAGE_SECONDS_PERCENTAGE: schemas.CPUUsageSecondsPercentage,
	ApiCommon.MetricType_MEMORY_USAGE_BYTES:           schemas.MemoryUsageBytes,
	ApiCommon.MetricType_POWER_USAGE_WATTS:            schemas.PowerUsageWatts,
	ApiCommon.MetricType_TEMPERATURE_CELSIUS:          schemas.TemperatureCelsius,
	ApiCommon.MetricType_DUTY_CYCLE:                   schemas.DutyCycle,
	ApiCommon.MetricType_CURRENT_OFFSET:               schemas.CurrentOffset,
	ApiCommon.MetricType_LAG:                          schemas.Lag,
	ApiCommon.MetricType_LATENCY:                      schemas.Latency,
	ApiCommon.MetricType_NUMBER:                       schemas.Number,
}
