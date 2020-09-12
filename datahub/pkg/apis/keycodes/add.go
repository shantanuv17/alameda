package keycodes

import (
	KeycodeMgt "github.com/containers-ai/alameda/datahub/pkg/account-mgt/keycodes"
	Errors "github.com/containers-ai/alameda/internal/pkg/errors"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	Keycodes "github.com/containers-ai/api/datahub/keycodes"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (c *ServiceKeycodes) AddKeycode(ctx context.Context, in *Keycodes.AddKeycodeRequest) (*Keycodes.AddKeycodeResponse, error) {
	scope.Debug("Request received from AddKeycode grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt(c.Config.InfluxDB)

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
		Keycode: TransformKeycode(keycode),
	}, nil
}
