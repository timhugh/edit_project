#!/usr/bin/env bash

set -euo pipefail

project_root=$(git rev-parse --show-toplevel)

version="$1"
expected_version=$(cat "$project_root/VERSION")
if [[ "$version" != "$expected_version" ]]; then
    echo "VERSION file does not match. expected '$expected_version', got '$version'"
    exit 1
fi

exists=$(gh release list --json tagName --jq ".[] | select(.tagName == \"$version\")")
if [ -n "$exists" ]; then
    echo "Release with tag $version already exists"
    exit 1
fi

echo "Version $version is valid and does not exist yet"
