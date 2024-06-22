package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/bluengop/cortex-client/internal/config"
	"github.com/bluengop/cortex-client/pkg/catalog/v1"
	"github.com/bluengop/cortex-client/pkg/client/v1"
)

var headers map[string]string

var (
	entityType    string
	listEntities  bool
	detailsEntity string
	applyFile     string
)

var rootCmd = &cobra.Command{
	Use:   "Cortex Client",
	Short: "A golang API REST client for Cortex to get and create/update catalog entities",
}

var listEntitiesCmd = &cobra.Command{
	Use:   "list-entities",
	Short: "List entities",
	Run: func(cmd *cobra.Command, args []string) {
		entityType, _ = cmd.Flags().GetString("type")
		listEntities = true
		log.Println("Listing entities")
		if entityType != "" {
			log.Printf("Filter by type: %s\n", entityType)
		}
	},
}

var detailsCmd = &cobra.Command{
	Use:   "details <entity>",
	Short: "Retrieve details of the entity",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		detailsEntity = args[0]
		log.Printf("Retrieving details for entity: %s\n", detailsEntity)
	},
}

var applyCmd = &cobra.Command{
	Use:   "apply <file>",
	Short: "Create or update an entity defined in a YAML file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		applyFile = args[0]
		log.Printf("Applying configuration from file: %s\n", applyFile)
	},
}

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

		// Unmarshall definition:
		var eksDef catalog.EKSClusterDefinition
		if err := json.Unmarshal(response.SuccessResponse.Definition, &eksDef); err == nil {
			log.Println("EKSClusterDefinition:", eksDef)
			log.Println("Cluster Name:", eksDef.Name)
		}
	} else {
		log.Println("An error occurred")
		log.Println("Message:", response.ErrorResponse.Message)
		log.Println("Details:", response.ErrorResponse.Details)
	}
	logger.Debug(response.SuccessResponse.Description)

	logger.Debug(fmt.Sprintf("cancel is of tipe %T", cancel))
}
