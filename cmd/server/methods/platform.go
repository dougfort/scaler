package methods

type deploymentData struct {
	instanceCount uint32
	// TODO: resources
}

type platformData struct {
	deployments map[string]deploymentData
}

func newPlatform(deploymentIDs []string) *platformData {
	var p platformData

	p.deployments = make(map[string]deploymentData)
	for _, d := range deploymentIDs {
		p.deployments[d] = deploymentData{}
	}

	return &p
}

func (p *platformData) enumerateDeployments() []string {
	var deploymentIDs []string

	for deploymentID := range p.deployments {
		deploymentIDs = append(deploymentIDs, deploymentID)
	}

	return deploymentIDs
}

// getDeploymentData returns false if the deployment does not exist on
// this platform. That is different from having zero instances.
func (p *platformData) getDeploymentData(
	deploymentID string,
) (deploymentData, bool, error) {
	data, ok := p.deployments[deploymentID]
	if !ok {
		return data, false, nil
	}
	return data, true, nil
}
