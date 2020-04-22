package schema

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
)

func GetKafkaTopicSchema() (schemas.Schema, error) {
	schema := schemas.Schema{}
	data, err := Asset("kafka-topic.json")
	if err != nil {
		return schema, errors.Wrap(err, "read bindata failed")
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		return schema, err
	}
	return schema, nil
}

func GetKafkaConsumerGroupSchema() (schemas.Schema, error) {
	schema := schemas.Schema{}
	data, err := Asset("kafka-consumergroup.json")
	if err != nil {
		return schema, errors.Wrap(err, "read bindata failed")
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		return schema, err
	}
	return schema, nil
}

func GetNginxSchema() (schemas.Schema, error) {
	schema := schemas.Schema{}
	data, err := Asset("nginx.json")
	if err != nil {
		return schema, errors.Wrap(err, "read bindata failed")
	}
	if err := json.Unmarshal(data, &schema); err != nil {
		return schema, err
	}
	return schema, nil
}
