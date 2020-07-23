package keycodes

import (
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/database/ldap"
)

const (
	defaultCliPath         = "/opt/prophetstor/federatorai/bin/license_main"
	defaultRefreshInterval = 180
)

// Configuration of keycode CLI
type Config struct {
	CliPath         string
	RefreshInterval int64
	AesKey          []byte
	InfluxDB        *influxdb.Config
	Ldap            *ldap.Config
}

// Provide default configuration for keycode CLI
func NewDefaultConfig() *Config {
	var config = Config{
		CliPath:         defaultCliPath,
		RefreshInterval: defaultRefreshInterval,
		InfluxDB:        influxdb.NewDefaultConfig(),
		Ldap:            ldap.NewDefaultConfig(),
	}
	return &config
}

// Confirm the keycode CLI configuration is validated
func (c *Config) Validate() error {
	return nil
}
