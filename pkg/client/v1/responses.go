package client

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

type Git struct {
	Alias         string `json:"alias"`
	Basepath      string `json:"basepath"`
	Provider      string `json:"provider"`
	Repository    string `json:"repository"`
	RepositoryURL string `json:"repositoryUrl"`
}

type Hierarchy struct {
	Children []HierarchyNode `json:"children"`
	Parents  []ParentNode    `json:"parents"`
}
type HierarchyNode struct {
	Children    []HierarchyNode `json:"children"`
	Definition  struct{}        `json:"definition"`
	Description string          `json:"description"`
	Groups      []string        `json:"groups"`
	Name        string          `json:"name"`
	Tag         string          `json:"tag"`
	Type        string          `json:"type"`
}

type ParentNode struct {
	Definition  struct{}   `json:"definition"`
	Description string     `json:"description"`
	Groups      []string   `json:"groups"`
	Name        string     `json:"name"`
	Parents     ParentInfo `json:"parents"`
	Tag         string     `json:"tag"`
	Type        string     `json:"type"`
}

type ParentInfo struct {
	Definition  struct{} `json:"definition"`
	Description string   `json:"description"`
	Groups      []string `json:"groups"`
	Name        string   `json:"name"`
	Tag         string   `json:"tag"`
	Type        string   `json:"type"`
}

type Link struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type MetadataItem struct {
	Key   string   `json:"key"`
	Value struct{} `json:"value"`
}

type OwnershipInfo struct {
	Emails        []EmailInfo    `json:"emails"`
	Groups        []GroupInfo    `json:"groups"`
	SlackChannels []SlackChannel `json:"slackChannels"`
}

type EmailInfo struct {
	Description string `json:"description"`
	Email       string `json:"email"`
	Inheritance string `json:"inheritance"`
}

type GroupInfo struct {
	Description string `json:"description"`
	GroupName   string `json:"groupName"`
	Inheritance string `json:"inheritance"`
	Provider    string `json:"provider"`
}

type SlackChannel struct {
	Channel              string `json:"channel"`
	Description          string `json:"description"`
	Inheritance          string `json:"inheritance"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
}
