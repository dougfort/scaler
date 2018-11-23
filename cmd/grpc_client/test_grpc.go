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

	deployments, err := client.EnumerateDeployments(
		ctx,
		&empty.Empty{},
		[]grpc.CallOption{}...,
	)
	if err != nil {
		return errors.Wrap(err, "EnumerateDeployments")
	}
	deploymentIDs := deployments.DeploymentId

	for _, deploymentID := range deploymentIDs {
		logger.Debug().Str("deploymentID", deploymentID).Msg("EnumerateDeployments")
	}

	platforms, err := client.EnumeratePlatforms(
		ctx,
		&empty.Empty{},
		[]grpc.CallOption{}...,
	)
	if err != nil {
		return errors.Wrap(err, "EnumeratePlatforms")
	}
	platformIDs := platforms.PlatformId

	for _, platformID := range platformIDs {
		logger.Debug().Str("platformID", platformID).Msg("EnumeratePlatforms")
	}

	// TODO: pick a random id
	testDeploymentID := deploymentIDs[0]

	getResponse, err := client.GetDeploymentStatus(
		ctx,
		&pb.GetDeploymentStatusRequest{DeploymentId: testDeploymentID},
	)
	if err != nil {
		return errors.Wrapf(err, "GetDeploymentStatus(%s) failed", testDeploymentID)
	}

	testDeploymentStatus := getResponse.DeploymentStatus

	for _, ds := range testDeploymentStatus {
		logger.Debug().Msgf("platform: %s; instances = %d", ds.PlatformId, ds.InstanceCount)
		ds.InstanceCount += 1
	}

	setResponse, err := client.SetDeploymentStatus(
		ctx,
		&pb.SetDeploymentStatusRequest{
			DeploymentId:     testDeploymentID,
			DeploymentStatus: testDeploymentStatus,
		},
	)
	if err != nil {
		return errors.Wrapf(err, "SetDeploymentStatus(%s) failed", testDeploymentID)
	}

	logger.Debug().Msgf("setResponse = %v", setResponse)

	getResponse, err = client.GetDeploymentStatus(
		ctx,
		&pb.GetDeploymentStatusRequest{DeploymentId: testDeploymentID},
	)
	if err != nil {
		return errors.Wrapf(err, "GetDeploymentStatus(%s) failed", testDeploymentID)
	}

	for _, ds := range getResponse.DeploymentStatus {
		logger.Debug().Msgf("platform: %s; instances = %d", ds.PlatformId, ds.InstanceCount)
	}

	return nil
}
