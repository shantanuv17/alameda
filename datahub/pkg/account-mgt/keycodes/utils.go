package keycodes

import (
	"github.com/containers-ai/alameda/datahub/pkg/schemamgt"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	"strconv"
)

type CapacityOccupied struct {
	CPUCores    int64
	MemoryBytes int64
}

func GetFederatoraiCapacityOccupied(influxCfg *influxdb.Config) (*CapacityOccupied, error) {
	var nameIndex = 0
	var clusterNameIndex = 0
	var coresIndex = 0
	var memoryIndex = 0

	// Get cluster-status node measurement
	measurement := GetMeasurement(schemas.Resource, "cluster_status", "node", "node", influxCfg)

	// Read all nodes
	groups, err := measurement.Read(influxdb.NewQuery(nil, "node"))
	if err != nil {
		scope.Errorf("failed to read nodes info when get federatorai capacity: %s", err.Error())
		return &CapacityOccupied{}, err
	}

	// Find indices of specific columns
	for index, column := range groups[0].Columns {
		if column == "name" {
			nameIndex = index
			continue
		}
		if column == "cluster_name" {
			clusterNameIndex = index
			continue
		}
		if column == "node_cpu_cores" {
			coresIndex = index
			continue
		}
		if column == "node_memory_bytes" {
			memoryIndex = index
			continue
		}
	}

	// Calculate the capacity occupied
	occupied := CapacityOccupied{}
	scope.Infof("Number of node: %d", len(groups[0].Rows))
	for _, row := range groups[0].Rows {
		cores, _ := strconv.ParseInt(row.Values[coresIndex], 10, 64)
		memory, _ := strconv.ParseInt(row.Values[memoryIndex], 10, 64)
		scope.Infof("  node: %s (cluster: %s), CPU Cores: %d, Memory Bytes: %d", row.Values[nameIndex], row.Values[clusterNameIndex], cores, memory)
		occupied.CPUCores += cores
		occupied.MemoryBytes += memory
	}

	return &occupied, nil
}

func GetMeasurement(scope schemas.Scope, category, schemaType, name string, influxCfg *influxdb.Config) *influxdb.InfluxMeasurement {
	schemaMgt := schemamgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(scope, category, schemaType)[0]
	cluster := schema.GetMeasurement(name, schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	return influxdb.NewMeasurement(schemas.DatabaseNameMap[scope], cluster, *influxCfg)
}
