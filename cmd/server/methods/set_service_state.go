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

	requestID, requestItem := s.newRequest(request.ServiceId, request.ServiceState)
	s.requests[requestID] = requestItem

	s.Logger.Debug().
		Str("service", request.ServiceId).
		Str("request-id", requestID).
		Msg("SetServiceState")

	response.RequestId = requestID
	return &response, nil
}

// finishSetServiceState simulates a long-running set of changes
func (s *serverData) finishSetServiceState(requestID string, req requestStruct) error {
	var platformsFound int
	var platformsSet int

SERVICE_STATE_LOOP:
	for _, ds := range req.serviceState {
		platform, ok := s.platforms[ds.PlatformId]
		if !ok {
			s.Logger.Warn().
				Str("request-id", requestID).
				Str("service", req.serviceID).
				Str("platform", ds.PlatformId).
				Msgf("unknown platform: %s", ds.PlatformId)
			continue SERVICE_STATE_LOOP
		}
		platformsFound++

		ok, err := platform.setServiceState(
			req.serviceID,
			serviceData{
				instanceCount: ds.InstanceCount,
			},
		)

		if err != nil {
			return errors.Wrapf(
				err, "platform: %s; setServiceState(%s) failed",
				ds.PlatformId, req.serviceID)
		}

		if !ok {
			s.Logger.Warn().
				Str("request-id", requestID).
				Str("service", req.serviceID).
				Str("platform", ds.PlatformId).
				Msgf("unknown service: %s", req.serviceID)
			continue SERVICE_STATE_LOOP
		}

		platformsSet++
	}

	s.Logger.Debug().
		Str("request-id", requestID).
		Str("service", req.serviceID).
		Int("platforms-found", platformsFound).
		Int("platforms-set", platformsSet).
		Msg("finishSetServiceState")

	return nil
}
