package requests

import (
	"context"
	"errors"
	"log"
	"net/http"
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
	Load *http.Request
}

// HTTP Request parameters
// So far I've only added the entity tag (x-cortex-tag)
// that identifies the entity.
type Parameters struct {
	Tag string `json:"tag,omitempty"`
}

// Returns new Request instance
func NewRequest(ctx *context.Context, path, method, url string, headers map[string]string) (*Request, error) {
	// Check if provided verb is allowed
	method = strings.ToTitle(method)
	if !slices.Contains(methods, method) {
		log.Printf("HTTP method %s is not allowed.\n", method)
		return nil, errors.New("bad http method")
	}

	req, err := http.NewRequestWithContext(*ctx, method, url, nil) // Body is io.Reader
	if err != nil {
		log.Println("Error when creating a new http request: ", err)
		return nil, err
	}

	// Adding default headers
	for key, value := range defaultHeaders {
		req.Header.Set(key, value)
	}

	// Adding specific headers, if so
	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	request := &Request{
		Load: req,
	}

	return request, nil
}
