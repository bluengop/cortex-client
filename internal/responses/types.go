package responses

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
