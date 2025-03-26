# Stage 1: Build the Go binary statically
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .
COPY ./templates ./templates
RUN go mod tidy
# Disable CGO to produce a fully static binary
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /autodock-be .

# Stage 2: Create the runtime image
FROM debian:bullseye-slim

# Install required packages: Nginx, certbot and its plugin, and other dependencies
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

# Copy the statically built binary from the builder stage
COPY --from=builder /autodock-be /autodock-be

# Expose necessary ports: 8888 for your app and 80 for Nginx
EXPOSE 8888 80

# Create an entrypoint script to start Nginx and your app
RUN echo '#!/bin/bash\nservice nginx start\n/autodock-be' > /entrypoint.sh && chmod +x /entrypoint.sh

CMD ["/entrypoint.sh"]
