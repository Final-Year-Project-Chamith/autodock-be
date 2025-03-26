# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /autodock-be

# Stage 2: Create the runtime image
FROM debian:bullseye-slim

# Install required packages: Nginx, certbot, and its nginx plugin, curl, and ca-certificates
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

# Copy the built Go binary from the builder stage
COPY --from=builder /autodock-be /autodock-be

# Expose necessary ports
# Port 8888 is for the Go Fiber application and port 80 is used by Nginx (for HTTP challenges, etc.)
EXPOSE 8888 80

# Create an entrypoint script to start Nginx and then your application
RUN echo '#!/bin/bash\nservice nginx start\n/autodock-be' > /entrypoint.sh && chmod +x /entrypoint.sh

# Set the PATH to include /usr/sbin so certbot and nginx can be found
ENV PATH="/usr/sbin:$PATH"

CMD ["/entrypoint.sh"]
