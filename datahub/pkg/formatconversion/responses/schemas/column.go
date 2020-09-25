package schemas

import (
	InfluxSchema "prophetstor.com/alameda/pkg/database/influxdb/schemas"
	ApiCommon "prophetstor.com/api/datahub/common"
	ApiSchema "prophetstor.com/api/datahub/schemas"
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
