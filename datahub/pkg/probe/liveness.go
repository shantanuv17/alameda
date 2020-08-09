package probe

import (
	"context"
	"fmt"
	"time"

	DatahubV1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
)

type LivenessProbeConfig struct {
	BindAddr string
}

func queryDatahub(bindAddr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost%s", bindAddr), grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(uint(3)))))

	if conn != nil {
		defer conn.Close()
	}
	if err != nil {
		return err
	}

	client := DatahubV1alpha1.NewDatahubServiceClient(conn)

	status, err := client.Ping(context.Background(), &empty.Empty{})
	if err != nil {
		return errors.Wrap(err, "failed to connect to datahub")
	}

	if status.GetCode() != int32(code.Code_OK) {
		return errors.Wrap(err, status.GetMessage())
	}

	return nil
}
