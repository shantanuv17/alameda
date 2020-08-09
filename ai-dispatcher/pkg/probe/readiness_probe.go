package probe

import (
	"context"
	"time"

	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type ReadinessProbeConfig struct {
	DatahubAddr string
	QueueURL    string
}

func queryDatahub(datahubAddr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, datahubAddr, grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(uint(3)))))
	if conn != nil {
		defer conn.Close()
	}

	if err != nil {
		return err
	}

	datahubServiceClnt := datahub_v1alpha1.NewDatahubServiceClient(conn)
	_, err = datahubServiceClnt.ListNodes(context.Background(), &datahub_resources.ListNodesRequest{})
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
