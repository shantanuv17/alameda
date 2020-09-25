// +build integration

package client

import (
	"testing"

	"prophetstor.com/alameda/internal/pkg/message-queue/kafka"
)

func TestClose(t *testing.T) {

	c, err := NewClient(kafka.Config{
		Addresses: []string{"localhost:9092"},
		Version:   "2.4.0",
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
