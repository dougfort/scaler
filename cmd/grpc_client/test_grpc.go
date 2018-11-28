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

	services, err := client.EnumerateServices(
		ctx,
		&empty.Empty{},
		[]grpc.CallOption{}...,
	)
	if err != nil {
		return errors.Wrap(err, "EnumerateServices")
	}
	serviceIDs := services.ServiceId

	for _, serviceID := range serviceIDs {
		logger.Debug().Str("serviceID", serviceID).Msg("EnumerateServices")
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
	testServiceID := serviceIDs[0]

	getResponse, err := client.GetServiceState(
		ctx,
		&pb.GetServiceStateRequest{ServiceId: testServiceID},
	)
	if err != nil {
		return errors.Wrapf(err, "GetServiceState(%s) failed", testServiceID)
	}

	testServiceState := getResponse.ServiceState

	for _, ds := range testServiceState {
		logger.Debug().Msgf("platform: %s; instances = %d", ds.PlatformId, ds.InstanceCount)
		ds.InstanceCount += 1
	}

	setResponse, err := client.SetServiceState(
		ctx,
		&pb.SetServiceStateRequest{
			ServiceId:    testServiceID,
			ServiceState: testServiceState,
		},
	)
	if err != nil {
		return errors.Wrapf(err, "SetServiceState(%s) failed", testServiceID)
	}

	logger.Debug().Msgf("setResponse = %v", setResponse)

	getResponse, err = client.GetServiceState(
		ctx,
		&pb.GetServiceStateRequest{ServiceId: testServiceID},
	)
	if err != nil {
		return errors.Wrapf(err, "GetServiceState(%s) failed", testServiceID)
	}

	for _, ds := range getResponse.ServiceState {
		logger.Debug().Msgf("platform: %s; instances = %d", ds.PlatformId, ds.InstanceCount)
	}

	return nil
}
