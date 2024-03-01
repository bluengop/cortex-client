package responses

import "time"

type Response struct {
	Success         bool            `json:"ok"`
	ErrorResponse   ErrorResponse   `json:"error,omitempty"`
	SuccessResponse SuccessResponse `json:"success,omitempty"`
}

type ErrorResponse struct {
	Details           string `json:"details"`
	GatewayHTTPStatus int    `json:"gatewayHttpStatus"`
	HTTPStatus        int    `json:"httpStatus"`
	Message           string `json:"message"`
	RequestID         string `json:"requestId"`
	Type              string `json:"type"`
}

type SuccessResponse struct {
	Definition    struct{}       `json:"definition"`
	Description   string         `json:"description"`
	Git           Git            `json:"git"`
	Groups        []string       `json:"groups"`
	Hierarchy     Hierarchy      `json:"hierarchy"`
	IsArchived    bool           `json:"isArchived"`
	LastUpdated   time.Time      `json:"lastUpdated"`
	Links         []Link         `json:"links"`
	Metadata      []MetadataItem `json:"metadata"`
	Name          string         `json:"name"`
	Ownership     OwnershipInfo  `json:"ownership"`
	SlackChannels []SlackChannel `json:"slackChannels"`
	Tag           string         `json:"tag"`
	Type          string         `json:"type"`
}
