package metrics

import (
	"prophetstor.com/alameda/datahub/pkg/account-mgt/keycodes"
	"prophetstor.com/alameda/datahub/pkg/schemamgt"
	"prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/alameda/pkg/database/influxdb"
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
	"prophetstor.com/api/datahub/events"
	"strconv"
	"time"
)

const (
	DefaultMeteringFederatoraiEnabled = true
	DefaultMeteringFederatoraiSpecs   = "0 0 0 * * *"
)

type MeteringFederatorai struct {
	AlertMetrics

	eventLevel  map[int]events.EventLevel
	eventPosted map[int]bool
	influxCfg   *influxdb.Config
}

func NewMeteringFederatorai(notifier *Notifier, influxCfg *influxdb.Config) *MeteringFederatorai {
	alert := MeteringFederatorai{}
	alert.notifier = notifier
	alert.name = "capacity"
	alert.alertType = "federator"
	alert.category = "metering"
	alert.criteriaType = CriteriaTypeUndefined
	alert.eventLevel = make(map[int]events.EventLevel, 0)
	alert.eventPosted = make(map[int]bool, 0)
	alert.influxCfg = influxCfg
	return &alert
}

func (c *MeteringFederatorai) Validate() {
	scope.Info("log metering data")

	tsNow := time.Now()
	remainder := tsNow.Unix() % 86400
	ts := time.Unix(tsNow.Unix()-remainder, 0)

	// Get schema read lock
	schemamgt.RWLock.RLock()
	defer schemamgt.RWLock.RUnlock()

	numberCluster := c.numberCluster()
	numberNode := c.numberNode()
	numberNamespace := c.numberNamespace()
	numberApplication := c.numberApplication()
	numberController := c.numberController()

	occupied, _ := keycodes.GetFederatoraiCapacityOccupied(c.influxCfg)
	cpuCores := strconv.FormatInt(occupied.CPUCores, 10)
	memoryBytes := strconv.FormatInt(occupied.MemoryBytes, 10)

	// Write data
	measurement := c.measurement(schemas.Metering, "federatorai", "capacity", "federatorai_capacity")
	columns := []string{"name", "number_cluster", "number_node", "number_namespace", "number_application", "number_controller", "cpu_cores", "memory_bytes"}
	row := &common.Row{Time: &ts, Values: []string{"federatorai", numberCluster, numberNode, numberNamespace, numberApplication, numberController, cpuCores, memoryBytes}}
	err := measurement.Write(columns, []*common.Row{row})
	if err != nil {
		scope.Error(err.Error())
	}
}

func (c *MeteringFederatorai) numberCluster() string {
	measurement := c.measurement(schemas.Resource, "cluster_status", "cluster", "cluster")
	groups, err := measurement.Read(c.query("cluster", measurement))
	if err != nil {
		scope.Errorf("failed to read number of cluster: %s", err.Error())
		return "0"
	}
	return groups[0].Rows[0].Values[0]
}

func (c *MeteringFederatorai) numberNode() string {
	measurement := c.measurement(schemas.Resource, "cluster_status", "node", "node")
	groups, err := measurement.Read(c.query("node", measurement))
	if err != nil {
		scope.Errorf("failed to read number of node: %s", err.Error())
		return "0"
	}
	return groups[0].Rows[0].Values[0]
}

func (c *MeteringFederatorai) numberNamespace() string {
	measurement := c.measurement(schemas.Resource, "cluster_status", "namespace", "namespace")
	groups, err := measurement.Read(c.query("namespace", measurement))
	if err != nil {
		scope.Errorf("failed to read number of namespace: %s", err.Error())
		return "0"
	}
	return groups[0].Rows[0].Values[0]
}

func (c *MeteringFederatorai) numberApplication() string {
	measurement := c.measurement(schemas.Resource, "cluster_status", "application", "application")
	groups, err := measurement.Read(c.query("application", measurement))
	if err != nil {
		scope.Errorf("failed to read number of application: %s", err.Error())
		return "0"
	}
	return groups[0].Rows[0].Values[0]
}

func (c *MeteringFederatorai) numberController() string {
	measurement := c.measurement(schemas.Resource, "cluster_status", "controller", "controller")
	groups, err := measurement.Read(c.query("controller", measurement))
	if err != nil {
		scope.Errorf("failed to read number of controller: %s", err.Error())
		return "0"
	}
	return groups[0].Rows[0].Values[0]
}

func (c *MeteringFederatorai) measurement(scope schemas.Scope, category, schemaType, name string) *influxdb.InfluxMeasurement {
	schema := SchemaMgt.GetSchemas(scope, category, schemaType)[0]
	cluster := schema.GetMeasurement(name, schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	return influxdb.NewMeasurement(schemas.DatabaseNameMap[scope], cluster, *c.influxCfg)
}

func (c *MeteringFederatorai) query(name string, measurement *influxdb.InfluxMeasurement) *influxdb.InfluxQuery {
	var fieldName string

	// Find out the first field which is required
	for _, field := range measurement.GetFields() {
		if field.Required {
			fieldName = field.Name
		}
	}

	// Generate query condition
	queryCondition := common.QueryCondition{
		Function: &common.Function{
			Type:   common.FunctionCount,
			Target: "value",
		},
		Selects: []string{fieldName},
	}

	return influxdb.NewQuery(&queryCondition, name)
}
