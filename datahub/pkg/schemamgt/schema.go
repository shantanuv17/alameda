package schemamgt

import (
	"prophetstor.com/alameda/pkg/database/influxdb"
	"prophetstor.com/alameda/pkg/utils/log"
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

func SchemaInit(config *influxdb.Config) error {
	InfluxConfig = config
	RWLock = new(sync.RWMutex)
	Schemas = NewSchemaMap()
	Schemas2Write = NewSchemaMap()
	return nil
}
