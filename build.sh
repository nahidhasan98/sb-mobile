#!/bin/bash

# Create bin directory if it doesn't exist
mkdir -p bin

# Build for Linux (production)
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
go build -ldflags "-s -w" -o bin/sb-mobile

echo "Build complete: bin/sb-mobile"
