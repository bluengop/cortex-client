package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	//"github.com/spf13/cobra"

	"github.com/bluengop/cortex-client/internal/config"
	"github.com/bluengop/cortex-client/pkg/client/v1"
)

var headers map[string]string

func main() {
	// TODO: logger should me moved apart to a different packagece
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Debug("Still TODO. By now using 'log' default library")

	// New Context
	ctx, cancel := context.WithCancel(context.Background())

	// Load Config
	configFile := "./config.env"
	log.Printf("Loading configuration from %s", configFile)
	conf, err := config.GetConfig(configFile)
	if err != nil {
		panic("unable to load configuration")
	}

	// Create request
	request, err := client.NewRequest(
		&ctx,
		"GET",
		conf.BaseUrl,
		"/catalog/us-east-1-stage-sandbox-bravo",
		headers,
	)
	if err != nil {
		logger.Error("Error when creating new HTTP Request: %s", err)
		panic("unable to create request")
	}

	// Create client
	client := client.NewClient(
		conf.ApiKey,
	)

	// Send request and print response
	log.Printf("Sending HTTP request...")
	log.Printf("Calling URL %s", request.Payload.URL)
	response, err := client.Send(&ctx, request)
	if err != nil {
		panic("an error occurred with the HTTP response")
	}

	if response.Success {
		log.Println("Id:", response.SuccessResponse.ID)
		log.Println("Type:", response.SuccessResponse.Type)
		log.Println("Name:", response.SuccessResponse.Name)
		log.Println("Tag:", response.SuccessResponse.Tag)
		log.Println("EKS Name:", response.SuccessResponse.Definition.Name)
		log.Println("EKS Region:", response.SuccessResponse.Definition.Region)
		log.Println("EKS container runtime:", response.SuccessResponse.Definition.ContainerRuntime)
		log.Println("EkS K8s version:", response.SuccessResponse.Definition.K8SVersion)
	} else {
		log.Println("An error occurred")
		log.Println("Message:", response.ErrorResponse.Message)
		log.Println("Details:", response.ErrorResponse.Details)
	}

	logger.Debug(fmt.Sprintf("cancel is of tipe %T", cancel))
}
