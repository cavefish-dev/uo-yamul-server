#!/bin/bash
# Wrapper for running yamul-gateway tests.
# Usage: run-tests.sh [go test flags...]
# Example: run-tests.sh -v ./...
# Example (integration): run-tests.sh -tags integration -v -timeout 30s ./integration/...
# Example (GRPC debug):  GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info run-tests.sh -tags integration -v ./integration/...

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../.." && pwd)"
cd "$REPO_ROOT/yamul-gateway"
exec go test "$@"
