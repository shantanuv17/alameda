package datahub

import (
	"errors"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/datahub/keycodes"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"time"
)

func (p *Client) Connect(retry, timeout int) error {
	// Check if connection is alive first
	if p.IsAlive() {
		scope.Info("connection to datahub is still alive, no necessary to connect again")
		return nil
	}

	// Get WRITE lock
	p.RWLock.Lock()
	defer p.RWLock.Unlock()

	// Do close connection before connecting to datahub
	if err := p.Close(); err != nil {
		scope.Errorf("failed to close connection to datahub: %s", err.Error())
		return err
	}

	// Create a client connection to datahub
	conn, err := grpc.Dial(p.Address,
		grpc.WithBlock(),
		grpc.WithTimeout(time.Duration(timeout) * time.Second),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(grpcretry.WithMax(uint(retry)))),
	)

	if err != nil {
		scope.Errorf("failed to dial to datahub via address(%s): %s", p.Address, err.Error())
		return err
	}

	scope.Info("successfully dial to to datahub")
	p.connection = conn
	p.DatahubServiceClient = datahub.NewDatahubServiceClient(p.connection)
	p.KeycodesServiceClient = keycodes.NewKeycodesServiceClient(p.connection)

	return nil
}

func (p *Client) Close() error {
	if p.connection != nil {
		if err := p.connection.Close(); err != nil {
			scope.Error(err.Error())
			return err
		}
	}
	p.connection = nil
	p.DatahubServiceClient = nil
	p.KeycodesServiceClient = nil
	return nil
}

func (p *Client) IsAlive() bool {
	// Get READ lock
	p.RWLock.RLock()
	defer p.RWLock.RUnlock()

	if p.connection != nil {
		state := p.connection.GetState()
		switch state {
		case connectivity.Idle:
			return true
		case connectivity.Connecting:
			return true
		case connectivity.Ready:
			return true
		case connectivity.TransientFailure:
			return false
		case connectivity.Shutdown:
			return false
		default:
			scope.Errorf("unknown connectivity state: %d", state)
			return false
		}
	}
	return false
}

func (p *Client) CheckConnection() error {
	if !p.IsAlive() {
		if err := p.Connect(3, 30); err != nil {
			scope.Errorf("failed to connect to datahub via address(%s): %s", p.Address, err.Error())
			return errors.New("failed to connect to datahub")
		}
	}
	return nil
}
