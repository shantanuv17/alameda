package datahub

import (
	"context"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/scores"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (p *Client) CreateSimulatedSchedulingScores(request *scores.CreateSimulatedSchedulingScoresRequest) (*status.Status, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.CreateSimulatedSchedulingScores(context.Background(), request)
}

func (p *Client) ListSimulatedSchedulingScores(request *scores.ListSimulatedSchedulingScoresRequest) (*scores.ListSimulatedSchedulingScoresResponse, error) {
	if err := p.CheckConnection(); err != nil {
		return nil, err
	}
	return p.DatahubServiceClient.ListSimulatedSchedulingScores(context.Background(), request)
}
