package v1alpha1

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/alameda/pkg/database/influxdb"
	"prophetstor.com/alameda/pkg/database/prometheus"
	"prophetstor.com/api/datahub/common"
	"prophetstor.com/api/datahub/rawdata"
)

// Read rawdata from database
func (s *ServiceV1alpha1) ReadRawdata(ctx context.Context, in *rawdata.ReadRawdataRequest) (*rawdata.ReadRawdataResponse, error) {
	scope.Debug("Request received from ReadRawdata grpc function")

	var (
		err   error
		rData = make([]*rawdata.ReadRawdata, 0)
	)

	switch in.GetDatabaseType() {
	case common.DatabaseType_INFLUXDB:
		rData, err = influxdb.ReadRawdata(s.Config.InfluxDB, in.GetQueries())
	case common.DatabaseType_PROMETHEUS:
		rData, err = prometheus.ReadRawdata(s.Config.Prometheus, in.GetQueries())
	default:
		err = errors.New(fmt.Sprintf("database type(%s) is not supported", common.DatabaseType_name[int32(in.GetDatabaseType())]))
	}

	if err != nil {
		scope.Errorf("api ReadRawdata failed: %v", err)
		response := &rawdata.ReadRawdataResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
			Rawdata: rData,
		}
		return response, err
	}

	response := &rawdata.ReadRawdataResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Rawdata: rData,
	}

	return response, nil
}

// Write rawdata to database
func (s *ServiceV1alpha1) WriteRawdata(ctx context.Context, in *rawdata.WriteRawdataRequest) (*status.Status, error) {
	scope.Debug("Request received from WriteRawdata grpc function")

	var (
		err error
	)

	switch in.GetDatabaseType() {
	case common.DatabaseType_INFLUXDB:
		err = influxdb.WriteRawdata(s.Config.InfluxDB, in.GetRawdata())
	case common.DatabaseType_PROMETHEUS:
		err = errors.New(fmt.Sprintf("database type(%s) is not supported yet", common.DatabaseType_name[int32(in.GetDatabaseType())]))
	default:
		err = errors.New(fmt.Sprintf("database type(%s) is not supported", common.DatabaseType_name[int32(in.GetDatabaseType())]))
	}

	if err != nil {
		scope.Errorf("api WriteRawdata failed: %v", err)
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, err
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}
