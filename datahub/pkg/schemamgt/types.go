package schemamgt

import (
	"prophetstor.com/alameda/datahub/pkg/entities"
	"prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/alameda/pkg/database/influxdb/schemas"
	"prophetstor.com/alameda/pkg/utils"
	"reflect"
	"strconv"
)

func NewSchema(entity interface{}) *schemas.Schema {
	schemaMeta := NewSchemaMeta(entity)
	schema := schemas.NewSchema(schemaMeta.Scope, schemaMeta.Category, schemaMeta.Type)
	schema.Measurements = append(schema.Measurements, NewMeasurement(entity))
	return schema
}

func NewSchemaMeta(entity interface{}) *schemas.SchemaMeta {
	datahubEntity := utils.ExtractField(entity, "DatahubEntity")
	schemaMeta := schemas.SchemaMeta{
		Scope:    schemas.ScopeValue[datahubEntity.Tag.Get("scope")],
		Category: datahubEntity.Tag.Get("category"),
		Type:     datahubEntity.Tag.Get("type"),
	}
	return &schemaMeta
}

func NewMeasurement(entity interface{}) *schemas.Measurement {
	m := utils.ExtractField(entity, "Measurement")
	measurement := schemas.Measurement{
		Name:       m.Tag.Get("name"),
		MetricType: schemas.MetricTypeValue[m.Tag.Get("metric")],
		Boundary:   schemas.ResourceBoundaryValue[m.Tag.Get("boundary")],
		Quota:      schemas.ResourceQuotaValue[m.Tag.Get("quota")],
		IsTS:       string2Bool(m.Tag.Get("ts")),
	}

	// Add columns
	entityType := reflect.TypeOf(entity).Elem()
	for i := entities.ColumnStartIndex + 1; i < entityType.NumField(); i++ {
		fieldType := entityType.Field(i)
		measurement.AddColumn(
			fieldType.Tag.Get("json"),
			string2Bool(fieldType.Tag.Get("required")),
			schemas.ColumnTypeValue[fieldType.Tag.Get("column")],
			common.DataTypeValue[fieldType.Type.Kind()],
		)
	}

	return &measurement
}

func string2Bool(str string) bool {
	valueBool, _ := strconv.ParseBool(str)
	return valueBool
}
