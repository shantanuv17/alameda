package datamappingmgt

import (
	"github.com/containers-ai/alameda/datahub/pkg/config"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"sync"
)

var (
	scope = log.RegisterScope("data-mapping-mgt", "data mapping management", 0)
)

var (
	InfluxConfig       *influxdb.Config
	RWLock             *sync.RWMutex
	DataMappings       *Map
	DataMappings2Write *Map
)

func DataMappingInit(config *config.Config) error {
	InfluxConfig = config.InfluxDB
	RWLock = new(sync.RWMutex)
	DataMappings = NewMap()
	DataMappings2Write = NewMap()
	return nil
}
