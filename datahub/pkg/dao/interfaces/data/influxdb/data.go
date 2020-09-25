package influxdb

import (
	"errors"
	DaoDataTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/data/types"
	SchemaMgt "prophetstor.com/alameda/datahub/pkg/schemamgt"
	"prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/alameda/pkg/database/influxdb"
	InfluxSchema "prophetstor.com/alameda/pkg/database/influxdb/schemas"
	Log "prophetstor.com/alameda/pkg/utils/log"
)

var (
	scope = Log.RegisterScope("dao-data-impl", "dao data implementation", 0)
)

type Data struct {
	InfluxDBConfig influxdb.Config
}

func NewDataWithConfig(config influxdb.Config) DaoDataTypes.DataDAO {
	return &Data{InfluxDBConfig: config}
}

func (p *Data) WriteData(request *DaoDataTypes.WriteDataRequest) error {
	if request.SchemaMeta == nil || request.WriteData == nil {
		return nil
	}

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)[0]
	for _, w := range request.WriteData {
		m := schema.GetMeasurement(w.Measurement, w.MetricType, w.Boundary, w.Quota)
		if m == nil {
			scope.Error("measurement is not found when writing data")
			return errors.New("measurement is not found")
		}
		measurement := influxdb.NewMeasurement(InfluxSchema.DatabaseNameMap[request.SchemaMeta.Scope], m, p.InfluxDBConfig)
		err := measurement.Write(w.Columns, w.Rows)
		if err != nil {
			scope.Error(err.Error())
			return err
		}
	}

	return nil
}

func (p *Data) ReadData(request *DaoDataTypes.ReadDataRequest) (*DaoDataTypes.Data, error) {
	if request.SchemaMeta == nil || request.ReadData == nil {
		return nil, nil
	}

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)[0]

	data := DaoDataTypes.Data{}
	data.SchemaMeta = InfluxSchema.NewSchemaMeta(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)
	for _, r := range request.ReadData {
		m := schema.GetMeasurement(r.Measurement, r.MetricType, r.Boundary, r.Quota)
		if m == nil {
			scope.Error("measurement is not found when reading data")
			return nil, errors.New("measurement is not found")
		}
		measurement := influxdb.NewMeasurement(InfluxSchema.DatabaseNameMap[request.SchemaMeta.Scope], m, p.InfluxDBConfig)
		groups, err := measurement.Read(influxdb.NewQuery(r.QueryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return nil, err
		}
		rawdata := DaoDataTypes.Rawdata{
			Measurement: measurement.Name,
			MetricType:  measurement.MetricType,
			Boundary:    measurement.Boundary,
			Quota:       measurement.Quota,
			Groups:      groups,
		}
		data.Rawdata = append(data.Rawdata, &rawdata)
	}

	return &data, nil
}

func (p *Data) DeleteData(request *DaoDataTypes.DeleteDataRequest) error {
	if request.SchemaMeta == nil || request.DeleteData == nil {
		return nil
	}

	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)[0]

	data := DaoDataTypes.Data{}
	data.SchemaMeta = InfluxSchema.NewSchemaMeta(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)
	for _, d := range request.DeleteData {
		m := schema.GetMeasurement(d.Measurement, d.MetricType, d.Boundary, d.Quota)
		if m == nil {
			scope.Error("measurement is not found when reading data")
			return errors.New("measurement is not found")
		}
		measurement := influxdb.NewMeasurement(InfluxSchema.DatabaseNameMap[request.SchemaMeta.Scope], m, p.InfluxDBConfig)
		err := measurement.Drop(influxdb.NewQuery(d.QueryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return err
		}
	}

	return nil
}

func (p *Data) WriteMeta(request *DaoDataTypes.WriteMetaRequest) error {
	schemaMgt := SchemaMgt.NewSchemaManagement()
	schema := schemaMgt.GetSchemas(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)[0]

	data := DaoDataTypes.Data{}
	data.SchemaMeta = InfluxSchema.NewSchemaMeta(request.SchemaMeta.Scope, request.SchemaMeta.Category, request.SchemaMeta.Type)
	for _, d := range request.WriteMeta {
		dataInDB := make([]*influxdb.InfluxData, 0)
		data2Update := make([]*influxdb.InfluxData, 0)
		data2Delete := make([]*influxdb.InfluxData, 0)

		m := schema.GetMeasurement(d.Measurement, d.MetricType, d.Boundary, d.Quota)
		if m == nil {
			scope.Error("measurement is not found when reading data")
			return errors.New("measurement is not found")
		}
		measurement := influxdb.NewMeasurement(InfluxSchema.DatabaseNameMap[request.SchemaMeta.Scope], m, p.InfluxDBConfig)

		// Generate the records which will be written or updated later
		for _, row := range d.Rows {
			data2Update = append(data2Update, influxdb.NewInfluxData(d.Columns, row, measurement.Columns))
		}

		// Read all records in influxdb by write meta
		queryCondition := common.QueryCondition{}
		queryCondition.WhereCondition = append(queryCondition.WhereCondition, d.Condition)
		groups, err := measurement.Read(influxdb.NewQuery(&queryCondition, measurement.Name))
		if err != nil {
			scope.Error(err.Error())
			return err
		}
		for _, group := range groups {
			for _, row := range group.Rows {
				dataInDB = append(dataInDB, influxdb.NewInfluxData(group.Columns, row, measurement.Columns))
			}
		}

		// Find out those records which are NOT in write meta
		for _, record := range dataInDB {
			found := false
			for _, r := range data2Update {
				if influxdb.CompareInfluxDataByTags(r, record) {
					found = true
					break
				}
			}
			if !found {
				data2Delete = append(data2Delete, record)
			}
		}

		// Do drop first (delete those records which are NOT in write meta)
		if len(data2Delete) > 0 {
			err = measurement.Drop(influxdb.NewQuery(influxdb.GenerateQueryConditionByInfluxData(data2Delete), measurement.Name))
			if err != nil {
				scope.Error(err.Error())
				return err
			}
		}

		// Then do write or update
		err = measurement.Write(d.Columns, d.Rows)
		if err != nil {
			scope.Error(err.Error())
			return err
		}
	}

	return nil
}
