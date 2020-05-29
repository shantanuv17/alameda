package datahub

import (
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"time"
)

func NewWriteDataRequest(entities interface{}, fields []string) *data.WriteDataRequest {
	request := data.WriteDataRequest{}
	request.SchemaMeta = NewSchemaMeta(entities)
	request.WriteData = append(request.WriteData, NewWriteData(entities, fields))
	return &request
}

func NewReadDataRequest(entities interface{}, startTime, endTime *time.Time, opts ...Option) *data.ReadDataRequest {
	request := data.ReadDataRequest{}
	request.SchemaMeta = NewSchemaMeta(entities)
	request.ReadData = append(request.ReadData, NewReadData(entities, startTime, endTime, opts...))
	return &request
}

func NewDeleteDataRequest(entities interface{}, opts ...Option) *data.DeleteDataRequest {
	request := data.DeleteDataRequest{}
	request.SchemaMeta = NewSchemaMeta(entities)
	request.DeleteData = append(request.DeleteData, NewDeleteData(entities, opts...))
	return &request
}
