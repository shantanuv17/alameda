package keycodes

import (
	"context"
	"errors"
	"fmt"
	"time"

	Keycodes "github.com/containers-ai/api/datahub/keycodes"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
)

func ListKeycodes(keycode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to datahub
	conn, err := grpc.DialContext(ctx, *datahubAddress, grpc.WithBlock(), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(grpc_retry.WithMax(uint(3)))))
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	client := Keycodes.NewKeycodesServiceClient(conn)

	// Generate request
	keycodes := make([]string, 0)
	if keycode != "" {
		keycodes = append(keycodes, keycode)
	}
	in := &Keycodes.ListKeycodesRequest{
		Keycodes: keycodes,
	}

	// Do API request
	response, err := client.ListKeycodes(context.Background(), in)
	if err != nil {
		fmt.Println("[Error]: failed to connect to datahub")
		fmt.Println(fmt.Sprintf("[Reason]: %s", err.Error()))
		return err
	}

	// Check API result
	retCode := int32(response.GetStatus().GetCode())
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
