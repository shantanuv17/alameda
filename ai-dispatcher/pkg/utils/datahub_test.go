package utils

import (
	"reflect"
	"testing"
	"time"

	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_data "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
)

func TestReadData(t *testing.T) {
	datahubAddr := "127.0.0.1:50050"
	conn, err := grpc.Dial(datahubAddr, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(5))))

	queryStartT := time.Now().AddDate(0, 0, -1).Unix()
	metricReadData := []*datahub_data.ReadData{}
	for _, node := range []string{
		"okd4-tsztm-worker-0-dxcqp",
		"okd4-tsztm-worker-0-hdfxb",
		"okd4-tsztm-worker-0-vdtr7",
	} {
		metricReadData = append(metricReadData, &datahub_data.ReadData{
			MetricType: datahub_common.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
			QueryCondition: &datahub_common.QueryCondition{
				Selects: []string{"value"},
				Order:   datahub_common.QueryCondition_DESC,
				Groups:  []string{"cluster_name", "name"},
				TimeRange: &datahub_common.TimeRange{
					Step: &duration.Duration{
						Seconds: 60,
					},
					StartTime: &timestamp.Timestamp{
						Seconds: queryStartT,
					},
					AggregateFunction: datahub_common.TimeRange_AVG,
				},
				WhereCondition: []*datahub_common.Condition{
					&datahub_common.Condition{
						Keys:      []string{"cluster_name", "name"},
						Values:    []string{"7ebc85ee-9c96-4940-a76e-dc7a60e1d2f2", node},
						Operators: []string{"=", "="},
						Types: []datahub_common.DataType{
							datahub_common.DataType_DATATYPE_STRING,
							datahub_common.DataType_DATATYPE_STRING,
						},
					},
				},
			},
		})
	}
	if err != nil {
		t.Errorf("datahub connection error = %v", err)
		return
	}

	type args struct {
		datahubServiceClnt datahub_v1alpha1.DatahubServiceClient
		schemaMeta         *datahub_schemas.SchemaMeta
		readData           []*datahub_data.ReadData
	}
	tests := []struct {
		name    string
		args    args
		want    *datahub_data.ReadDataResponse
		wantErr bool
	}{
		{
			name: "node metrics read data",
			args: args{
				datahubServiceClnt: datahub_v1alpha1.NewDatahubServiceClient(conn),
				schemaMeta: &datahub_schemas.SchemaMeta{
					Scope:    datahub_schemas.Scope_SCOPE_METRIC,
					Category: "cluster_status",
					Type:     "node",
				},
				readData: metricReadData,
			},
		},
	}
	for _, ttt := range tests {
		tt := ttt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ReadData(tt.args.datahubServiceClnt, tt.args.schemaMeta, tt.args.readData)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() = %v, want %v", got, tt.want)
			}
		})
	}
}
