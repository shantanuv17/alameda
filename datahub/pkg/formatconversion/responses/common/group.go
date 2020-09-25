package common

import (
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	ApiCommon "prophetstor.com/api/datahub/common"
)

func NewGroup(group *DBCommon.Group) *ApiCommon.Group {
	if group == nil {
		return nil
	}
	g := ApiCommon.Group{}
	g.Columns = group.Columns
	g.Rows = NewRows(group.Rows)
	return &g
}
