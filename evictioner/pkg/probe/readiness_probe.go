package probe

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	datahub_client "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
)

type ReadinessProbeConfig struct {
	DatahubAddr         string
	AdmissionController *AdmissionController
}

type AdmissionController struct {
	SvcName   string
	Namespace string
	Port      int32
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

	datahubServiceClnt := datahub_client.NewDatahubServiceClient(conn)
	_, err = datahubServiceClnt.ListNodes(context.Background(), &datahub_resources.ListNodesRequest{})
	if err != nil {
		return err
	}

	// if len(res.GetNodes()) == 0 {
	// 	return fmt.Errorf("No nodes found in datahub")
	// }

	return err
}

func queryAdmissionControlerSvc(admCtlSvcName string, admCtlSvcNS string, admCtlPort int32) error {
	svcURL := fmt.Sprintf("https://%s.%s:%s", admCtlSvcName, admCtlSvcNS, fmt.Sprint(admCtlPort))
	curlCmd := exec.Command("curl", "-k", svcURL)

	_, err := curlCmd.CombinedOutput()
	if err != nil {
		return err
	}

	return err
}
