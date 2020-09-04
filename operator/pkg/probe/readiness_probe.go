package probe

import (
	"os/exec"

	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
)

type ReadinessProbeConfig struct {
	WHSrvPort   int32
	DatahubAddr string
}

func queryDatahub(datahubAddr string) error {
	datahubClient := datahubpkg.NewClient(datahubAddr)
	_, err := datahubClient.ListNodes(&datahub_resources.ListNodesRequest{})
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
