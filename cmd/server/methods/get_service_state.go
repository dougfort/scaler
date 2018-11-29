package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// GetServiceState lists the state of the service on all platforms
func (s *serverData) GetServiceState(
	_ context.Context,
	request *pb.GetServiceStateRequest,
) (*pb.GetServiceStateResponse, error) {
	var response pb.GetServiceStateResponse

	for platformID, platform := range s.platforms {
		data, isDeployed, err := platform.getServiceState(request.ServiceId)
		if err != nil {
			return nil, errors.Wrapf(err, "platform: %s; getServiceState(%s)",
				platformID, request.ServiceId)
		}
		if isDeployed {
			response.ServiceState = append(
				response.ServiceState,
				&pb.ServiceState{
					PlatformId:    platformID,
					InstanceCount: data.instanceCount,
				},
			)
		}
	}

	s.Logger.Debug().
		Str("method", "GetServiceState").
		Str("serviceID", request.ServiceId).
		Int("platforms-reporting", len(response.ServiceState)).
		Msg("GetServiceState")

	return &response, nil
}
