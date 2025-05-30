#!/bin/bash
set -e

# Configuration
IMAGE_NAME="pg17-4-timescaledb"
VERSION="2.19.3-slim"
REGISTRY="docker.io/cedrickiama/"

# Get build date and git commit for image labels
BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
VCS_REF=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Print build information
echo "==================================================================="
echo "Building secure PostgreSQL/TimescaleDB image with the following info:"
echo "IMAGE: ${REGISTRY}${IMAGE_NAME}:${VERSION}"
echo "BUILD DATE: ${BUILD_DATE}"
echo "VCS REF: ${VCS_REF}"
echo "==================================================================="

# Ensure buildx is set up with containerd driver
echo "Setting up buildx with containerd driver..."
if ! docker buildx inspect containerd-builder >/dev/null 2>&1; then
    docker buildx create --name containerd-builder --driver docker-container --use
else
    # Reset the existing builder to avoid potential stale cache issues
    docker buildx rm containerd-builder
    docker buildx create --name containerd-builder --driver docker-container --use
fi
docker buildx inspect --bootstrap containerd-builder

# Build with SBOM and provenance using buildx
echo "Building image with SBOM and provenance using buildx..."
docker buildx build \
  --builder containerd-builder \
  --provenance=true \
  --sbom=true \
  --no-cache \
  --platform linux/amd64 \
  --push \
  --build-arg BUILD_DATE="${BUILD_DATE}" \
  --build-arg VCS_REF="${VCS_REF}" \
  --build-arg VERSION="${VERSION}" \
  -t "${REGISTRY}${IMAGE_NAME}:${VERSION}" .

# Tag latest
docker tag "${REGISTRY}${IMAGE_NAME}:${VERSION}" "${REGISTRY}${IMAGE_NAME}:latest"

# Display image size
echo "Image size:"
docker images "${REGISTRY}${IMAGE_NAME}:${VERSION}" --format "{{.Size}}"

# Run security scan if Docker Scout is available
if docker scout --version >/dev/null 2>&1; then
    # Run security scan
    echo "==================================================================="
    echo "Running Docker Scout security scan..."
    docker scout cves "${REGISTRY}${IMAGE_NAME}:${VERSION}" --only-severity critical,high || {
        echo "Docker Scout scan found critical/high vulnerabilities. Review the report and address issues."
        exit 1
    }

    # Run recommendations scan
    echo "==================================================================="
    echo "Getting Docker Scout recommendations..."
    docker scout recommendations "${REGISTRY}${IMAGE_NAME}:${VERSION}" || echo "Docker Scout recommendations failed, continuing build..."
else
    echo "==================================================================="
    echo "Docker Scout not available. Skipping security scanning."
    echo "Consider installing Docker Desktop or the Scout CLI for security scanning."
fi

# Prompt for pushing to registry (only latest tag, since --push already handled main tag)
echo "==================================================================="
read -p "Do you want to push the 'latest' tag to Docker Hub? (y/n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    # Push latest tag to Docker Hub
    echo "Pushing latest tag to Docker Hub..."
    docker push "${REGISTRY}${IMAGE_NAME}:latest"
    echo "Latest tag pushed successfully!"
else
    echo "Skipping push of latest tag to Docker Hub."
fi

# Run final compliance check if Docker Scout is available
if docker scout --version >/dev/null 2>&1; then
    echo "==================================================================="
    echo "Running final compliance check..."
    docker scout quickview "${REGISTRY}${IMAGE_NAME}:${VERSION}" || echo "Docker Scout quickview failed, continuing..."
fi

echo "Build completed successfully!"
echo "==================================================================="

# Provide instructions for running with docker-compose
echo "To run the image with docker-compose:"
echo "docker-compose up -d"
echo "==================================================================="
