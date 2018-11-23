package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// GetDeploymentStatus lists the status of the deployment on all platforms
func (s *serverData) GetDeploymentStatus(
	ctx context.Context,
	request *pb.GetDeploymentStatusRequest,
) (*pb.GetDeploymentStatusResponse, error) {

	var response pb.GetDeploymentStatusResponse

	for platformID, platform := range s.platforms {
		data, isDeployed, err := platform.getDeploymentData(request.DeploymentId)
		if err != nil {
			return nil, errors.Wrapf(err, "platform: %s; GetDeploymentStatus(%s)",
				platformID, request.DeploymentId)
		}
		if isDeployed {
			response.DeploymentStatus = append(
				response.DeploymentStatus,
				&pb.DeploymentStatus{
					PlatformId:    platformID,
					InstanceCount: data.instanceCount,
				},
			)
		}
	}

	s.Logger.Debug().
		Str("method", "GetDeploymentStatus").
		Str("deploymentID", request.DeploymentId).
		Int("platforms-reporting", len(response.DeploymentStatus)).
		Msg("GetDeploymentStatus")

	return &response, nil
}
