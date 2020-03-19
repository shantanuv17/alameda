package data

import (
	"errors"
	"github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/data/types"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/requests/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
)

func NewWriteData(writeData *data.WriteData) *types.WriteData {
	if writeData == nil {
		return nil
	}
	wData := types.WriteData{}
	wData.Measurement = writeData.GetMeasurement()
	wData.MetricType = enumconv.MetricTypeNameMap[writeData.GetMetricType()]
	wData.Boundary = enumconv.ResourceBoundaryNameMap[writeData.GetResourceBoundary()]
	wData.Quota = enumconv.ResourceQuotaNameMap[writeData.GetResourceQuota()]
	wData.Columns = writeData.GetColumns()
	wData.Rows = common.NewRows(writeData.GetRows())
	return &wData
}

func NewReadData(readData *data.ReadData) *types.ReadData {
	if readData == nil {
		return nil
	}
	rData := types.ReadData{}
	rData.Measurement = readData.GetMeasurement()
	rData.MetricType = enumconv.MetricTypeNameMap[readData.GetMetricType()]
	rData.Boundary = enumconv.ResourceBoundaryNameMap[readData.GetResourceBoundary()]
	rData.Quota = enumconv.ResourceQuotaNameMap[readData.GetResourceQuota()]
	rData.QueryCondition = common.NewQueryCondition(readData.GetQueryCondition())
	return &rData
}

func NewDeleteData(deleteData *data.DeleteData) *types.DeleteData {
	if deleteData == nil {
		return nil
	}
	dData := types.DeleteData{}
	dData.Measurement = deleteData.GetMeasurement()
	dData.MetricType = enumconv.MetricTypeNameMap[deleteData.GetMetricType()]
	dData.Boundary = enumconv.ResourceBoundaryNameMap[deleteData.GetResourceBoundary()]
	dData.Quota = enumconv.ResourceQuotaNameMap[deleteData.GetResourceQuota()]
	dData.QueryCondition = common.NewQueryCondition(deleteData.GetQueryCondition())
	return &dData
}

func isSchemaMetaComplete(schemaMeta *schemas.SchemaMeta) error {
	if int(schemaMeta.Scope) == 0 {
		return errors.New("schema meta(scope) is not given")
	}
	if schemaMeta.Category == "" {
		return errors.New("schema meta(category) is not given")
	}
	if schemaMeta.Type == "" {
		return errors.New("schema meta(type) is not given")
	}
	return nil
}

func isMeasurementFound(schema *schemas.Schema, measurement string, metricType schemas.MetricType, boundary schemas.ResourceBoundary, quota schemas.ResourceQuota) error {
	m := schema.GetMeasurement(measurement, metricType, boundary, quota)
	if m == nil {
		return errors.New("measurement is not found in schema meta")
	}
	return nil
}
