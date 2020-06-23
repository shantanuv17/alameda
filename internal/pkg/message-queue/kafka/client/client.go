package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/containers-ai/alameda/internal/pkg/message-queue/kafka"

	"github.com/Shopify/sarama"
)

var (
	defaultDialTimeout = 10 * time.Second
	defaultKeepAlive   = 300 * time.Second
)

type client struct {
	brokerAddresses []string
	config          sarama.Config

	lock   sync.Mutex
	admin  sarama.ClusterAdmin
	client sarama.Client
}

// NewClient returns implementation for internal kafka client interface,
// The connection od the client will be lazily initialization.
func NewClient(config kafka.Config) (kafka.Client, error) {
	config = setConfigDefaults(config)

	cfg := sarama.NewConfig()
	cfg.Net.DialTimeout = *config.DialTimeout
	cfg.Net.KeepAlive = *config.KeepAlive
	version, err := sarama.ParseKafkaVersion(config.Version)
	if err != nil {
		return nil, errors.Wrap(err, "parse kafka version failed")
	}
	cfg.Version = version
	if config.TLS.Enabled {
		cfg.Net.TLS.Enable = true
		cfg.Net.TLS.Config = &tls.Config{
			InsecureSkipVerify: config.TLS.InsecureSkipVerify,
		}
	}
	if config.SASL != nil {
		cfg.Net.SASL.Enable = config.SASL.Enabled
		cfg.Net.SASL.User = config.SASL.Username
		cfg.Net.SASL.Password = config.SASL.Password
	}

	return &client{
		brokerAddresses: config.Addresses,
		config:          *cfg,
	}, nil
}

func (c *client) Open() error {
	if c.admin != nil {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	client, err := sarama.NewClient(c.brokerAddresses, &c.config)
	if err != nil {
		return errors.Wrap(err, "new kafka client failed")
	}
	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		return errors.Wrap(err, "new kafka clusterAdmin failed")
	}
	c.admin = admin
	c.client = client
	return nil
}

func (c *client) Close() error {
	if c.admin == nil {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	if err := c.admin.Close(); err != nil {
		return errors.Wrap(err, "close admin failed")
	}
	c.admin = nil
	if err := c.client.Close(); err != nil {
		return errors.Wrap(err, "close client failed")
	}
	c.client = nil
	return nil
}

func (c *client) ListTopics(ctx context.Context) ([]string, error) {
	if err := c.Open(); err != nil {
		return nil, errors.Wrap(err, "open client failed")
	}
	topicsDetail, err := c.admin.ListTopics()
	if err = c.handleError(err); err != nil {
		return nil, errors.Wrap(err, "list topics failed")
	}

	topics := make([]string, 0, len(topicsDetail))
	for topic := range topicsDetail {
		topics = append(topics, topic)
	}
	return topics, nil
}

func (c *client) ListConsumerGroups(ctx context.Context) ([]string, error) {
	if err := c.Open(); err != nil {
		return nil, errors.Wrap(err, "open client failed")
	}
	consumerGroupMap, err := c.admin.ListConsumerGroups()
	if err = c.handleError(err); err != nil {
		return nil, errors.Wrap(err, "list consumerGroups failed")
	}

	consumerGroups := make([]string, 0, len(consumerGroupMap))
	for consumerGroup := range consumerGroupMap {
		consumerGroups = append(consumerGroups, consumerGroup)
	}
	return consumerGroups, nil
}

func (c *client) ListConsumeTopics(ctx context.Context, consumerGroup string) ([]string, error) {
	if err := c.Open(); err != nil {
		return nil, errors.Wrap(err, "open client failed")
	}

	if c.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	if err := c.client.RefreshCoordinator(consumerGroup); err != nil {
		return nil, err
	}

	if c.admin == nil {
		return nil, fmt.Errorf("admin is nil")
	}
	resp, err := c.admin.ListConsumerGroupOffsets(consumerGroup, nil)
	if err = c.handleError(err); err != nil {
		return nil, errors.Wrap(err, "list consumerGroup offsets failed")
	}

	topics := make([]string, 0, len(resp.Blocks))
	for topic := range resp.Blocks {
		topics = append(topics, topic)
	}
	return topics, nil
}

// ListTopicsPartitionCounts returns map from topic name to partition counts.
func (c *client) ListTopicsPartitionCounts(ctx context.Context, topics []string) (map[string]int, error) {
	if err := c.Open(); err != nil {
		return nil, errors.Wrap(err, "open client failed")
	}
	topicsMetadata, err := c.admin.DescribeTopics(topics)
	if err = c.handleError(err); err != nil {
		return nil, errors.Wrap(err, "describe topics failed")
	}

	topicToPartitionMap := make(map[string]int, len(topicsMetadata))
	for _, metadata := range topicsMetadata {
		if metadata == nil {
			continue
		}

		if metadata.Err != sarama.ErrNoError {
			return nil, errors.New(metadata.Err.Error())
		}

		topicToPartitionMap[metadata.Name] = len(metadata.Partitions)
	}

	return topicToPartitionMap, nil
}

func (c *client) handleError(err error) error {
	if err == nil {
		return nil
	}

	netErr, ok := err.(net.Error)
	if ok {
		return c.handleErrorNet(netErr)
	}

	if strings.Contains(err.Error(), "EOF") {
		if err := c.Close(); err != nil {
			return errors.Wrap(err, "close client failed")
		}
	}

	return err
}

func (c *client) handleErrorNet(err net.Error) error {
	if strings.Contains(err.Error(), "write: broken pipe") {
		if err := c.Close(); err != nil {
			return errors.Wrap(err, "close client failed")
		}
	}
	return err
}

func setConfigDefaults(config kafka.Config) kafka.Config {
	if config.DialTimeout == nil {
		copyTime := defaultDialTimeout
		config.DialTimeout = &copyTime
	}

	if config.KeepAlive == nil {
		copyTime := defaultKeepAlive
		config.KeepAlive = &copyTime
	}

	if config.TLS == nil {
		tlsConfig := kafka.TLSConfig{
			Enabled: false,
		}
		config.TLS = &tlsConfig
	}

	return config
}
