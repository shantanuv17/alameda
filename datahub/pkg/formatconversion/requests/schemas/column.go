package schemas

import (
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxSchema "prophetstor.com/alameda/pkg/database/influxdb/schemas"
	ApiSchema "prophetstor.com/api/datahub/schemas"
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
