package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/plannings"
)

func (p *Client) CreateApplicationPlannings(request *plannings.CreateApplicationPlanningsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateApplicationPlannings(context.Background(), request)
}

func (p *Client) ListApplicationPlannings(request *plannings.ListApplicationPlanningsRequest) (*plannings.ListApplicationPlanningsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListApplicationPlannings(context.Background(), request)
}

func (p *Client) CreateClusterPlannings(request *plannings.CreateClusterPlanningsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateClusterPlannings(context.Background(), request)
}

func (p *Client) ListClusterPlannings(request *plannings.ListClusterPlanningsRequest) (*plannings.ListClusterPlanningsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListClusterPlannings(context.Background(), request)
}

func (p *Client) CreateControllerPlannings(request *plannings.CreateControllerPlanningsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateControllerPlannings(context.Background(), request)
}

func (p *Client) ListControllerPlannings(request *plannings.ListControllerPlanningsRequest) (*plannings.ListControllerPlanningsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListControllerPlannings(context.Background(), request)
}

func (p *Client) CreateNamespacePlannings(request *plannings.CreateNamespacePlanningsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNamespacePlannings(context.Background(), request)
}

func (p *Client) ListNamespacePlannings(request *plannings.ListNamespacePlanningsRequest) (*plannings.ListNamespacePlanningsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNamespacePlannings(context.Background(), request)
}

func (p *Client) CreateNodePlannings(request *plannings.CreateNodePlanningsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNodePlannings(context.Background(), request)
}

func (p *Client) ListNodePlannings(request *plannings.ListNodePlanningsRequest) (*plannings.ListNodePlanningsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNodePlannings(context.Background(), request)
}

func (p *Client) CreatePodPlannings(request *plannings.CreatePodPlanningsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreatePodPlannings(context.Background(), request)
}

func (p *Client) ListPodPlannings(request *plannings.ListPodPlanningsRequest) (*plannings.ListPodPlanningsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListPodPlannings(context.Background(), request)
}
