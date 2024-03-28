package client

type Definition struct {
	Definition any `json:"definition"`
}
type Children struct {
	Children    *Children  `json:"children"`
	Definition  Definition `json:"definition"`
	Description string     `json:"description"`
	Groups      []string   `json:"groups"`
	Name        string     `json:"name"`
	Tag         string     `json:"tag"`
	Type        string     `json:"type"`
}

type Email struct {
	Description string `json:"description"`
	Email       string `json:"email"`
	Inheritance string `json:"inheritance"`
}

type ErrorResponse struct {
	Details           string `json:"details"`
	GatewayHTTPStatus int32  `json:"gatewayHttpStatus"`
	HTTPStatus        int32  `json:"httpStatus"`
	Message           string `json:"message"`
	RequestID         string `json:"requestId"`
	Type              string `json:"type"`
}

type Git struct {
	Alias         string `json:"alias"`
	Basepath      string `json:"basepath"`
	Provider      string `json:"provider"`
	Repository    string `json:"repository"`
	RepositoryURL string `json:"repositoryUrl"`
}

type Group struct {
	Description string `json:"description"`
	GroupName   string `json:"groupName"`
	Inheritance string `json:"inheritance"`
	Provider    string `json:"provider"`
}

type Hierarchy struct {
	Children []Children `json:"children"`
	Parents  []Parents  `json:"parents"`
}

type Individual struct {
	Description string `json:"description"`
	Email       string `json:"email"`
}

type Link struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OwnersV2 struct {
	Individuals []Individual `json:"individuals"`
	Teams       []Team       `json:"teams"`
}

type Ownership struct {
	Emails []Email `json:"emails"`
	Groups []Group `json:"groups"`
}

type Parents struct {
	Definition  Definition `json:"definition"`
	Description string     `json:"description"`
	Groups      []string   `json:"groups"`
	Name        string     `json:"name"`
	Parents     *Parents   `json:"parents"`
	Tag         string     `json:"tag"`
	Type        string     `json:"type"`
}

type Response struct {
	Success         bool            `json:"ok"`
	ErrorResponse   ErrorResponse   `json:"error,omitempty"`
	SuccessResponse SuccessResponse `json:"success,omitempty"`
}

type SlackChannel struct {
	Description          string `json:"description"`
	Name                 string `json:"name"`
	NotificationsEnabled bool   `json:"notificationsEnabled"`
}

type SuccessResponse struct {
	Definition    Definition     `json:"definition"`
	Description   string         `json:"description"`
	Git           Git            `json:"git"`
	Groups        []string       `json:"groups"`
	Hierarchy     Hierarchy      `json:"hierarchy"`
	IsArchived    bool           `json:"isArchived"`
	LastUpdated   string         `json:"lastUpdated"` //TODO: use time.Time, RFC3339 ?
	Links         []Link         `json:"links"`
	Metadata      []Metadata     `json:"metadata"`
	Name          string         `json:"name"`
	OwnersV2      OwnersV2       `json:"ownersV2"`
	Ownership     Ownership      `json:"ownership"`
	SlackChannels []SlackChannel `json:"slackChannels"`
	Tag           string         `json:"tag"`
	Type          string         `json:"type"`
}

type Team struct {
	Description string `json:"description"`
	IsArchived  bool   `json:"isArchived"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
}
