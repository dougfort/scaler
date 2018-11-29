package methods

import (
	"google.golang.org/grpc"

	pb "github.com/dougfort/scaler/protobuf"
	"github.com/rs/zerolog"
)

type serverData struct {
	zerolog.Logger
	nextRequestID int64 // We should use a UUID
	requests      map[string]requestStruct
	platforms     map[string]platformData
}

// CreateAndRegisterServer returns an object that implements the  interface
func CreateAndRegisterServer(
	logger zerolog.Logger,
	grpcServer *grpc.Server,
) {
	s := &serverData{
		Logger:        logger,
		nextRequestID: 1,
		requests:      make(map[string]requestStruct),
		platforms:     make(map[string]platformData),
	}

	s.loadTestData()

	pb.RegisterScalerServer(grpcServer, s)
	s.Logger.Debug().Str("method", "CreateAndRegisterServer").Msg("registered")
}
