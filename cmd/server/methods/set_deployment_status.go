package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// SetDeploymentStatus sets the desired deployment status
func (s *serverData) SetDeploymentStatus(
	ctx context.Context,
	request *pb.SetDeploymentStatusRequest,
) (*pb.SetDeploymentStatusResponse, error) {
	var response pb.SetDeploymentStatusResponse

DEPLOYMENT_STATUS_LOOP:
	for _, ds := range request.DeploymentStatus {
		platform, ok := s.platforms[ds.PlatformId]
		if !ok {
			s.Logger.Warn().
				Str("deployment", request.DeploymentId).
				Str("platform", ds.PlatformId).
				Msgf("unknown platform: %s", ds.PlatformId)
			continue DEPLOYMENT_STATUS_LOOP
		}
		response.PlatformsFound++

		ok, err := platform.setDeploymentData(
			request.DeploymentId,
			deploymentData{
				instanceCount: ds.InstanceCount,
			},
		)

		if err != nil {
			return nil, errors.Wrapf(
				err, "platfrom: %s; setDeploymentData(%s) failed", ds.PlatformId, request.DeploymentId)
		}

		if !ok {
			s.Logger.Warn().
				Str("deployment", request.DeploymentId).
				Str("platform", ds.PlatformId).
				Msgf("unknown deployment: %s", request.DeploymentId)
			continue DEPLOYMENT_STATUS_LOOP
		}

		response.PlatformsSet++
	}

	return &response, nil
}
