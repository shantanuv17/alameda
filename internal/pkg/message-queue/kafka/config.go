package kafka

import (
	"time"
)

type Config struct {
	Addresses   []string       `mapstructure:"brokerAddresses"`
	Version     string         `mapstructure:"version"`
	DialTimeout *time.Duration `mapstructure:""`
	KeepAlive   *time.Duration `mapstructure:""`

	SASL *SASLConfig `mapstructure:"sasl"`
	TLS  *TLSConfig  `mapstructure:"tls"`
}

type TLSConfig struct {
	Enabled            bool `mapstructure:"enabled"`
	InsecureSkipVerify bool `mapstructure:"insecureSkipVerify"`
}

type SASLConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
