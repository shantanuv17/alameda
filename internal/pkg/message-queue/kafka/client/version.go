package client

import (
	"github.com/Shopify/sarama"
)

type SupportedVersion sarama.KafkaVersion

func GetSupportedVersion() []sarama.KafkaVersion {
	return sarama.SupportedVersions
}
