package methods

import (
	"golang.org/x/net/context"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
)

// EnumeratePlatforms lists all known platform IDs
func (s *serverData) EnumeratePlatforms(
	ctx context.Context,
	request *empty.Empty,
) (*pb.PlatformsResponse, error) {
	var response pb.PlatformsResponse

	for platformID := range s.platforms {
		response.PlatformId = append(response.PlatformId, platformID)
	}

	s.Logger.Debug().
		Str("method", "EnumeratePlatforms").
		Int("platforms", len(response.PlatformId)).
		Msg("EnumeratePlatforms")

	return &response, nil
}
