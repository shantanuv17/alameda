package probe

import (
	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/streadway/amqp"
)

type ReadinessProbeConfig struct {
	DatahubAddr string
	QueueURL    string
}

func queryDatahub(datahubAddr string) error {
	datahubServiceClnt := datahubpkg.NewClient(datahubAddr)
	_, err := datahubServiceClnt.ListNodes(&datahub_resources.ListNodesRequest{})
	if err != nil {
		return err
	}

	// if len(res.GetNodes()) == 0 {
	// 	return fmt.Errorf("No nodes found in datahub")
	// }

	return err
}

func connQueue(url string) error {
	conn, err := amqp.Dial(url)
	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return err
	}
	return nil
}
