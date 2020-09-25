package config

import (
	"errors"
	Keycodes "prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	Licenses "prophetstor.com/alameda/datahub/pkg/account-mgt/licenses"
	Apis "prophetstor.com/alameda/datahub/pkg/apis"
	Notifier "prophetstor.com/alameda/datahub/pkg/notifier"
	InternalRabbitMQ "prophetstor.com/alameda/internal/pkg/message-queue/rabbitmq"
	InternalWeaveScope "prophetstor.com/alameda/internal/pkg/weavescope"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	LDAP "prophetstor.com/alameda/pkg/database/ldap"
	Prometheus "prophetstor.com/alameda/pkg/database/prometheus"
	"prophetstor.com/alameda/pkg/utils/log"
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
