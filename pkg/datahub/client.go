package datahub

import (
	"sync"

	"google.golang.org/grpc"
	"prophetstor.com/alameda/pkg/utils/log"
	"prophetstor.com/api/datahub"
)

var (
	scope = log.RegisterScope("datahub-client", "datahub client library", 0)
)

type Client struct {
	datahub.DatahubServiceClient

	RWLock  *sync.RWMutex
	Address string

	connection *grpc.ClientConn
}

func NewClient(address string) *Client {
	client := Client{}
	client.Address = address
	client.RWLock = new(sync.RWMutex)

	// Connect to datahub
	if err := client.Connect(3, 30); err != nil {
		scope.Errorf("failed to connect to datahub: %s", err.Error())
		return nil
	}

	return &client
}
