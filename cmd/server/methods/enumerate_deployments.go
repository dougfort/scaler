package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
)

// EnumerateDeployments lists all known deployment IDs
func (s *serverData) EnumerateDeployments(ctx context.Context, request *empty.Empty) (*pb.DeploymentsResponse, error) {
	return nil, errors.New("not implemented")
}
