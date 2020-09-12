package keycodes

import (
	"fmt"
	KeycodeMgt "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	Errors "github.com/containers-ai/alameda/internal/pkg/errors"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	Keycodes "github.com/containers-ai/api/datahub/keycodes"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (c *ServiceKeycodes) ActivateRegistrationData(ctx context.Context, in *Keycodes.ActivateRegistrationDataRequest) (*status.Status, error) {
	scope.Debug("Request received from ActivateRegistrationData grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(c.Config.InfluxDB)

	// Validate request
	if in.GetData() == "" {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: Errors.GetReason(Errors.ReasonMissingFieldReq, "Data"),
		}, nil
	}

	filePath := fmt.Sprintf("/tmp/%s.dat", AlamedaUtils.GenerateUUID())

	// Create empty registration file
	if AlamedaUtils.CreateFile(filePath) != nil {
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: "failed to create empty registration file",
		}, nil
	}

	// Write registration file
	if err := AlamedaUtils.WriteFile(filePath, []string{in.GetData()}); err != nil {
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: "failed to write registration file",
		}, nil
	}

	// Activation
	if err := keycodeMgt.PutSignatureDataFile(filePath); err != nil {
		AlamedaUtils.DeleteFile(filePath)
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	// Delete registration file
	if AlamedaUtils.DeleteFile(filePath) != nil {
		scope.Error("failed to delete registration file")
	}

	scope.Info("Successfully to activate keycode")

	if err := keycodeMgt.PostEvent(); err != nil {
		scope.Errorf("failed to post activate-keycode event: %s", err.Error())
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}
