#!/bin/bash

set -euxo pipefail

# assume we are in the service directory
SERVICEDIR=$PWD
CLIENTDIR="${SERVICEDIR}/cmd/http_client"

(
	cd $CLIENTDIR
	go build -o=$GOPATH/bin/scaler_http_client
)
