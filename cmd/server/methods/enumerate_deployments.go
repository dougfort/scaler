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
	var response pb.DeploymentsResponse

	for deploymentID := range s.DeploymentStatusMap {
		response.DeploymentId = append(response.DeploymentId, deploymentID)
	}

	s.Logger.Debug().
		Str("method", "EnumerateDeployments").
		Int("deployments", len(response.DeploymentId)).
		Msg("EnumerateDeployments")

	return &response, nil
}
