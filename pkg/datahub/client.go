package datahub

import (
	"context"
	"errors"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"google.golang.org/grpc"
	"time"
)

var (
	scope = log.RegisterScope("datahub-client", "datahub client library", 0)
)

type Client struct {
	Address string
}

func NewClient(address string) *Client {
	client := Client{}
	client.Address = address
	return &client
}

func (p *Client) Create(entities interface{}, fields []string) error {
	// Create connection to datahub
	conn, err := grpc.Dial(p.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Write data to datahub
	client := datahub.NewDatahubServiceClient(conn)
	request := NewWriteDataRequest(entities, fields)
	status, err := client.WriteData(context.Background(), request)

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
	// Create connection to datahub
	conn, err := grpc.Dial(p.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Read data from datahub
	client := datahub.NewDatahubServiceClient(conn)
	request := NewReadDataRequest(entities, nil, nil, opts...)
	response, err := client.ReadData(context.Background(), request)

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

func (p *Client) ListTS(entities interface{}, startTime, endTime *time.Time, opts ...Option) error {
	// Create connection to datahub
	conn, err := grpc.Dial(p.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Read data from datahub
	client := datahub.NewDatahubServiceClient(conn)
	request := NewReadDataRequest(entities, startTime, endTime, opts...)
	response, err := client.ReadData(context.Background(), request)

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

func (p *Client) Delete(entities interface{}, opts ...Option) error {
	// Create connection to datahub
	conn, err := grpc.Dial(p.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Write data to datahub
	client := datahub.NewDatahubServiceClient(conn)
	request := NewDeleteDataRequest(entities, opts...)
	status, err := client.DeleteData(context.Background(), request)

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
