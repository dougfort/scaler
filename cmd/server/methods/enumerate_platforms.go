package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
)

// EnumeratePlatforms lists all known platform IDs
func (s *serverData) EnumeratePlatforms(ctx context.Context, request *empty.Empty) (*pb.PlatformsResponse, error) {
	return nil, errors.New("not implemented")
}
