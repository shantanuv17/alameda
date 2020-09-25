package keycodes

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"prophetstor.com/api/datahub"
	"prophetstor.com/api/datahub/keycodes"
)

func ListKeycodes(keycode string) error {
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
	keys := make([]string, 0)
	if keycode != "" {
		keys = append(keys, keycode)
	}
	in := &keycodes.ListKeycodesRequest{
		Keycodes: keys,
	}

	// Do API request
	response, err := client.ListKeycodes(context.Background(), in)
	if err != nil {
		fmt.Println("[Error]: failed to connect to datahub")
		fmt.Println(fmt.Sprintf("[Reason]: %s", err.Error()))
		return err
	}

	// Check API result
	retCode := response.GetStatus().GetCode()
	if retCode == int32(code.Code_OK) {
		fmt.Println(fmt.Sprintf("[Result]: %s", code.Code_name[retCode]))

		fmt.Println("[Keycode]")
		if len(response.GetKeycodes()) > 0 {
			PrintKeycode(response.GetKeycodes()[0])
		}

		fmt.Println("")

		fmt.Println("[Summary]")
		if response.GetSummary() != nil {
			PrintKeycode(response.GetSummary())
		}
	} else {
		fmt.Println(fmt.Sprintf("[Result]: %s", code.Code_name[retCode]))
		fmt.Println(fmt.Sprintf("[Reason]: %s", response.GetStatus().GetMessage()))
		return errors.New(response.GetStatus().GetMessage())
	}

	return nil
}
