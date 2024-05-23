#!/bin/bash

# Stop the price-aggregator service
if [ -f price-aggregator.pid ]; then
    PID=$(cat price-aggregator.pid)
    echo "Stopping price-aggregator with PID $PID..."
    kill $PID
    if [ $? -eq 0 ]; then
        echo "price-aggregator stopped."
        rm price-aggregator.pid
    else
        echo "Failed to stop price-aggregator."
    fi
else
    echo "price-aggregator is not running (no PID file found)."
fi