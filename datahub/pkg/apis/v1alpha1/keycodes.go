package v1alpha1

import (
	"fmt"
	KeycodeMgt "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	FormatKeycode "github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/keycodes"
	Errors "github.com/containers-ai/alameda/internal/pkg/errors"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	Keycodes "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/keycodes"
	"github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type IError = Errors.InternalError

func (s *ServiceV1alpha1) AddKeycode(ctx context.Context, in *Keycodes.AddKeycodeRequest) (*Keycodes.AddKeycodeResponse, error) {
	scope.Debug("Request received from AddKeycode grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(s.Config.InfluxDB)

	// Validate request
	if in.GetKeycode() == "" {
		return &Keycodes.AddKeycodeResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INVALID_ARGUMENT),
				Message: Errors.GetReason(Errors.ReasonMissingFieldReq, "Keycode"),
			},
		}, nil
	}

	// Add keycode
	if err := keycodeMgt.AddKeycode(in.GetKeycode()); err != nil {
		scope.Error(err.Error())
		return &Keycodes.AddKeycodeResponse{
			Status: &status.Status{
				Code:    CategorizeKeycodeErrorId(err.(*IError).ErrorID),
				Message: err.Error(),
			},
		}, nil
	}

	scope.Infof("successfully to add keycode(%s)", in.GetKeycode())

	if err := keycodeMgt.PostEvent(); err != nil {
		scope.Errorf("failed to post add-keycode event: %s", err.Error())
	}

	keycode, _ := keycodeMgt.GetKeycode(in.GetKeycode())
	return &Keycodes.AddKeycodeResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Keycode: FormatKeycode.NewKeycode(keycode),
	}, nil
}

func (s *ServiceV1alpha1) ListKeycodes(ctx context.Context, in *Keycodes.ListKeycodesRequest) (*Keycodes.ListKeycodesResponse, error) {
	scope.Debug("Request received from ListKeycodes grpc function: " + AlamedaUtils.InterfaceToString(in))

	var (
		err      error
		keycodes []*KeycodeMgt.Keycode
		summary  *KeycodeMgt.Keycode
	)

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(s.Config.InfluxDB)

	if len(in.GetKeycodes()) == 0 {
		// Read all keycodes
		keycodes, summary, err = keycodeMgt.GetAllKeycodes()
	} else {
		// Read keycodes
		keycodes, summary, err = keycodeMgt.GetKeycodes(in.GetKeycodes())
	}

	if err != nil {
		scope.Error(err.Error())
		response := &Keycodes.ListKeycodesResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}
		return response, nil
	}

	// Prepare response
	response := &Keycodes.ListKeycodesResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Keycodes: FormatKeycode.NewKeycodeList(keycodes),
		Summary:  FormatKeycode.NewKeycode(summary),
	}

	return response, nil
}

func (s *ServiceV1alpha1) DeleteKeycode(ctx context.Context, in *Keycodes.DeleteKeycodeRequest) (*status.Status, error) {
	scope.Debug("Request received from DeleteKeycode grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(s.Config.InfluxDB)

	// Validate request
	if in.GetKeycode() == "" {
		return &status.Status{
			Code:    int32(code.Code_INVALID_ARGUMENT),
			Message: Errors.GetReason(Errors.ReasonMissingFieldReq, "Keycode"),
		}, nil
	}

	// Delete keycode
	if err := keycodeMgt.DeleteKeycode(in.GetKeycode()); err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    CategorizeKeycodeErrorId(err.(*IError).ErrorID),
			Message: err.Error(),
		}, nil
	}

	scope.Infof("successfully to delete keycode(%s)", in.GetKeycode())

	if err := keycodeMgt.PostEvent(); err != nil {
		scope.Errorf("failed to post delete-keycode event: %s", err.Error())
	}

	return &status.Status{Code: int32(code.Code_OK)}, nil
}

func (s *ServiceV1alpha1) GenerateRegistrationData(ctx context.Context, in *empty.Empty) (*Keycodes.GenerateRegistrationDataResponse, error) {
	scope.Debug("Request received from GenerateRegistrationData grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(s.Config.InfluxDB)

	// Generate registration data
	registrationData, err := keycodeMgt.GetRegistrationData()
	if err != nil {
		scope.Error(err.Error())
		return &Keycodes.GenerateRegistrationDataResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	return &Keycodes.GenerateRegistrationDataResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Data: registrationData,
	}, nil
}

func (s *ServiceV1alpha1) ActivateRegistrationData(ctx context.Context, in *Keycodes.ActivateRegistrationDataRequest) (*status.Status, error) {
	scope.Debug("Request received from ActivateRegistrationData grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(s.Config.InfluxDB)

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

func CategorizeKeycodeErrorId(errorId int) int32 {
	switch errorId {
	case Errors.ReasonKeycodeInvalidKeycode:
		return int32(code.Code_INVALID_ARGUMENT)
	}
	return int32(code.Code_INTERNAL)
}
