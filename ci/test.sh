#!/usr/bin/env bash

set -euo pipefail

project_root=$(git rev-parse --show-toplevel)
cd "$project_root"

echo "Preparing test environment..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

echo "Running linters and tests..."
golangci-lint run
go test ./...
echo "All tests passed!"
