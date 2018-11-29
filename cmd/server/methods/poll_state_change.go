package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// PollStateChange checks the status of a SetServiceState request
func (s *serverData) PollStateChange(
	_ context.Context,
	request *pb.PollStateChangeRequest,
) (*pb.PollStateChangeResponse, error) {
	var response pb.PollStateChangeResponse

	req, ok := s.requests[request.RequestId]
	if !ok {
		return nil, errors.Errorf("No active request for id %s",
			request.RequestId)
	}

	// the first time through, pretend we aren't done
	if !req.finished {
		s.Logger.Debug().
			Str("service", req.serviceID).
			Str("request-id", request.RequestId).
			Bool("finished", req.finished).
			Msg("PollStateChange")

		req.finished = true
		s.requests[request.RequestId] = req

		response.Message = "in progress"
		return &response, nil
	}

	delete(s.requests, request.RequestId)

	err := s.finishSetServiceState(request.RequestId, req)

	// now, really finish the request
	response.IsCompleted = true
	if err != nil {
		response.IsSuccessful = false
		response.Message = err.Error()
	} else {
		response.IsSuccessful = true
	}

	return &response, nil
}
