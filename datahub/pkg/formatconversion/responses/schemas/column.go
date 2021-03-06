package schemas

import (
	InfluxSchema "github.com/containers-ai/alameda/pkg/database/influxdb/schemas"
	ApiCommon "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	ApiSchema "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func NewColumn(column *InfluxSchema.Column) *ApiSchema.Column {
	if column != nil {
		c := ApiSchema.Column{}
		c.Name = column.Name
		c.Required = column.Required
		c.ColumnType = ApiCommon.ColumnType(column.ColumnType)
		c.DataType = ApiCommon.DataType(column.DataType)
		return &c
	}
	return nil
}
