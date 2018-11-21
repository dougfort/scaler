package methods

import (
	"google.golang.org/grpc"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/rs/zerolog"
)

type serverData struct {
	zerolog.Logger
}

// NewScalerServer returns an object that implements the  interface
func CreateAndRegisterServer(
	logger zerolog.Logger,
	grpcServer *grpc.Server,
) {
	var server pb.ScalerServer = &serverData{
		logger,
	}

	pb.RegisterScalerServer(grpcServer, server)

}
