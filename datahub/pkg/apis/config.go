package apis

// Configuration of APIs
type Config struct {
	Metrics *MetricsConfig `mapstructure:"metrics"`
}

// Configuration of metrics related APIs
type MetricsConfig struct {
	Source         string          `mapstructure:"source"`
	Target         string          `mapstructure:"target"`
	RollupFunction *RollupFunction `mapstructure:"rollupFunction"`
}

type RollupFunction struct {
	Function string `mapstructure:"function"`
	Number   int64  `mapstructure:"number"`
}

// Provide default configuration for APIs
func NewDefaultConfig() *Config {
	var config = Config{
		Metrics: &MetricsConfig{
			Source: "prometheus",
			Target: "influxdb",
			RollupFunction: &RollupFunction{
				Function: "percentile",
				Number:   95,
			},
		},
	}
	return &config
}

// Confirm the APIs configuration is validated
func (c *Config) Validate() error {
	return nil
}
