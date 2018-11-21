package methods

func (s *serverData) loadTestData() {
	// TODO: read test data from a file

	for _, item := range []struct {
		platformID    string
		deploymentIDs []string
	}{
		{
			platformID:    "platform-001",
			deploymentIDs: []string{"deployment-001"},
		},
	} {
		s.platforms[item.platformID] = newPlatform(item.deploymentIDs)
	}

}
