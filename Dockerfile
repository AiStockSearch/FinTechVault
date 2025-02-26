# Stage 1: Build the application
FROM golang:1.20 AS builder
# Set working directory
WORKDIR /app
# Copy go.mod and go.sum for dependency management
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Copy the rest of the code
COPY . .
# Build the Wails application
RUN wails build -p -d

# Stage 2: Create a minimal runtime image
FROM debian:bullseye-slim
# Install necessary dependencies
RUN apt-get update && \
    apt-get install -y libgtk-3-0 libayatana-appindicator3-1 libappindicator3-1 && \
    rm -rf /var/lib/apt/lists/*
# Set working directory
WORKDIR /app
# Copy the built binary from the builder stage
COPY --from=builder /app/build/* .
# Expose the app as an entrypoint
ENTRYPOINT ["./fintech"]