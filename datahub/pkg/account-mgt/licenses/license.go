package licenses

// Configuration
var (
	CPUCapacityGracePeriod int64 = defaultCPUCapacityGracePeriod
)

func LicenseInit(config *Config) error {
	CPUCapacityGracePeriod = config.Capacity.CPU.GracePeriod
	return nil
}
