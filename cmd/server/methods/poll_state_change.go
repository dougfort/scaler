package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// PollStateChange checks the status of a SetServiceState request
func (s *serverData) PollStateChange(ctx context.Context, request *pb.PollStateChangeRequest) (*pb.PollStateChangeResponse, error) {
	return nil, errors.New("not implemented")
}
