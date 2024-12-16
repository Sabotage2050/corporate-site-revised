# Base stage for common setup
FROM golang:1.23 AS base

WORKDIR /app

# Create a non-root user
RUN useradd -u 1000 -m appuser

# Set up a global GOPATH and PATH
ENV GOPATH /home/appuser/go
ENV PATH $GOPATH/bin:$PATH

# Switch to the non-root user
USER appuser

# Install development tools in a separate layer
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/air-verse/air@latest
# Cache dependencies to avoid reinstalling on every build
COPY --chown=appuser:appuser ./go.mod ./go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Development stage
FROM base AS development

WORKDIR /app
EXPOSE 8080 40000

# Use Air for hot-reload
CMD ["air", "-c", ".air.toml"]

# Test-debug stage
FROM base AS test-debug

WORKDIR /app
EXPOSE 8080 40000

# Run Delve in test mode
CMD ["sh", "-c", "dlv test --headless --listen=:40000 --api-version=2 --log ./tests/feature"]

# Builder stage for compiling - using root user for build
FROM golang:1.23-bullseye AS builder
WORKDIR /app

# Keep root user for build process
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# CI stage - minimal image for testing
FROM golang:1.23-bullseye AS ci
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
RUN useradd -u 1000 -m appuser

# Copy only necessary files for testing
COPY --from=builder /app /app

# Set permissions after copying
RUN chown -R appuser:appuser /app

USER appuser
EXPOSE 8080
# Keep container running for debugging
CMD ["tail", "-f", "/dev/null"]
# CMD ["./main"]

# Staging stage - Similar to production but with additional debugging capabilities
FROM debian:bullseye-slim AS staging
WORKDIR /app

# Install additional packages for debugging and monitoring
RUN apt-get update && \
    apt-get install -y \
    ca-certificates \
    curl \
    htop \
    netcat \
    procps \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -u 1000 -m appuser

# Copy the compiled binary and necessary files
COPY --from=builder /app/main .
# COPY .env.staging ./.env
COPY data ./data

# Set permissions after copying
RUN chown -R appuser:appuser /app

USER appuser

# Environment variable to indicate staging environment
ENV DEBUG=true

EXPOSE 8080
# Additional port for debugging if needed
EXPOSE 8081

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# CMD ["tail", "-f", "/dev/null"]
CMD ["./main"]

# Production stage - minimal image for running
FROM debian:bullseye-slim AS production
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
RUN useradd -u 1000 -m appuser

# Copy only the compiled binary and necessary files
COPY --from=builder /app/main .
COPY .env* ./
COPY data ./data

# Set permissions after copying
RUN chown -R appuser:appuser /app

USER appuser
EXPOSE 8080
CMD ["./main"]




























