package licenses

const (
	defaultCPUCapacityGracePeriod = 1209600 // 14 days in seconds
)

type CPUCapacity struct {
	GracePeriod int64
}

type CapacityConfig struct {
	CPU *CPUCapacity
}

type Config struct {
	Capacity *CapacityConfig
}

func NewDefaultConfig() *Config {
	var config = Config{
		Capacity: NewCapacityConfig(),
	}
	return &config
}

func NewCapacityConfig() *CapacityConfig {
	var config = CapacityConfig{
		CPU: NewCPUCapacity(),
	}
	return &config
}

func NewCPUCapacity() *CPUCapacity {
	var config = CPUCapacity{
		GracePeriod: defaultCPUCapacityGracePeriod,
	}
	return &config
}
