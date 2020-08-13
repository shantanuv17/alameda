package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/predictions"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) CreateApplicationPredictions(request *predictions.CreateApplicationPredictionsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateApplicationPredictions(context.Background(), request)
}

func (p *Client) ListApplicationPredictions(request *predictions.ListApplicationPredictionsRequest) (*predictions.ListApplicationPredictionsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListApplicationPredictions(context.Background(), request)
}

func (p *Client) CreateClusterPredictions(request *predictions.CreateClusterPredictionsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateClusterPredictions(context.Background(), request)
}

func (p *Client) ListClusterPredictions(request *predictions.ListClusterPredictionsRequest) (*predictions.ListClusterPredictionsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListClusterPredictions(context.Background(), request)
}

func (p *Client) CreateControllerPredictions(request *predictions.CreateControllerPredictionsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateControllerPredictions(context.Background(), request)
}

func (p *Client) ListControllerPredictions(request *predictions.ListControllerPredictionsRequest) (*predictions.ListControllerPredictionsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListControllerPredictions(context.Background(), request)
}

func (p *Client) CreateNamespacePredictions(request *predictions.CreateNamespacePredictionsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNamespacePredictions(context.Background(), request)
}

func (p *Client) ListNamespacePredictions(request *predictions.ListNamespacePredictionsRequest) (*predictions.ListNamespacePredictionsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNamespacePredictions(context.Background(), request)
}

func (p *Client) CreateNodePredictions(request *predictions.CreateNodePredictionsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNodePredictions(context.Background(), request)
}

func (p *Client) ListNodePredictions(request *predictions.ListNodePredictionsRequest) (*predictions.ListNodePredictionsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNodePredictions(context.Background(), request)
}

func (p *Client) CreatePodPredictions(request *predictions.CreatePodPredictionsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreatePodPredictions(context.Background(), request)
}

func (p *Client) ListPodPredictions(request *predictions.ListPodPredictionsRequest) (*predictions.ListPodPredictionsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListPodPredictions(context.Background(), request)
}
