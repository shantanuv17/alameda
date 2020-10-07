package datahub

import (
	"reflect"
	"sync"
	"testing"

	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"prophetstor.com/api/datahub"
	"prophetstor.com/api/datahub/configs"
)

func TestClient_CreateConfigs(t *testing.T) {
	type fields struct {
		DatahubServiceClient datahub.DatahubServiceClient
		RWLock               *sync.RWMutex
		Address              string
		connection           *grpc.ClientConn
	}
	type args struct {
		request *configs.CreateConfigsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *status.Status
		wantErr bool
	}{
		{
			args: args{
				request: &configs.CreateConfigsRequest{
					Configs: []*configs.Config{
						{Kind: &configs.Config_DetectionConfig{}},
						{Kind: &configs.Config_ScalerConfig{}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Client{
				DatahubServiceClient: tt.fields.DatahubServiceClient,
				RWLock:               tt.fields.RWLock,
				Address:              tt.fields.Address,
				connection:           tt.fields.connection,
			}
			got, err := p.CreateConfigs(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.CreateConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}
