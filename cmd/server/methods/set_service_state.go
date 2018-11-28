package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// SetServiceState initiates a request to change the state of a service
// Returns a unique request_id for tracking the status of the request
func (s *serverData) SetServiceState(
	_ context.Context,
	request *pb.SetServiceStateRequest,
) (*pb.SetServiceStateResponse, error) {
	var response pb.SetServiceStateResponse
	var platformsFound int
	var platformsSet int

SERVICE_STATE_LOOP:
	for _, ds := range request.ServiceState {
		platform, ok := s.platforms[ds.PlatformId]
		if !ok {
			s.Logger.Warn().
				Str("service", request.ServiceId).
				Str("platform", ds.PlatformId).
				Msgf("unknown platform: %s", ds.PlatformId)
			continue SERVICE_STATE_LOOP
		}
		platformsFound++

		ok, err := platform.setServiceState(
			request.ServiceId,
			serviceData{
				instanceCount: ds.InstanceCount,
			},
		)

		if err != nil {
			return nil, errors.Wrapf(
				err, "platform: %s; setServiceState(%s) failed",
				ds.PlatformId, request.ServiceId)
		}

		if !ok {
			s.Logger.Warn().
				Str("service", request.ServiceId).
				Str("platform", ds.PlatformId).
				Msgf("unknown service: %s", request.ServiceId)
			continue SERVICE_STATE_LOOP
		}

		platformsSet++
	}

	s.Logger.Debug().
		Str("service", request.ServiceId).
		Int("platforms-found", platformsFound).
		Int("platforms-set", platformsSet).
		Msg("SetServiceState")

	return &response, nil
}
