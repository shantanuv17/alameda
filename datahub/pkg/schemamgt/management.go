package schemamgt

import (
	"errors"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/models"
	InternalSchemas "github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"strconv"
)

type SchemaManagement struct {
}

func NewSchemaManagement() *SchemaManagement {
	schemaManagement := SchemaManagement{}
	return &schemaManagement
}

func (p *SchemaManagement) Refresh() error {
	Schemas.Empty()
	Schemas.Schemas = make(map[InternalSchemas.Scope][]*InternalSchemas.Schema)

	// Build measurement definition
	for table, measurement := range MeasurementNameMap {
		Schemas.Schemas[table] = make([]*InternalSchemas.Schema, 0)
		results, err := p.read(measurement)
		if len(results) == 0 {
			continue
		}
		if err != nil {
			scope.Error(err.Error())
			scope.Error("failed to read measurement definition")
			return errors.New("failed to refresh schemas")
		}
		if results[0].GetGroupNum() == 0 {
			continue
		}
		for _, row := range results[0].GetGroup(0).GetRows() {
			schema := InternalSchemas.NewSchema(table, row["category"], row["type"])
			Schemas.Schemas[table] = append(Schemas.Schemas[table], schema)
		}
	}

	// Build measurement schema definition
	for table, measurement := range MeasurementSchemaNameMap {
		results, err := p.read(measurement)
		if len(results) == 0 {
			continue
		}
		if err != nil {
			scope.Error(err.Error())
			scope.Error("failed to read measurement schema definition")
			return errors.New("failed to refresh schemas")
		}
		if results[0].GetGroupNum() == 0 {
			continue
		}
		for _, row := range results[0].GetGroup(0).GetRows() {
			for _, schema := range Schemas.Schemas[table] {
				if schema.SchemaMeta.Category == row["category"] && schema.SchemaMeta.Type == row["type"] {
					metricType, _ := strconv.ParseInt(row["metric_type"], 10, 64)
					boundary, _ := strconv.ParseInt(row["resource_boundary"], 10, 64)
					quota, _ := strconv.ParseInt(row["resource_quota"], 10, 64)
					schema.AddMeasurement(row["measurement"], InternalSchemas.MetricType(metricType), InternalSchemas.ResourceBoundary(boundary), InternalSchemas.ResourceQuota(quota), row["columns"])
				}
			}
		}
	}

	return nil
}

func (p *SchemaManagement) AddSchemas(schemas []*InternalSchemas.Schema) {
	if schemas != nil {
		for _, schema := range schemas {
			Schemas2Write.AddSchema(schema)
		}
	}
}

func (p *SchemaManagement) GetSchemas(scope InternalSchemas.Scope, category, schemaType string) []*InternalSchemas.Schema {
	// TODO: check if need to get READ LOCK
	// Filter table
	tables := make([]*InternalSchemas.Schema, 0)
	if scope == InternalSchemas.ScopeUndefined {
		for _, s := range Schemas.Schemas {
			for _, schema := range s {
				tables = append(tables, schema)
			}
		}
	} else {
		for _, schema := range Schemas.Schemas[scope] {
			tables = append(tables, schema)
		}
	}

	// Filter category
	categories := make([]*InternalSchemas.Schema, 0)
	if category != "" {
		for _, schema := range tables {
			if schema.SchemaMeta.Category == category {
				categories = append(categories, schema)
			}
		}
	} else {
		for _, schema := range tables {
			categories = append(categories, schema)
		}
	}

	// Filter type
	schemas := make([]*InternalSchemas.Schema, 0)
	if schemaType != "" {
		for _, schema := range categories {
			if schema.SchemaMeta.Type == schemaType {
				schemas = append(schemas, schema)
			}
		}
	} else {
		for _, schema := range categories {
			schemas = append(schemas, schema)
		}
	}

	return schemas
}

func (p *SchemaManagement) DeleteSchemas(table InternalSchemas.Scope, category, schemaType string) error {
	influxClient := &influxdb.InfluxClient{
		Address:  InfluxConfig.Address,
		Username: InfluxConfig.Username,
		Password: InfluxConfig.Password,
	}

	// Delete measurement definition
	cmd := p.buildDropMeasurementsQuery(table, category, schemaType)
	_, err := influxClient.QueryDB(cmd, "alameda_schema")
	if err != nil {
		scope.Error(err.Error())
		return errors.New("failed to delete measurement definition")
	}

	// Delete measurement schema definition
	cmd = p.buildDropSchemasQuery(table, category, schemaType)
	_, err = influxClient.QueryDB(cmd, "alameda_schema")
	if err != nil {
		scope.Error(err.Error())
		return errors.New("failed to delete measurement schema definition")
	}

	err = p.Refresh()
	if err != nil {
		scope.Error(err.Error())
		return errors.New("failed to refresh schemas")
	}

	return nil
}

func (p *SchemaManagement) Flush() error {
	points := make([]*InfluxClient.Point, 0)

	// Build influxdb data points
	for table, schemas := range Schemas2Write.Schemas {
		pts, err := p.buildMeasurementPoints(table, schemas)
		if err != nil {
			return err
		}
		for _, pt := range pts {
			points = append(points, pt)
		}

		pts, err = p.buildSchemaPoints(table, schemas)
		if err != nil {
			return err
		}
		for _, pt := range pts {
			points = append(points, pt)
		}
	}

	// Batch write influxdb data points
	influxClient := &influxdb.InfluxClient{
		Address:  InfluxConfig.Address,
		Username: InfluxConfig.Username,
		Password: InfluxConfig.Password,
	}
	err := influxClient.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: "alameda_schema",
	})
	if err != nil {
		scope.Error(err.Error())
		return errors.New("failed to batch write influxdb data points")
	}

	// Append Schemas2Write to Schemas and empty it
	for _, schemas := range Schemas2Write.Schemas {
		for _, schema := range schemas {
			Schemas.AddSchema(schema)
		}
	}
	Schemas2Write.Empty()

	return nil
}

func (p *SchemaManagement) buildMeasurementPoints(table InternalSchemas.Scope, schemas []*InternalSchemas.Schema) ([]*InfluxClient.Point, error) {
	points := make([]*InfluxClient.Point, 0)

	for _, schema := range schemas {
		// Pack influx tags
		tags := map[string]string{
			"category": schema.SchemaMeta.Category,
			"type":     schema.SchemaMeta.Type,
		}

		// Pack influx fields
		fields := map[string]interface{}{
			"measurements": schema.String(),
		}

		// Add to influx point list
		pt, err := InfluxClient.NewPoint(MeasurementNameMap[table], tags, fields, influxdb.ZeroTime)
		if err != nil {
			scope.Error(err.Error())
			return make([]*InfluxClient.Point, 0), errors.New("failed to instance measurement influxdb data point")
		}
		points = append(points, pt)
	}

	return points, nil
}

func (p *SchemaManagement) buildSchemaPoints(table InternalSchemas.Scope, schemas []*InternalSchemas.Schema) ([]*InfluxClient.Point, error) {
	points := make([]*InfluxClient.Point, 0)

	for _, schema := range schemas {
		for _, measurement := range schema.Measurements {
			// Pack influx tags
			tags := map[string]string{
				"category":          schema.SchemaMeta.Category,
				"type":              schema.SchemaMeta.Type,
				"measurement":       measurement.Name,
				"metric_type":       strconv.FormatInt(int64(measurement.MetricType), 10),
				"resource_boundary": strconv.FormatInt(int64(measurement.Boundary), 10),
				"resource_quota":    strconv.FormatInt(int64(measurement.Quota), 10),
			}

			// Pack influx fields
			fields := map[string]interface{}{
				"columns": measurement.String(),
			}

			// Add to influx point list
			pt, err := InfluxClient.NewPoint(MeasurementSchemaNameMap[table], tags, fields, influxdb.ZeroTime)
			if err != nil {
				scope.Error(err.Error())
				return make([]*InfluxClient.Point, 0), errors.New("failed to instance schema influxdb data point")
			}
			points = append(points, pt)
		}
	}

	return points, nil
}

func (p *SchemaManagement) buildDropMeasurementsQuery(table InternalSchemas.Scope, category, schemaType string) string {
	keys := make([]string, 0)
	values := make([]string, 0)
	operators := make([]string, 0)
	types := make([]common.DataType, 0)

	if category != "" {
		keys = append(keys, "category")
		values = append(values, category)
		operators = append(operators, "=")
		types = append(types, common.String)
	}

	if schemaType != "" {
		keys = append(keys, "type")
		values = append(values, schemaType)
		operators = append(operators, "=")
		types = append(types, common.String)
	}

	query := influxdb.NewQuery(nil, MeasurementNameMap[table])
	query.AppendCondition(keys, values, operators, types)

	return query.BuildDropCmd()
}

func (p *SchemaManagement) buildDropSchemasQuery(table InternalSchemas.Scope, category, schemaType string) string {
	keys := make([]string, 0)
	values := make([]string, 0)
	operators := make([]string, 0)
	types := make([]common.DataType, 0)

	if category != "" {
		keys = append(keys, "category")
		values = append(values, category)
		operators = append(operators, "=")
		types = append(types, common.String)
	}

	if schemaType != "" {
		keys = append(keys, "type")
		values = append(values, schemaType)
		operators = append(operators, "=")
		types = append(types, common.String)
	}

	query := influxdb.NewQuery(nil, MeasurementSchemaNameMap[table])
	query.AppendCondition(keys, values, operators, types)

	return query.BuildDropCmd()
}

func (p *SchemaManagement) read(measurement string) ([]*models.InfluxResultExtend, error) {
	influxClient := &influxdb.InfluxClient{
		Address:  InfluxConfig.Address,
		Username: InfluxConfig.Username,
		Password: InfluxConfig.Password,
	}

	if !influxClient.MeasurementExist("alameda_schema", measurement) {
		scope.Infof("measurement(%s: %s) is not found", "alameda_schema", measurement)
		return make([]*models.InfluxResultExtend, 0), nil
	}

	statement := influxdb.Statement{
		Measurement: influxdb.Measurement(measurement),
	}
	cmd := statement.BuildQueryCmd()

	response, err := influxClient.QueryDB(cmd, "alameda_schema")
	if err != nil {
		scope.Error(err.Error())
		return make([]*models.InfluxResultExtend, 0), errors.New("failed to read database")
	}

	return models.NewInfluxResults(response), nil
}
