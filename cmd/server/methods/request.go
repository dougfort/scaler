package methods

import (
	"fmt"

	pb "github.com/dougfort/scaler/protobuf"
)

type requestStruct struct {
	serviceID    string
	serviceState []pb.ServiceState
	finished     bool
}

func (s *serverData) newRequest(
	serviceID string,
	serviceState []*pb.ServiceState,
) (string, requestStruct) {
	var req requestStruct

	req.serviceID = serviceID
	for _, s := range serviceState {
		req.serviceState = append(req.serviceState, *s)
	}
	requestID := fmt.Sprintf("%08d", s.nextRequestID)
	s.nextRequestID++

	return requestID, req
}
