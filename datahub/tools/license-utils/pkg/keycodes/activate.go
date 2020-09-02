package keycodes

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/pkg/utils"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/keycodes"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
)

func Activate(filePath string) error {
	// Check if registration file is found
	if !utils.FileExists(filePath) {
		reason := fmt.Sprintf("registration file(%s) is not found", filePath)
		fmt.Println(fmt.Sprintf("[Error]: %s", reason))
		return errors.New(reason)
	}

	// Read registration file
	registrationFile, err := utils.ReadFile(filePath)
	if err != nil {
		reason := fmt.Sprintf("failed to read registration file(%s)", filePath)
		fmt.Println(fmt.Sprintf("[Error]: %s", reason))
		return errors.New(reason)
	}

	// Check if registration file is empty
	if len(registrationFile) == 0 {
		reason := fmt.Sprintf("registration file(%s) is empty", filePath)
		fmt.Println(fmt.Sprintf("[Error]: %s", reason))
		return errors.New(reason)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Connect to datahub
	conn, err := grpc.DialContext(ctx, *datahubAddress, grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(retry.WithMax(uint(3)))))
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := datahub.NewDatahubServiceClient(conn)

	// Generate request
	in := &keycodes.ActivateRegistrationDataRequest{
		Data: registrationFile[0],
	}

	// Do API request
	stat, err := client.ActivateRegistrationData(context.Background(), in)
	if err != nil {
		fmt.Println("[Error]: failed to connect to datahub")
		fmt.Println(fmt.Sprintf("[Reason]: %s", err.Error()))
		return err
	}

	// Check API result
	retCode := stat.GetCode()
	if retCode == int32(code.Code_OK) {
		fmt.Println(fmt.Sprintf("[Result]: %s", code.Code_name[retCode]))
	} else {
		fmt.Println(fmt.Sprintf("[Result]: %s", code.Code_name[retCode]))
		fmt.Println(fmt.Sprintf("[Reason]: %s", stat.GetMessage()))
		return errors.New(stat.GetMessage())
	}

	return nil
}
