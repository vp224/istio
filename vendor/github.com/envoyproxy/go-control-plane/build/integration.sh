#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

##
## Integration testing script for the control plane library against the Envoy binary.
## This is a wrapper around the test app `pkg/test/main` that spawns/kills Envoy.
##

# Management server type. Valid values are "ads", "xds", "rest"
XDS=${XDS:-ads}

(bin/test --xds=${XDS} "$@")&
SERVER_PID=$!

# Envoy start-up command
ENVOY=${ENVOY:-/usr/local/bin/envoy}

# Start envoy: important to keep drain time short
(${ENVOY} -c sample/bootstrap-${XDS}.yaml --drain-time-s 1 --v2-config-only -l debug 2> envoy.log)&
ENVOY_PID=$!

function cleanup() {
  kill ${ENVOY_PID}
}
trap cleanup EXIT

wait ${SERVER_PID}
