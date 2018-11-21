package methods

import (
	"golang.org/x/net/context"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
)

// EnumerateDeployments lists all known deployment IDs
func (s *serverData) EnumerateDeployments(
	ctx context.Context,
	request *empty.Empty,
) (*pb.DeploymentsResponse, error) {
	deploymentIDSet := make(map[string]struct{})
	for _, platform := range s.platforms {
		for _, deploymentID := range platform.enumerateDeployments() {
			deploymentIDSet[deploymentID] = struct{}{}
		}
	}

	var response pb.DeploymentsResponse
	for deploymentID := range deploymentIDSet {
		response.DeploymentId = append(response.DeploymentId, deploymentID)
	}

	s.Logger.Debug().
		Str("method", "EnumerateDeployments").
		Int("deployments", len(response.DeploymentId)).
		Msg("EnumerateDeployments")

	return &response, nil
}
