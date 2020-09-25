package metrics

import (
	"prophetstor.com/alameda/pkg/database/influxdb"
)

const (
	DutyCycle            influxdb.Measurement = "nvidia_gpu_duty_cycle"
	MemoryTotalBytes     influxdb.Measurement = "nvidia_gpu_memory_total_bytes"
	MemoryUsedBytes      influxdb.Measurement = "nvidia_gpu_memory_used_bytes"
	NumDevices           influxdb.Measurement = "nvidia_gpu_num_devices"
	PowerUsageMilliWatts influxdb.Measurement = "nvidia_gpu_power_usage_milliwatts"
	TemperatureCelsius   influxdb.Measurement = "nvidia_gpu_temperature_celsius"
)
