package methods

// DeploymentStatus is the status of one deployment on a single platform
type DeploymentStatus struct {
	PlatformID    string
	InstanceCount int32
}
