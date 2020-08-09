package utils

import (
	"context"
	"reflect"
	"testing"
	"time"

	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_data "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	datahub_schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
)

func TestReadData(t *testing.T) {
	datahubAddr := "127.0.0.1:50050"
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, datahubAddr, grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(uint(5)))))

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
					{
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

func TestGetGranularityStr(t *testing.T) {
	type args struct {
		granularitySec int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "30",
			args: args{
				granularitySec: 30,
			},
			want: "30s",
		},
		{
			name: "60",
			args: args{
				granularitySec: 60,
			},
			want: "1m",
		},
		{
			name: "120",
			args: args{
				granularitySec: 120,
			},
			want: "2m",
		},
		{
			name: "180",
			args: args{
				granularitySec: 180,
			},
			want: "3m",
		},
		{
			name: "3600",
			args: args{
				granularitySec: 3600,
			},
			want: "1h",
		},
		{
			name: "21600",
			args: args{
				granularitySec: 21600,
			},
			want: "6h",
		},
		{
			name: "86400",
			args: args{
				granularitySec: 86400,
			},
			want: "24h",
		},
	}
	for _, ttt := range tests {
		tt := ttt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetGranularityStr(tt.args.granularitySec); got != tt.want {
				t.Errorf("GetGranularityStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGranularitySec(t *testing.T) {
	type args struct {
		granularityStr string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "30s",
			args: args{
				granularityStr: "30s",
			},
			want: 30,
		},
		{
			name: "1m",
			args: args{
				granularityStr: "1m",
			},
			want: 60,
		},
		{
			name: "2m",
			args: args{
				granularityStr: "2m",
			},
			want: 120,
		},
		{
			name: "3m",
			args: args{
				granularityStr: "3m",
			},
			want: 180,
		},
		{
			name: "1h",
			args: args{
				granularityStr: "1h",
			},
			want: 3600,
		},
		{
			name: "6h",
			args: args{
				granularityStr: "6h",
			},
			want: 21600,
		},
		{
			name: "24h",
			args: args{
				granularityStr: "24h",
			},
			want: 86400,
		},
	}
	for _, ttt := range tests {
		tt := ttt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetGranularitySec(tt.args.granularityStr); got != tt.want {
				t.Errorf("GetGranularitySec() = %v, want %v", got, tt.want)
			}
		})
	}
}
