package datahub

import (
	"errors"
	"net/url"

	datahubutils "github.com/containers-ai/alameda/operator/pkg/utils/datahub"
)

type retry struct {
	Interval int64 `mapstructure:"interval"`
	Timeout  int64 `mapstructure:"timeout"`
}

type Config struct {
	Address string `mapstructure:"address"`
	Retry   retry  `mapstructure:"retry"`
}

func NewConfig() *Config {

	c := Config{}
	c.init()
	return &c
}

func (c *Config) init() {
	c.Address = datahubutils.GetDatahubAddress()
	c.Retry = retry{
		Interval: 3,
		Timeout:  30,
	}
}

func (c *Config) Validate() error {

	var err error

	_, err = url.Parse(c.Address)
	if err != nil {
		return errors.New("datahub config validate failed: " + err.Error())
	}

	return nil
}
