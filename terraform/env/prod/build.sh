#!/bin/bash
set -e

# Navigate to the lambda directory
cd ../../../lambda || exit 1

# Check if go.mod exists
if [ ! -f go.mod ]; then
    echo "go.mod not found. Initializing Go module..."
    go mod init lambda-function
fi

# Build the Go binary
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap -tags lambda.norpc main.go

echo "Build process completed successfully."