package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bluengop/cortex-client/internal/requests"
	"github.com/bluengop/cortex-client/internal/responses"
)

const BaseURLv1 = "https://api.getcortexapp.com/api/v1"

// Cortex API Client
type Client struct {
	BaseURL    string
	ApiKey     string
	HTTPClient *http.Client
}

// NewClient() returns a new HTTP Client for Cortex
func NewClient(url, apikey string) *Client {
	if url == "" {
		url = BaseURLv1
	}

	return &Client{
		BaseURL: url,
		ApiKey:  apikey,
		HTTPClient: &http.Client{
			Timeout: time.Minute, // TODO: parametrize timeout
		},
	}
}

// HTTP Request parameters
// So far I've only added the entity tag (x-cortex-tag)
// that identifies the entity.
type Parameters struct {
	Tag string `json:"tag,omitempty"`
}

// Send HTTP Request Method
func (c *Client) SendRequest(ctx *context.Context, req *requests.Request) (*responses.Response, error) {
	// Add Bearer Token to the headers
	req.Load.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	// Execute the request
	res, err := c.HTTPClient.Do(req.Load)
	if err != nil {
		log.Println("Error when executing the http request: ", err)
		return nil, err
	}
	defer res.Body.Close()

	// Instance new response and fail (error) objects
	response := &responses.Response{}
	var fail error = nil

	// Check status code to select the proper data struct
	switch res.StatusCode {
	case 200:
		response.Success = true
		fmt.Println(res.Body)
		if err := json.NewDecoder(res.Body).Decode(&response.SuccessResponse); err != nil {
			log.Println("Failure when trying to decode the JSON response from API")
			return nil, err
		}
	case 401:
		response.Success = false
		fail = fmt.Errorf("unauthorized request")
		log.Printf("Unauthorized request. Please, check your API token")
		return nil, fail
	case 404:
		response.Success = false
		body, err := io.ReadAll(res.Body)
		if err == nil {
			log.Printf("Response Body: %s\n", body)
		}
		if err := json.Unmarshal(body, &response.ErrorResponse); err != nil {
			log.Printf("Failure when trying to decode the JSON response from API: %s\n", err)
			log.Printf("Printing raw response body for debugging:\n%s\n", body)
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return response, fail
}
