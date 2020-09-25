package probe

import (
	"fmt"
	"os/exec"

	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	datahub_resources "prophetstor.com/api/datahub/resources"
)

type ReadinessProbeConfig struct {
	WHSrvPort   int32
	DatahubAddr string
}

func queryDatahub(datahubAddr string) error {
	datahubClient := datahubpkg.NewClient(datahubAddr)
	if datahubClient == nil {
		return fmt.Errorf("New datahub client failed")
	}
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
