package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// SetDeploymentStatus sets the desired deployment status
// Returns the known deployment status at time of execution
func (s *serverData) SetDeploymentStatus(ctx context.Context, request *pb.SetDeploymentStatusRequest) (*pb.DeploymentStatusResponse, error) {
	return nil, errors.New("not implemented")
}
