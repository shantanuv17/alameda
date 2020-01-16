// +build integration

package client

import (
	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"
	"testing"
)

func TestClose(t *testing.T) {

	c, err := NewClient(kafka.Config{
		Addresses: []string{"localhost:9092"},
		Version:   "0.11.0.0",
	})
	if err != nil {
		panic(err)
	}

	if err := c.Open(); err != nil {
		panic(err)
	}

	if err := c.Close(); err != nil {
		panic(err)
	}
}
