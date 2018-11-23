package methods

func (s *serverData) loadTestData() {
	// TODO: read test data from a file

	for _, item := range []struct {
		platformID    string
		deploymentIDs []string
	}{
		{
			platformID: "platform-001",
			deploymentIDs: []string{
				"deployment-001",
				"deployment-002",
				"deployment-003",
			},
		},
		{
			platformID: "platform-002",
			deploymentIDs: []string{
				"deployment-001",
				"deployment-002",
				"deployment-004",
			},
		},
		{
			platformID: "platform-003",
			deploymentIDs: []string{
				"deployment-001",
				"deployment-003",
				"deployment-004",
			},
		},
		{
			platformID: "platform-004",
			deploymentIDs: []string{
				"deployment-002",
				"deployment-003",
				"deployment-004",
			},
		},
	} {
		s.platforms[item.platformID] = newPlatform(item.deploymentIDs)
	}

}
