#!/bin/bash
set -e

# Configuration
IMAGE_NAME="pg17-4-timescaledb"
VERSION="2.19.3-slim"
REGISTRY="docker.io/cedrickiama/"  # Using your Docker Hub account

# Build the image
echo "Building image: ${REGISTRY}${IMAGE_NAME}:${VERSION}"
docker build --no-cache -t "${REGISTRY}${IMAGE_NAME}:${VERSION}" .

# Tag latest
docker tag "${REGISTRY}${IMAGE_NAME}:${VERSION}" "${REGISTRY}${IMAGE_NAME}:latest"

# Display image size
echo "Image size:"
docker images "${REGISTRY}${IMAGE_NAME}:${VERSION}" --format "{{.Size}}"

# Push to Docker Hub (uncomment for production deployment)
 docker push "${REGISTRY}${IMAGE_NAME}:${VERSION}"
 docker push "${REGISTRY}${IMAGE_NAME}:latest"

echo "Build completed successfully!"
