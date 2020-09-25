package common

import (
	"github.com/golang/protobuf/ptypes"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	"prophetstor.com/api/datahub/common"
)

func NewRows(rows []*common.Row) []*DBCommon.Row {
	if rows == nil {
		return nil
	}
	rs := make([]*DBCommon.Row, 0)
	for _, row := range rows {
		rs = append(rs, NewRow(row))
	}
	return rs
}

func NewRow(row *common.Row) *DBCommon.Row {
	if row == nil {
		return nil
	}
	r := DBCommon.Row{}
	if row.GetTime() != nil {
		timestamp, _ := ptypes.Timestamp(row.GetTime())
		r.Time = &timestamp
	}
	r.Values = row.GetValues()
	return &r
}
