package kafka

import (
	"context"
)

type Client interface {
	Open() error
	Close() error
	ListTopics(ctx context.Context) ([]string, error)
	ListConsumerGroups(ctx context.Context) ([]string, error)
	ListConsumeTopics(ctx context.Context, consumerGroup string) ([]string, error)
	ListTopicsPartitionCounts(ctx context.Context,topics []string) (map[string]int, error)
}
