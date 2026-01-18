#!/usr/bin/env bash

set -euo pipefail

project_root=$(git rev-parse --show-toplevel)
target=$1

local_os=$(uname | tr '[:upper:]' '[:lower:]')
GOOS=${GOOS:-$local_os}
case $(uname -m) in
    x86_64) local_arch="amd64" ;;
    aarch64 | arm64) local_arch="arm64" ;;
    *) echo "Unsupported architecture: $(uname -m)" ; exit 1 ;;
esac
GOARCH=${GOARCH:-$local_arch}

echo "Building target '$target' for $GOOS/$GOARCH"
target_name="${target##*/}"
cd "$project_root"
mkdir -p dist
CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" go build -o "dist/${target_name}" "${target}"

PACKAGE_FORMAT="${PACKAGE_FORMAT:-tar.gz}"
package_name="${target_name}-${GOOS}-${GOARCH}.${PACKAGE_FORMAT}"
cd dist
if [ "$PACKAGE_FORMAT" == "tar.gz" ]; then
    tar -czf "$package_name" "$target_name"
elif [ "$PACKAGE_FORMAT" == "zip" ]; then
    zip "$package_name" "$target_name"
else
    echo "Unsupported package format: $PACKAGE_FORMAT"
    exit 1
fi
echo "Packaged $package_name"
