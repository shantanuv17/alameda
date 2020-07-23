package common

import (
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/golang/protobuf/ptypes"
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
