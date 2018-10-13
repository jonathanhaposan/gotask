#!/bin/bash

# Start the first process
./cmd/web/web & ./cmd/tcp/tcp;

echo "Called..."