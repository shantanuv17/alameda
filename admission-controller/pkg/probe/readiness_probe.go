package probe

import (
	"fmt"
	"os/exec"

	datahubpkg "github.com/containers-ai/alameda/pkg/datahub"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
)

type ReadinessProbeConfig struct {
	DatahubAddr   string
	AdmCtrSrvPort int32
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

func queryWebhookSrv(port int32) error {

	svcURL := fmt.Sprintf("https://localhost:%s", fmt.Sprint(port))
	curlCmd := exec.Command("curl", "-k", svcURL)

	_, err := curlCmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
