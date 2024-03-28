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
		log.Println("Response:", *response)
		log.Println("Definition:", response.SuccessResponse.Definition)
		log.Println("Description:", response.SuccessResponse.Description)
		log.Println("Git:", response.SuccessResponse.Git)
		log.Println("Groups:", response.SuccessResponse.Groups)
		log.Println("Hierarchy:", response.SuccessResponse.Hierarchy)
		log.Println("IsArchived:", response.SuccessResponse.IsArchived)
		log.Println("LastUpdated:", response.SuccessResponse.LastUpdated)
		log.Println("Links:", response.SuccessResponse.Links)
		log.Println("Metadata:", response.SuccessResponse.Metadata)
		log.Println("Name:", response.SuccessResponse.Name)
		log.Println("OwnersV2:", response.SuccessResponse.OwnersV2)
		log.Println("Ownership:", response.SuccessResponse.Ownership)
		log.Println("SlackChannels:", response.SuccessResponse.SlackChannels)
		log.Println("Tag:", response.SuccessResponse.Tag)
		log.Println("Type:", response.SuccessResponse.Type)
	} else {
		log.Printf(response.ErrorResponse.Message)
		log.Printf(response.ErrorResponse.Details)
	}

	logger.Debug(fmt.Sprintf("cancel is of tipe %T", cancel))
}
