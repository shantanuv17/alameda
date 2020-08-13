package datahub

import (
	"github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/datahub/keycodes"
	"google.golang.org/grpc"
	"sync"
)

var (
	scope = log.RegisterScope("datahub-client", "datahub client library", 0)
)

type Client struct {
	datahub.DatahubServiceClient
	keycodes.KeycodesServiceClient

	RWLock     *sync.RWMutex
	Address    string

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
