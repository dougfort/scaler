package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"

	"google.golang.org/grpc"

	"github.com/dougfort/scaler/cmd/server/config"
	"github.com/dougfort/scaler/cmd/server/methods"

	// we don't use this directly, but need it in vendor for gateway grpc plugin
	_ "github.com/golang/glog"
	_ "github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/ghodss/yaml"
)

func main() {
	var tlsServerConf *tls.Config
	var err error

	logger := zerolog.New(os.Stdout).
		With().Timestamp().Str("service", "scaler").Logger().
		Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger.Info().Msg("starting")

	ctx, cancelFunc := context.WithCancel(context.Background())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger.Debug().Msg("initializing config")
	if err = config.Initialize(); err != nil {
		logger.Fatal().AnErr("config.Initialize()", err).Msg("")
	}

	if tlsServerConf, err = buildServerTLSConfigIfNeeded(logger); err != nil {
		logger.Fatal().AnErr("buildServerTLSConfigIfNeeded", err).Msg("")
	}

	logger.Debug().Str("host", viper.GetString("grpc_server_host")).
		Int("port", viper.GetInt("grpc_server_port")).
		Msg("creating listener")

	lis, err := net.Listen(
		"tcp",
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("grpc_server_host"),
			viper.GetInt("grpc_server_port"),
		),
	)
	if err != nil {
		logger.Fatal().AnErr("net.Listen", err).Msg("")
	}

	opts := getTLSOptsIfNeeded(tlsServerConf)
	grpcServer := grpc.NewServer(opts...)

	methods.CreateAndRegisterServer(logger, grpcServer)

	logger.Debug().Msg("starting grpc server")
	go grpcServer.Serve(lis)

	if viper.GetBool("use_gateway_proxy") {
		logger.Debug().Msg("starting gateway proxy")
		if err = startGatewayProxy(ctx, logger); err != nil {
			logger.Fatal().AnErr("startGatewayProxy", err).Msg("")
		}
	}

	s := <- sigChan
	logger.Info().Str("signal", s.String()).Msg("shutting down")
	cancelFunc()
	grpcServer.Stop()
}
