package methods

func (s *serverData) loadTestData() {
	// TODO: read test data from a file

	platformIDSet := make(map[string]struct{})

	for _, item := range []struct {
		deploymentID     string
		deploymentStatus DeploymentStatus
	}{
		{
			deploymentID: "deployment-001",
			deploymentStatus: DeploymentStatus{
				PlatformID:    "platform-001",
				InstanceCount: 0,
			},
		},
	} {
		s.DeploymentStatusMap[item.deploymentID] =
			append(s.DeploymentStatusMap[item.deploymentID], item.deploymentStatus)
		platformIDSet[item.deploymentStatus.PlatformID] = struct{}{}
	}

	for platformID := range platformIDSet {
		s.PlatformIDs = append(s.PlatformIDs, platformID)
	}
}
