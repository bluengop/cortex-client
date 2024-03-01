package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/bluengop/cortex-client/internal/config"
	"github.com/bluengop/cortex-client/pkg/client/v1"
)

var headers map[string]string

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Hello world")

	// New Context
	ctx, cancel := context.WithCancel(context.Background())

	// Load Config
	conf, err := config.GetConfig("./config.env")
	if err != nil {
		fmt.Printf("The error is %s\n", err)
		logger.Error("Unable to load configuration")
	}

	// Create request
	request, err := client.NewRequest(
		&ctx,
		"/catalog",
		"get",
		conf.BaseUrl,
		headers,
	)
	if err != nil {
		logger.Error("Error when creating new HTTP Request: %s", err)
	}

	// Create client
	client := client.NewClient(
		conf.BaseUrl,
		conf.ApiKey,
	)

	// Send request and print response
	response, error := client.Send(&ctx, request)
	logger.Info(fmt.Sprintf("Response is %#v\n", response))
	logger.Info(fmt.Sprintf("Error is %s\n", error))
	logger.Debug(fmt.Sprintf("Cancel is of type %T\n", cancel))
}
