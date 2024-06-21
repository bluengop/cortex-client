package client

import "encoding/json"

type Response struct {
	Success         bool            `json:"ok"`
	ErrorResponse   ErrorResponse   `json:"error"`
	SuccessResponse SuccessResponse `json:"success"`
}

type ErrorResponse struct {
	Details           string `json:"details"`
	GatewayHTTPStatus int32  `json:"gatewayHttpStatus"`
	HTTPStatus        int32  `json:"httpStatus"`
	Message           string `json:"message"`
	RequestID         string `json:"requestId"`
	Type              string `json:"type"`
}

type SuccessResponse struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Tag           string          `json:"tag"`
	Description   string          `json:"description"`
	Type          string          `json:"type"`
	Groups        []any           `json:"groups"`
	Owners        Owners          `json:"owners"`
	Ownership     Ownership       `json:"ownership"`
	OwnersV2      OwnersV2        `json:"ownersV2"`
	Metadata      []any           `json:"metadata"`
	Links         []any           `json:"links"`
	Definition    json.RawMessage `json:"definition"`
	Hierarchy     Hierarchy       `json:"hierarchy"`
	LastUpdated   string          `json:"lastUpdated"`
	IsArchived    bool            `json:"isArchived"`
	Git           any             `json:"git"`
	SlackChannels []any           `json:"slackChannels"`
}

type Owners struct {
	Groups        []any `json:"groups"`
	SlackChannels []any `json:"slackChannels"`
	Emails        []any `json:"emails"`
}

type Ownership struct {
	Groups        []any `json:"groups"`
	SlackChannels []any `json:"slackChannels"`
	Emails        []any `json:"emails"`
}

type OwnersV2 struct {
	Teams       []any `json:"teams"`
	Individuals []any `json:"individuals"`
}

type Hierarchy struct {
	Parents  []any `json:"parents"`
	Children []any `json:"children"`
}
