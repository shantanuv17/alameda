package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/metrics"
)

func (p *Client) CreateApplicationMetrics(request *metrics.CreateApplicationMetricsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateApplicationMetrics(context.Background(), request)
}

func (p *Client) ListApplicationMetrics(request *metrics.ListApplicationMetricsRequest) (*metrics.ListApplicationMetricsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListApplicationMetrics(context.Background(), request)
}

func (p *Client) CreateClusterMetrics(request *metrics.CreateClusterMetricsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateClusterMetrics(context.Background(), request)
}

func (p *Client) ListClusterMetrics(request *metrics.ListClusterMetricsRequest) (*metrics.ListClusterMetricsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListClusterMetrics(context.Background(), request)
}

func (p *Client) CreateControllerMetrics(request *metrics.CreateControllerMetricsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateControllerMetrics(context.Background(), request)
}

func (p *Client) ListControllerMetrics(request *metrics.ListControllerMetricsRequest) (*metrics.ListControllerMetricsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListControllerMetrics(context.Background(), request)
}

func (p *Client) CreateNamespaceMetrics(request *metrics.CreateNamespaceMetricsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNamespaceMetrics(context.Background(), request)
}

func (p *Client) ListNamespaceMetrics(request *metrics.ListNamespaceMetricsRequest) (*metrics.ListNamespaceMetricsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNamespaceMetrics(context.Background(), request)
}

func (p *Client) CreateNodeMetrics(request *metrics.CreateNodeMetricsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNodeMetrics(context.Background(), request)
}

func (p *Client) ListNodeMetrics(request *metrics.ListNodeMetricsRequest) (*metrics.ListNodeMetricsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNodeMetrics(context.Background(), request)
}

func (p *Client) CreatePodMetrics(request *metrics.CreatePodMetricsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreatePodMetrics(context.Background(), request)
}

func (p *Client) ListPodMetrics(request *metrics.ListPodMetricsRequest) (*metrics.ListPodMetricsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListPodMetrics(context.Background(), request)
}
