package predictions

import (
	"prophetstor.com/alameda/pkg/database/influxdb"
)

const (
	DutyCycle                      influxdb.Measurement = "nvidia_gpu_duty_cycle"
	DutyCycleLowerBound            influxdb.Measurement = "nvidia_gpu_duty_cycle_lower_bound"
	DutyCycleUpperBound            influxdb.Measurement = "nvidia_gpu_duty_cycle_upper_bound"
	MemoryUsagePercentage          influxdb.Measurement = "nvidia_gpu_memory_usage_percentage"
	MemoryUsedBytes                influxdb.Measurement = "nvidia_gpu_memory_used_bytes"
	MemoryUsedBytesLowerBound      influxdb.Measurement = "nvidia_gpu_memory_used_bytes_lower_bound"
	MemoryUsedBytesUpperBound      influxdb.Measurement = "nvidia_gpu_memory_used_bytes_upper_bound"
	PowerUsageMilliWatts           influxdb.Measurement = "nvidia_gpu_power_usage_milliwatts"
	PowerUsageMilliWattsLowerBound influxdb.Measurement = "nvidia_gpu_power_usage_milliwatts_lower_bound"
	PowerUsageMilliWattsUpperBound influxdb.Measurement = "nvidia_gpu_power_usage_milliwatts_upper_bound"
	TemperatureCelsius             influxdb.Measurement = "nvidia_gpu_temperature_celsius"
	TemperatureCelsiusLowerBound   influxdb.Measurement = "nvidia_gpu_temperature_celsius_lower_bound"
	TemperatureCelsiusUpperBound   influxdb.Measurement = "nvidia_gpu_temperature_celsius_upper_bound"
)
