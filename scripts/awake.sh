#!/bin/bash

# URL to ping
URL="https://studious-doodle.onrender.com"

while true
do
    # Record the start time
    start_time=$(date +%s%N)

    # Print ping message
    echo "$(date +%Y-%m-%d\ %H:%M:%S) ping"

    # Perform the HTTP GET request
    response=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

    # Check if the request was successful
    if [ "$response" -eq 200 ]; then
        echo "$(date +%Y-%m-%d\ %H:%M:%S) pong"
    else
        echo "$(date +%Y-%m-%d\ %H:%M:%S) error"
    fi

    # Calculate the response time
    end_time=$(date +%s%N)
    response_time=$(( (end_time - start_time) / 1000000 )) # Convert nanoseconds to milliseconds

    # Print the response time
    echo "$(date +%Y-%m-%d\ %H:%M:%S) Response time for restart: ${response_time}ms"
    echo "$(date +%Y-%m-%d\ %H:%M:%S) see you 5 mins later"

    # Sleep for 5 minutes
    sleep 300
done
