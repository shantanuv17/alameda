package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) CreateApplications(request *resources.CreateApplicationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateApplications(context.Background(), request)
}

func (p *Client) ListApplications(request *resources.ListApplicationsRequest) (*resources.ListApplicationsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListApplications(context.Background(), request)
}

func (p *Client) DeleteApplications(request *resources.DeleteApplicationsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteApplications(context.Background(), request)
}

func (p *Client) CreateClusters(request *resources.CreateClustersRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateClusters(context.Background(), request)
}

func (p *Client) ListClusters(request *resources.ListClustersRequest) (*resources.ListClustersResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListClusters(context.Background(), request)
}

func (p *Client) DeleteClusters(request *resources.DeleteClustersRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteClusters(context.Background(), request)
}

func (p *Client) CreateControllers(request *resources.CreateControllersRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateControllers(context.Background(), request)
}

func (p *Client) ListControllers(request *resources.ListControllersRequest) (*resources.ListControllersResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListControllers(context.Background(), request)
}

func (p *Client) DeleteControllers(request *resources.DeleteControllersRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteControllers(context.Background(), request)
}

func (p *Client) CreateNamespaces(request *resources.CreateNamespacesRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNamespaces(context.Background(), request)
}

func (p *Client) ListNamespaces(request *resources.ListNamespacesRequest) (*resources.ListNamespacesResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNamespaces(context.Background(), request)
}

func (p *Client) DeleteNamespaces(request *resources.DeleteNamespacesRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteNamespaces(context.Background(), request)
}

func (p *Client) CreateNodes(request *resources.CreateNodesRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateNodes(context.Background(), request)
}

func (p *Client) ListNodes(request *resources.ListNodesRequest) (*resources.ListNodesResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListNodes(context.Background(), request)
}

func (p *Client) DeleteNodes(request *resources.DeleteNodesRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeleteNodes(context.Background(), request)
}

func (p *Client) CreatePods(request *resources.CreatePodsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreatePods(context.Background(), request)
}

func (p *Client) ListPods(request *resources.ListPodsRequest) (*resources.ListPodsResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListPods(context.Background(), request)
}

func (p *Client) DeletePods(request *resources.DeletePodsRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.DeletePods(context.Background(), request)
}
