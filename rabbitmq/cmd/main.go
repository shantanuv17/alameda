package main

import (
	"os"

	rabbit_app "prophetstor.com/alameda/rabbitmq/cmd/app"
)

func main() {
	if err := rabbit_app.PublishCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
