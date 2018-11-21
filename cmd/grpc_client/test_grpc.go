package main

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	pb "github.com/dougfort/scaler/protobuf"
)

func runTest(logger zerolog.Logger, client pb.ScalerClient) error {
	ctx := context.Background()

	deployments, err := client.EnumerateDeployments(ctx, &empty.Empty{}, []grpc.CallOption{}...)
	if err != nil {
		return errors.Wrap(err, "EnumerateDeployments")
	}

	for _, deploymentID := range deployments.DeploymentId {
		logger.Debug().Str("deploymentID", deploymentID).Msg("EnumerateDeployments")
	}

	platforms, err := client.EnumeratePlatforms(ctx, &empty.Empty{}, []grpc.CallOption{}...)
	if err != nil {
		return errors.Wrap(err, "EnumeratePlatforms")
	}

	for _, platformID := range platforms.PlatformId {
		logger.Debug().Str("platformID", platformID).Msg("EnumeratePlatforms")
	}

	return nil
}
