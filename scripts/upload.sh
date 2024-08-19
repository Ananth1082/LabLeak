#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 "\<section\>/\<subject\>/\<lab_program\>" <file_path> <auth_token>"
    exit 1
fi

# Variables
API_PATH=$1
FILE_PATH=$2
AUTH_TOKEN=$3

# Execute curl command
curl -X POST "https://localhost:8080/$API_PATH" \
     -H "Content-Type: multipart/form-data" \
     -F "file=@$FILE_PATH" \
     -H "Authorization: Bearer $AUTH_TOKEN"
