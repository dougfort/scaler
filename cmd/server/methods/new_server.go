package methods

import (
	"google.golang.org/grpc"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/rs/zerolog"
)

type serverData struct {
	zerolog.Logger
	DeploymentStatusMap map[string][]DeploymentStatus
	PlatformIDs         []string
}

// CreateAndRegisterServer returns an object that implements the  interface
func CreateAndRegisterServer(
	logger zerolog.Logger,
	grpcServer *grpc.Server,
) {
	s := &serverData{
		Logger:              logger,
		DeploymentStatusMap: make(map[string][]DeploymentStatus),
	}

	s.loadTestData()

	pb.RegisterScalerServer(grpcServer, s)
	s.Logger.Debug().Str("method", "CreateAndRegisterServer").Msg("registered")
}
