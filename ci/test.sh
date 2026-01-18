#!/usr/bin/env bash

set -euo pipefail

project_root=$(git rev-parse --show-toplevel)

cd "$project_root"

golangci-lint run
go test ./...

echo "All tests passed!"
