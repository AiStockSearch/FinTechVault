# Stage 1: Build the application
FROM golang:1.20-bullseye AS builder

# Add labels for better maintainability
LABEL maintainer="Developer"
LABEL version="1.0"
LABEL description="FinTech Wails Application"

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Wails application
RUN wails build -p -d && \
    chmod +x build/fintech

# Stage 2: Create a minimal runtime image
FROM debian:bullseye-slim

# Install necessary dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    libgtk-3-0 \
    libayatana-appindicator3-1 \
    libappindicator3-1 \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    groupadd -r appuser && \
    useradd -r -g appuser appuser

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/build/* .

# Change ownership of the application files
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Add healthcheck
HEALTHCHECK --interval=30s --timeout=3s \
  CMD pgrep fintech || exit 1

# Expose the app as an entrypoint
ENTRYPOINT ["./fintech"]