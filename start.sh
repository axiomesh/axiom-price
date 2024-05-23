#!/bin/bash

# Start the price-aggregator service
echo "Starting price-aggregator..."
nohup ./price-aggregator > price-aggregator.log 2>&1 &

# Save the process ID
echo $! > price-aggregator.pid

echo "price-aggregator started with PID $(cat price-aggregator.pid)."
