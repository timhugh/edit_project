#!/usr/bin/env bash

set -euo pipefail

project_root=$(git rev-parse --show-toplevel)

version=$1
gh release create "$version" ./dist/*.{tar.gz,zip} --notes "Release $version"
