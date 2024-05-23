#!/bin/bash

# Build the price-aggregator binary
echo "Building price-aggregator..."
go build -o price-aggregator ./cmd/server

if [ $? -eq 0 ]; then
    echo "Build successful."
else
    echo "Build failed."
    exit 1
fi
