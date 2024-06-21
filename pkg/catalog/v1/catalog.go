package catalog

type EKSClusterDefinition struct {
	Name                   string `json:"name"`
	Region                 string `json:"region"`
	AwsAccount             string `json:"aws_account"`
	K8SVersion             string `json:"k8s_version"`
	ContainerRuntime       string `json:"container_runtime"`
	KarpenterEnabled       bool   `json:"karpenter_enabled"`
	AwsEbsCsiDriverVersion string `json:"aws_ebs_csi_driver_version"`
}

type CoreServiceDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CoreServiceInClusterDefinition struct {
	Version     string `json:"version"`
	ClusterName string `json:"cluster_name"`
	CoreService string `json:"core_service"`
}
