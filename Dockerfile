# Stage 1: Build the Go binary along with all project files
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project including templates
COPY . .

# Build the Go application
RUN go build -o /autodock-be

# Stage 2: Create the runtime image
FROM debian:bullseye-slim

# Install required packages: Nginx, certbot and its nginx plugin, curl, and certificates
RUN apt-get update && apt-get install -y \
    nginx \
    python3-certbot-nginx \
    certbot \
    curl \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Install Docker CLI and Docker Compose
RUN curl -fsSL https://get.docker.com -o get-docker.sh && \
    sh get-docker.sh && \
    curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" \
    -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose && \
    rm get-docker.sh

# Set working directory in the final image
WORKDIR /app

# Copy the built binary and the templates directory from the builder stage
COPY --from=builder /autodock-be /autodock-be
COPY --from=builder /app/templates ./templates

# Expose application and HTTP challenge ports
EXPOSE 8888 80

# Run the application
CMD ["/autodock-be"]