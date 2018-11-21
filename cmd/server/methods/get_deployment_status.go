package methods

import (
	"golang.org/x/net/context"

	"github.com/pkg/errors"

	pb "github.com/dougfort/scaler/protobuf"
)

// GetDeploymentStatus lists the status of the deployment on all platforms
func (s *serverData) GetDeploymentStatus(ctx context.Context, request *pb.GetDeploymentStatusRequest) (*pb.DeploymentStatusResponse, error) {
	return nil, errors.New("not implemented")
}
