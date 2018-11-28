package methods

import (
	"golang.org/x/net/context"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
)

// EnumerateServices lists all known service IDs
func (s *serverData) EnumerateServices(
	_ context.Context,
	_ *empty.Empty,
) (*pb.EnumerateServicesResponse, error) {
	serviceIDSet := make(map[string]struct{})
	for _, platform := range s.platforms {
		for _, serviceID := range platform.enumerateServices() {
			serviceIDSet[serviceID] = struct{}{}
		}
	}

	var response pb.EnumerateServicesResponse
	for serviceID := range serviceIDSet {
		response.ServiceId = append(response.ServiceId, serviceID)
	}

	s.Logger.Debug().
		Str("method", "EnumerateServices").
		Int("services", len(response.ServiceId)).
		Msg("EnumerateServices")

	return &response, nil
}
