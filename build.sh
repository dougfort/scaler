#!/bin/bash

set -euxo pipefail

GOPATH="${GOPATH:-$HOME/go}"
GOPATH=$(echo $GOPATH | cut -d: -f1)

TOPDIR=$GOPATH/src/github.com/dougfort/scaler
CLIENTDIR="${TOPDIR}/cmd/grpc_client"
SERVERDIR="${TOPDIR}/cmd/server"

(
	cd $SERVERDIR
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o=$GOPATH/bin/scaler
	DOCKERDIR="${SERVERDIR}/docker"
	cp $GOPATH/bin/scaler "$DOCKERDIR/."
	cp "settings.toml" "${DOCKERDIR}/."
	cd $DOCKERDIR
	docker build -t scaler .
)

(
	cd $CLIENTDIR
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o=$GOPATH/bin/scaler_client
	DOCKERDIR="${CLIENTDIR}/docker"
	cp $GOPATH/bin/scaler_client "$DOCKERDIR/."
	cd $DOCKERDIR
	docker build -t scaler_client .
)
