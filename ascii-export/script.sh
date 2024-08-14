#!/bin/bash

# Stop ascii-art container if running
docker stop ascii-art-container >/dev/null
# Remove the container
docker rm ascii-art-container >/dev/null
# Delete the ascii-art image if exists
docker rmi ascii-art >/dev/null
#Build docker image
docker build -t ascii-art .
#Run the image on ascii-art-container on port 8080
docker run -p 8080:8080 --name ascii-art-container ascii-art
