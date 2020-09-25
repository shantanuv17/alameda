package data

import (
	"prophetstor.com/alameda/datahub/pkg/dao/interfaces/data/types"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/responses/common"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/responses/enumconv"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/responses/schemas"
	"prophetstor.com/api/datahub/data"
)

func NewData(readData *types.Data) *data.Data {
	if readData == nil {
		return nil
	}
	rawdata := data.Data{}
	rawdata.SchemaMeta = schemas.NewSchemaMeta(readData.SchemaMeta)
	rawdata.Rawdata = NewRawdata(readData.Rawdata)
	return &rawdata
}

func NewRawdata(rawdata []*types.Rawdata) []*data.Rawdata {
	if rawdata == nil {
		return nil
	}
	rawdataList := make([]*data.Rawdata, 0)
	for _, d := range rawdata {
		rd := data.Rawdata{}
		rd.Measurement = d.Measurement
		rd.MetricType = enumconv.MetricTypeNameMap[d.MetricType]
		rd.ResourceBoundary = enumconv.ResourceBoundaryNameMap[d.Boundary]
		rd.ResourceQuota = enumconv.ResourceQuotaNameMap[d.Quota]
		for _, group := range d.Groups {
			rd.Groups = append(rd.Groups, common.NewGroup(group))
		}
		rawdataList = append(rawdataList, &rd)
	}
	return rawdataList
}
