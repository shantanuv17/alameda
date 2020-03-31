package datamappingmgt

import (
	"errors"
	"github.com/containers-ai/alameda/datahub/pkg/datamappingmgt/datamapping"
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/models"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"strconv"
)

type DataMappingManagement struct {
}

func NewDataMappingManagement() *DataMappingManagement {
	management := DataMappingManagement{}
	return &management
}

func (p *DataMappingManagement) Refresh() error {
	DataMappings.Empty()
	DataMappings.Initialize()

	for table, measurement := range MeasurementNameMap {
		results, err := p.read(measurement)
		if len(results) == 0 {
			continue
		}
		if err != nil {
			scope.Error(err.Error())
			scope.Error("failed to read data mapping definition")
			return errors.New("failed to refresh data mapping")
		}
		if results[0].GetGroupNum() == 0 {
			continue
		}
		for _, row := range results[0].GetGroup(0).GetRows() {
			category := row["category"]
			schemaType := row["type"]
			name := row["column"]
			mapping := row["mapping"]
			metricType, _ := strconv.ParseInt(row["metric_type"], 10, 64)
			source, _ := strconv.ParseInt(row["data_source"], 10, 64)

			dataMappings := p.GetDataMappings(table, category, schemaType, schemas.MetricType(metricType))
			if len(dataMappings) == 0 {
				dataMapping := datamapping.NewDataMapping(table, category, schemaType, schemas.MetricType(metricType))
				dataMapping.AddColumn(name, datamapping.Source(source), mapping)
				DataMappings.AddDataMapping(dataMapping)
			} else {
				dataMapping := dataMappings[0]
				dataMapping.AddColumn(name, datamapping.Source(source), mapping)
			}
		}
	}

	return nil
}

func (p *DataMappingManagement) AddDataMappings(sources []*datamapping.DataMapping) {
	if sources != nil {
		for _, source := range sources {
			DataMappings2Write.AddDataMapping(source)
		}
	}
}

func (p *DataMappingManagement) GetDataMappings(scope schemas.Scope, category, schemaType string, metricType schemas.MetricType) []*datamapping.DataMapping {
	// Filter table
	tables := make([]*datamapping.DataMapping, 0)
	if scope == schemas.ScopeUndefined {
		for _, dataMappings := range DataMappings.DataMappings {
			for _, dataMapping := range dataMappings {
				tables = append(tables, dataMapping)
			}
		}
	} else {
		for _, dataMapping := range DataMappings.DataMappings[scope] {
			tables = append(tables, dataMapping)
		}
	}

	// Filter category
	categories := make([]*datamapping.DataMapping, 0)
	if category != "" {
		for _, dataMapping := range tables {
			if dataMapping.SchemaMeta.Category == category {
				categories = append(categories, dataMapping)
			}
		}
	} else {
		for _, dataMapping := range tables {
			categories = append(categories, dataMapping)
		}
	}

	// Filter type
	schemaTypes := make([]*datamapping.DataMapping, 0)
	if schemaType != "" {
		for _, dataMapping := range categories {
			if dataMapping.SchemaMeta.Type == schemaType {
				schemaTypes = append(schemaTypes, dataMapping)
			}
		}
	} else {
		for _, dataMapping := range categories {
			schemaTypes = append(schemaTypes, dataMapping)
		}
	}

	// Filter metric type
	mappings := make([]*datamapping.DataMapping, 0)
	if metricType != schemas.MetricTypeUndefined {
		for _, dataMapping := range schemaTypes {
			if dataMapping.MetricType == metricType {
				mappings = append(mappings, dataMapping)
			}
		}
	} else {
		for _, dataMapping := range schemaTypes {
			mappings = append(mappings, dataMapping)
		}
	}

	return mappings
}

func (p *DataMappingManagement) DeleteDataMappings(table schemas.Scope, category, schemaType string, metricType schemas.MetricType) error {
	influxClient := &influxdb.InfluxClient{
		Address:  InfluxConfig.Address,
		Username: InfluxConfig.Username,
		Password: InfluxConfig.Password,
	}

	cmd := p.drop(table, category, schemaType, metricType)
	_, err := influxClient.QueryDB(cmd, DatabaseName)
	if err != nil {
		scope.Error(err.Error())
		return errors.New("failed to delete data mappings")
	}

	err = p.Refresh()
	if err != nil {
		scope.Error(err.Error())
		return errors.New("failed to refresh data mappings")
	}

	return nil
}

func (p *DataMappingManagement) Flush() error {
	points := make([]*InfluxClient.Point, 0)

	// Build influxdb data points
	for scope, dataSources := range DataMappings2Write.DataMappings {
		pts, err := p.buildColumnPoints(scope, dataSources)
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
		Database: DatabaseName,
	})
	if err != nil {
		scope.Error(err.Error())
		return errors.New("failed to batch write influxdb data points")
	}

	// Append DataSources2Write to DataSources and empty it
	for _, dataSources := range DataMappings2Write.DataMappings {
		for _, dataSource := range dataSources {
			DataMappings.AddDataMapping(dataSource)
		}
	}
	DataMappings2Write.Empty()

	return nil
}

func (p *DataMappingManagement) buildColumnPoints(table schemas.Scope, dataMappings []*datamapping.DataMapping) ([]*InfluxClient.Point, error) {
	points := make([]*InfluxClient.Point, 0)

	for _, dataMapping := range dataMappings {
		for _, column := range dataMapping.Columns {
			for _, sourceMapping := range column.SourceMappings {
				// Pack influx tags
				tags := map[string]string{
					"category":    dataMapping.SchemaMeta.Category,
					"type":        dataMapping.SchemaMeta.Type,
					"column":      column.Name,
					"metric_type": strconv.FormatInt(int64(dataMapping.MetricType), 10),
					"data_source": strconv.FormatInt(int64(sourceMapping.Source), 10),
				}

				// Pack influx fields
				fields := map[string]interface{}{
					"mapping": sourceMapping.Mapping,
				}

				// Add to influx point list
				pt, err := InfluxClient.NewPoint(MeasurementNameMap[table], tags, fields, influxdb.ZeroTime)
				if err != nil {
					scope.Error(err.Error())
					return make([]*InfluxClient.Point, 0), errors.New("failed to instance schema influxdb data point")
				}
				points = append(points, pt)
			}
		}
	}

	return points, nil
}

func (p *DataMappingManagement) read(measurement string) ([]*models.InfluxResultExtend, error) {
	influxClient := &influxdb.InfluxClient{
		Address:  InfluxConfig.Address,
		Username: InfluxConfig.Username,
		Password: InfluxConfig.Password,
	}

	if !influxClient.MeasurementExist(DatabaseName, measurement) {
		scope.Infof("measurement(%s: %s) is not found", DatabaseName, measurement)
		return make([]*models.InfluxResultExtend, 0), nil
	}

	statement := influxdb.Statement{
		Measurement: influxdb.Measurement(measurement),
	}
	cmd := statement.BuildQueryCmd()

	response, err := influxClient.QueryDB(cmd, DatabaseName)
	if err != nil {
		scope.Error(err.Error())
		return make([]*models.InfluxResultExtend, 0), errors.New("failed to read database")
	}

	return models.NewInfluxResults(response), nil
}

func (p *DataMappingManagement) drop(table schemas.Scope, category, schemaType string, metricType schemas.MetricType) string {
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

	if metricType != schemas.MetricTypeUndefined {
		keys = append(keys, "metric_type")
		values = append(values, strconv.FormatInt(int64(metricType), 10))
		operators = append(operators, "=")
		types = append(types, common.String)
	}

	query := influxdb.NewQuery(nil, MeasurementNameMap[table])
	query.AppendCondition(keys, values, operators, types)

	return query.BuildDropCmd()
}
