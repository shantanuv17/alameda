package probe

import (
	"context"
	"os/exec"
	"time"

	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
)

type ReadinessProbeConfig struct {
	WHSrvPort   int32
	DatahubAddr string
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

	return err
}

func queryWebhookSrv(svcURL string) error {
	curlCmd := exec.Command("curl", "-k", svcURL)

	_, err := curlCmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
