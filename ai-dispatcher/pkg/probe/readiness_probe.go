package probe

import (
	"fmt"

	"github.com/streadway/amqp"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	datahub_resources "prophetstor.com/api/datahub/resources"
)

type ReadinessProbeConfig struct {
	DatahubAddr string
	QueueURL    string
}

func queryDatahub(datahubAddr string) error {
	datahubServiceClnt := datahubpkg.NewClient(datahubAddr)
	if datahubServiceClnt == nil {
		return fmt.Errorf("failed to new datahub client")
	}
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
