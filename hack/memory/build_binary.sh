#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

BINARY="simulate-memory"
OUTPUT="_output"
BIN_PATH="bin"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${OUTPUT}/${BIN_PATH}/${BINARY} pkg/simulate-memory/simulate_memory.go