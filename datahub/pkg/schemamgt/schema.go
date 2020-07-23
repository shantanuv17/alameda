package schemamgt

import (
	"github.com/containers-ai/alameda/datahub/pkg/config"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"sync"
)

var (
	scope = log.RegisterScope("schema", "schema library", 0)
)

var (
	InfluxConfig  *influxdb.Config
	RWLock        *sync.RWMutex
	Schemas       *SchemaMap
	Schemas2Write *SchemaMap
)

func SchemaInit(config *config.Config) error {
	InfluxConfig = config.InfluxDB
	RWLock = new(sync.RWMutex)
	Schemas = NewSchemaMap()
	Schemas2Write = NewSchemaMap()
	return nil
}
