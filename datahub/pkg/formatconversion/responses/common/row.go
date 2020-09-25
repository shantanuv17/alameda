package common

import (
	"github.com/golang/protobuf/ptypes"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/api/datahub/common"
)

func NewRows(rows []*DBCommon.Row) []*common.Row {
	if rows == nil {
		return nil
	}
	rs := make([]*common.Row, 0)
	for _, row := range rows {
		rs = append(rs, NewRow(row))
	}
	return rs
}

func NewRow(row *DBCommon.Row) *common.Row {
	if row == nil {
		return nil
	}
	r := common.Row{}
	if row.Time != nil {
		timestamp, _ := ptypes.TimestampProto(*row.Time)
		r.Time = timestamp
	}
	r.Values = row.Values
	return &r
}
