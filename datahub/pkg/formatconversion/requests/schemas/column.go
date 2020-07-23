package schemas

import (
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	InfluxSchema "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewColumn(column *ApiSchema.Column) *InfluxSchema.Column {
	if column != nil {
		c := InfluxSchema.Column{}
		c.Name = column.Name
		c.Required = column.Required
		c.ColumnType = InfluxSchema.ColumnType(column.ColumnType)
		c.DataType = DBCommon.DataType(column.DataType)
		return &c
	}
	return nil
}
