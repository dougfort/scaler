package methods

type deploymentData struct {
	instanceCount uint32
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
