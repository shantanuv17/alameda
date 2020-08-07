package datahub

import (
	"context"
	"errors"
	Entities "github.com/containers-ai/alameda/datahub/pkg/entities"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"reflect"
	"sync"
	"time"

)

var (
	scope = log.RegisterScope("datahub-client", "datahub client library", 0)
)

type Client struct {
	datahub.DatahubServiceClient
	Connection *grpc.ClientConn
	RWLock     *sync.RWMutex
	Address    string
}

func NewClient(address string) *Client {
	client := Client{}
	client.Address = address
	client.RWLock = new(sync.RWMutex)

	// Create a client connection to datahub
	if err := client.Reconnect(5); err != nil {
		scope.Errorf("failed to dial to datahub via address(%s): %s", address, err.Error())
		return nil
	}

	return &client
}

func (p *Client) Reconnect(retry int) error {
	// Check if connection is alive first
	if p.IsAlive() {
		scope.Info("connection to datahub is alive, no necessary to reconnect")
		return nil
	}

	// Get WRITE lock
	p.RWLock.Lock()
	defer p.RWLock.Unlock()

	// Do close connection before connecting to datahub
	if err := p.Close(); err != nil {
		scope.Errorf("failed to close reconnect of datahub: %s", err.Error())
		return err
	}

	for i := 0; i < retry; i++ {
		// Create a client connection to datahub
		conn, err := grpc.Dial(p.Address,
			grpc.WithBlock(),
			grpc.WithTimeout(30 * time.Second),
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(grpcretry.UnaryClientInterceptor(grpcretry.WithMax(uint(3)))),
		)

		if err == nil {
			scope.Info("successfully create connection to datahub")
			p.Connection = conn
			p.DatahubServiceClient = datahub.NewDatahubServiceClient(p.Connection)
			break
		}

		scope.Errorf("failed to create connection to datahub, try again: %s", err.Error())
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (p *Client) Close() error {
	if p.Connection != nil {
		if err := p.Connection.Close(); err != nil {
			scope.Error(err.Error())
			return err
		}
	}
	p.Connection = nil
	return nil
}

func (p *Client) IsAlive() bool {
	// Get READ lock
	p.RWLock.RLock()
	defer p.RWLock.RUnlock()

	if p.Connection != nil {
		state := p.Connection.GetState()
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

func (p *Client) Create(entities interface{}, fields ...string) error {
	request := NewWriteDataRequest(entities, fields)
	status, err := p.WriteData(context.Background(), request)

	// Check error
	if err != nil {
		return err
	}

	// Check response status code
	if status.Code != 0 {
		return errors.New(status.GetMessage())
	}

	return nil
}

func (p *Client) List(entities interface{}, opts ...Option) error {
	request := NewReadDataRequest(entities, nil, nil, nil, opts...)
	response, err := p.ReadData(context.Background(), request)

	// Check error
	if err != nil {
		return err
	}

	// Check response status code
	if response.Status.Code != 0 {
		return errors.New(response.Status.GetMessage())
	}

	// Copy results
	DeepCopyEntity(entities, response.Data)

	return nil
}

func (p *Client) ListTS(entities interface{}, timeRange *TimeRange, function *Function, fields []string, opts ...Option) error {
	request := NewReadDataRequest(entities, fields, timeRange, function, opts...)
	response, err := p.ReadData(context.Background(), request)

	// Check error
	if err != nil {
		return err
	}

	// Check response status code
	if response.Status.Code != 0 {
		return errors.New(response.Status.GetMessage())
	}

	// Copy results
	DeepCopyEntity(entities, response.Data)

	return nil
}

// Delete by tags
func (p *Client) Delete(entities interface{}) error {
	opts := make([]Option, 0)

	values := reflect.ValueOf(entities).Elem()

	// If length of entities list is ZERO which means to delete nothing
	if values.Len() == 0 {
		return nil
	}

	// Iterate the entities to find all the tags
	for i := 0; i < values.Len(); i++ {
		entity := values.Index(i).Interface()
		datahubEntity := values.Index(i).Field(0).Interface().(Entities.DatahubEntity)
		tags := datahubEntity.TagNames(entity)
		opts = append(opts, Option{Entity: entity, Fields: tags})
	}

	request := NewDeleteDataRequest(entities, opts...)
	status, err := p.DeleteData(context.Background(), request)

	// Check error
	if err != nil {
		return err
	}

	// Check response status code
	if status.Code != 0 {
		return errors.New(status.GetMessage())
	}

	return nil
}

// Entity is indicator, delete by options
func (p *Client) DeleteByOpts(entity interface{}, opts ...Option) error {
	request := NewDeleteDataRequest(entity, opts...)
	status, err := p.DeleteData(context.Background(), request)

	// Check error
	if err != nil {
		return err
	}

	// Check response status code
	if status.Code != 0 {
		return errors.New(status.GetMessage())
	}

	return nil
}
