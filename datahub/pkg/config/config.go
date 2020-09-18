package config

import (
	"errors"
	Keycodes "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	Licenses "github.com/containers-ai/alameda/datahub/pkg/account-mgt/licenses"
	Apis "github.com/containers-ai/alameda/datahub/pkg/apis"
	Notifier "github.com/containers-ai/alameda/datahub/pkg/notifier"
	InternalRabbitMQ "github.com/containers-ai/alameda/internal/pkg/message-queue/rabbitmq"
	InternalWeaveScope "github.com/containers-ai/alameda/internal/pkg/weavescope"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	LDAP "github.com/containers-ai/alameda/pkg/database/ldap"
	Prometheus "github.com/containers-ai/alameda/pkg/database/prometheus"
	"github.com/containers-ai/alameda/pkg/utils/log"
)

const (
	defaultBindAddress = ":50050"
)

type Config struct {
	BindAddress string `mapstructure:"bindAddress"`
	ClusterUID  string
	Prometheus  *Prometheus.Config         `mapstructure:"prometheus"`
	InfluxDB    *InfluxDB.Config           `mapstructure:"influxdb"`
	Apis        *Apis.Config               `mapstructure:"apis"`
	Ldap        *LDAP.Config               `mapstructure:"ldap"`
	Keycode     *Keycodes.Config           `mapstructure:"keycode"`
	License     *Licenses.Config           `mapstructure:"license"`
	Notifier    *Notifier.Config           `mapstructure:"notifier"`
	WeaveScope  *InternalWeaveScope.Config `mapstructure:"weavescope"`
	RabbitMQ    *InternalRabbitMQ.Config   `mapstructure:"rabbitmq"`
	Log         *log.Config                `mapstructure:"log"`
}

func NewDefaultConfig() Config {
	var (
		defaultLogConfig        = log.NewDefaultConfig()
		defaultPrometheusConfig = Prometheus.NewDefaultConfig()
		defaultInfluxDBConfig   = InfluxDB.NewDefaultConfig()
		defaultApisConfig       = Apis.NewDefaultConfig()
		defaultLdapConfig       = LDAP.NewDefaultConfig()
		defaultKeycodeConfig    = Keycodes.NewDefaultConfig()
		defaultLicenseConfig    = Licenses.NewDefaultConfig()
		defaultNotifierConfig   = Notifier.NewDefaultConfig()
		defaultWeaveScopeConfig = InternalWeaveScope.NewDefaultConfig()
		defaultRabbitMQConfig   = InternalRabbitMQ.NewDefaultConfig()
		config                  = Config{
			BindAddress: defaultBindAddress,
			Prometheus:  defaultPrometheusConfig,
			InfluxDB:    defaultInfluxDBConfig,
			Apis:        defaultApisConfig,
			Ldap:        defaultLdapConfig,
			Keycode:     defaultKeycodeConfig,
			License:     defaultLicenseConfig,
			Notifier:    defaultNotifierConfig,
			WeaveScope:  defaultWeaveScopeConfig,
			RabbitMQ:    defaultRabbitMQConfig,
			Log:         &defaultLogConfig,
		}
	)

	defaultKeycodeConfig.InfluxDB = defaultInfluxDBConfig
	defaultKeycodeConfig.Ldap = nil // TODO: defaultLdapConfig

	return config
}

func (c *Config) Validate() error {
	var err error

	err = c.Prometheus.Validate()
	if err != nil {
		return errors.New("failed to validate gRPC config: " + err.Error())
	}

	return nil
}
