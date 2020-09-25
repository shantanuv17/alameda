package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/gpu"
)

func (p *Client) ListGpus(request *gpu.ListGpusRequest) (*gpu.ListGpusResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListGpus(context.Background(), request)
}

func (p *Client) ListGpuMetrics(request *gpu.ListGpuMetricsRequest) (*gpu.ListGpuMetricsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListGpuMetrics(context.Background(), request)
}

func (p *Client) ListGpuPredictions(request *gpu.ListGpuPredictionsRequest) (*gpu.ListGpuPredictionsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListGpuPredictions(context.Background(), request)
}

func (p *Client) CreateGpuPredictions(request *gpu.CreateGpuPredictionsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateGpuPredictions(context.Background(), request)
}
