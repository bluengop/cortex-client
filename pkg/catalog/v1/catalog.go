package catalog

import "time"

// Structs and Data Types for Catalog Entity Descriptors
// https://docs.cortex.io/docs/api/get-entity-descriptor
type EntityDescriptor struct {
	Info    Info     `json:"info"`
	Openapi string   `json:"openapi"`
	Servers []Server `json:"servers"`
}

type Info struct {
	Title             string         `json:"title"`
	XCortexTag        string         `json:"x-cortex-tag"`
	XCortexType       string         `json:"x-cortex-type"`
	XCortexDefinition map[string]any `json:"x-cortex-definition"`
}

type Server struct {
	URL string `json:"url"`
}

// Structs and Data Types for Catalog Entity Details
// https://docs.cortex.io/docs/api/get-entity-details
type EntityDetails struct {
	Name          string         `json:"name"`
	Tag           string         `json:"tag"`
	Description   any            `json:"description"`
	Type          string         `json:"type"`
	Groups        []string       `json:"groups"`
	Owners        []Owners       `json:"owners"`
	Ownership     []Owners       `json:"ownership"`
	Metadata      []string       `json:"metadata"`
	Links         []string       `json:"links"`
	Definition    map[string]any `json:"definition"`
	Hierarchy     Hierarchy      `json:"hierarchy"`
	LastUpdated   time.Time      `json:"lastUpdated"`
	IsArchived    bool           `json:"isArchived"`
	Git           any            `json:"git"`
	SlackChannels []string       `json:"slackChannels"`
}

type Owners struct {
	Groups        []string `json:"groups"`
	SlackChannels []string `json:"slackChannels"`
	Emails        []string `json:"emails"`
}

type Hierarchy struct {
	Parents  []string `json:"parents"`
	Children []string `json:"children"`
}

/*
type EKSCluster struct {
	Name        string               `json:"name"`
	Tag         string               `json:"tag"`
	Description any                  `json:"description"`
	Type        string               `json:"type"`
	LastUpdated string               `json:"lastUpdated"`
	Definition  EKSClusterDefinition `json:"definition"`
}

type EKSClusterDefinition struct {
	Region                 string `json:"region"`
	AwsAccount             string `json:"aws_account"`
	K8SVersion             string `json:"k8s_version"`
	ContainerRuntime       string `json:"container_runtime"`
	KarpenterEnabled       bool   `json:"karpenter_enabled"`
	AwsEbsCsiDriverVersion string `json:"aws_ebs_csi_driver_version"`
}
*/
