package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Cortex API Client
type Client struct {
	ApiKey     string
	HTTPClient *http.Client
}

// NewClient() returns a new HTTP Client for Cortex
// it requires a valid API Key (Bearer token).
func NewClient(apikey string) *Client {
	return &Client{
		ApiKey: apikey,
		HTTPClient: &http.Client{
			Timeout: time.Minute, // TODO: parametrize timeout
		},
	}
}

// Send HTTP Request Method
// given a context and a request pointers, sends
// the http request and returns a response object
func (c *Client) Send(ctx *context.Context, req *Request) (*Response, error) {
	// Add Bearer Token to the headers
	req.Payload.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	// Execute the request
	res, err := c.HTTPClient.Do(req.Payload)
	if err != nil {
		log.Println("Error when executing the http request: ", err)
		return nil, err
	}
	defer res.Body.Close()

	response := &Response{Success: false}

	// Check status code to select the proper data struct
	switch res.StatusCode {

	case http.StatusOK:
		response.Success = true
		if err := json.NewDecoder(res.Body).Decode(&response.SuccessResponse); err != nil {
			log.Printf("Failure when trying to decode the JSON response from API: %s", err)
			return nil, err
		}

	case http.StatusUnauthorized:
		log.Printf("Unauthorized request. Please, check your API token")
		return nil, errors.New("401: Unauthorized Request")

	case http.StatusNotFound:
		log.Printf("404: Not found")
		if err := json.NewDecoder(res.Body).Decode(&response.ErrorResponse); err != nil {
			log.Printf("Failure when trying to decode the JSON response from API: %s", err)
			return nil, err
		}

	default:
		log.Printf("The status code %d is not expected, so printing raw response and exiting", res.StatusCode)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("Unable to read response's body")
			return nil, err
		}
		log.Printf("Raw response:\n%s", body)
		return nil, nil
	}

	return response, nil
}
