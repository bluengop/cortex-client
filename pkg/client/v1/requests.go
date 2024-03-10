package client

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strings"
)

// Allowed HTTP methods
var methods = []string{
	"GET",
	"POST",
}

// Default headers for the requests
var defaultHeaders = map[string]string{
	"Content-Type": "application/json; charset=utf-8",
	"Accept":       "application/json; charset=utf-8",
}

// Request struct
type Request struct {
	Payload *http.Request
}

// HTTP Request parameters
// So far I've only added the entity tag (x-cortex-tag)
// that identifies the entity.
type Parameters struct {
	Tag string `json:"tag,omitempty"`
}

// Returns new Request instance
// The parameters accepted by the function are:
// 1. A pointer to a context.
// 2. The HTTP method (case insensitive).
// 3. The Base URL for the API call.
// 4. The path inside the above Base URL (resource route).
// 5. (Optionally) custom headers for the request.
func NewRequest(ctx *context.Context, method, baseurl, path string, headers ...map[string]string) (*Request, error) {
	// Check if provided verb is allowed
	method = strings.ToTitle(method)
	if !slices.Contains(methods, method) {
		log.Printf("HTTP method %s is not allowed.\n", method)
		return nil, errors.New("bad http method")
	}

	// Form full URL
	url, err := url.JoinPath(baseurl, path)
	if err != nil {
		log.Printf("Unable to form the full URL with provided parameters:\n")
		log.Printf("Base URL: %s\nPath: %s", baseurl, path)
		return nil, err
	}

	// Create new HTTP Request
	payload := strings.NewReader(``)
	req, err := http.NewRequestWithContext(*ctx, method, url, payload)
	if err != nil {
		log.Printf("Error when creating a new http request: %s\n", err)
		return nil, err
	}

	// Adding default headers
	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}

	// Adding specific headers, if so
	if len(headers) > 0 {
		for _, headerMap := range headers {
			for key, value := range headerMap {
				req.Header.Set(key, value)
			}
		}
	}

	request := &Request{
		Payload: req,
	}

	return request, nil
}
