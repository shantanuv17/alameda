package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) CreateApplicationRecommendations(request *recommendations.CreateApplicationRecommendationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateApplicationRecommendations(context.Background(), request)
}

func (p *Client) ListApplicationRecommendations(request *recommendations.ListApplicationRecommendationsRequest) (*recommendations.ListApplicationRecommendationsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListApplicationRecommendations(context.Background(), request)
}

func (p *Client) CreateClusterRecommendations(request *recommendations.CreateClusterRecommendationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateClusterRecommendations(context.Background(), request)
}

func (p *Client) ListClusterRecommendations(request *recommendations.ListClusterRecommendationsRequest) (*recommendations.ListClusterRecommendationsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListClusterRecommendations(context.Background(), request)
}

func (p *Client) CreateControllerRecommendations(request *recommendations.CreateControllerRecommendationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateControllerRecommendations(context.Background(), request)
}

func (p *Client) ListControllerRecommendations(request *recommendations.ListControllerRecommendationsRequest) (*recommendations.ListControllerRecommendationsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListControllerRecommendations(context.Background(), request)
}

func (p *Client) CreateNamespaceRecommendations(request *recommendations.CreateNamespaceRecommendationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNamespaceRecommendations(context.Background(), request)
}

func (p *Client) ListNamespaceRecommendations(request *recommendations.ListNamespaceRecommendationsRequest) (*recommendations.ListNamespaceRecommendationsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNamespaceRecommendations(context.Background(), request)
}

func (p *Client) CreateNodeRecommendations(request *recommendations.CreateNodeRecommendationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNodeRecommendations(context.Background(), request)
}

func (p *Client) ListNodeRecommendations(request *recommendations.ListNodeRecommendationsRequest) (*recommendations.ListNodeRecommendationsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNodeRecommendations(context.Background(), request)
}

func (p *Client) CreatePodRecommendations(request *recommendations.CreatePodRecommendationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreatePodRecommendations(context.Background(), request)
}

func (p *Client) ListPodRecommendations(request *recommendations.ListPodRecommendationsRequest) (*recommendations.ListPodRecommendationsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListPodRecommendations(context.Background(), request)
}
