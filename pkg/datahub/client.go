package datahub

import (
	"context"
	"errors"
	Entities "github.com/containers-ai/alameda/datahub/pkg/entities"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"reflect"
	"time"
)

var (
	scope = log.RegisterScope("datahub-client", "datahub client library", 0)
)

type Client struct {
	datahub.DatahubServiceClient
	Connection *grpc.ClientConn
	Address    string
}

func NewClient(address string) *Client {
	// Create a client connection to datahub
	conn, err := grpc.Dial(address,
		grpc.WithBlock(),
		grpc.WithTimeout(30*time.Second),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(uint(3)))),
	)

	if err != nil {
		scope.Errorf("failed to dial to datahub via address(%s): %s", address, err.Error())
		return nil
	}

	// Create datahub service client and initialize member variable
	client := Client{DatahubServiceClient: datahub.NewDatahubServiceClient(conn)}
	client.Address = address
	client.Connection = conn

	return &client
}

func (p *Client) Close() error {
	return p.Connection.Close()
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
		opts = append(opts, Option{Entity: entity, Fields: tags,})
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
