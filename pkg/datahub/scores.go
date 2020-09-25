package datahub

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"
	"prophetstor.com/api/datahub/scores"
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
