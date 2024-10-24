#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -lt 3 ]; then
    echo "Usage: $0 <section>/<subject>/<lab_program> <auth_token> <file_path> <attachments...>"
    exit 1
fi

# Variables
API_PATH=$1
AUTH_TOKEN=$2
FILE_PATH=$3

# Initialize attachments variable
attachments=""

# Loop through remaining arguments as attachments
for arg in "${@:4}"; do
    attachments="$attachments -F \"attachments=@$arg\""
done

# Execute curl command
eval curl -X POST "http://localhost:8080/$API_PATH" \
    -H "Authorization: Bearer $AUTH_TOKEN"
    -H "Content-Type: multipart/form-data" \
    -F "file=@$FILE_PATH" \
    #$attachments \
    
