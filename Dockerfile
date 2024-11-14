# Start from the official Golang image
FROM golang:1.22

# Set up the working directory
WORKDIR /app

# Copy and install Go dependencies
COPY . .
RUN go mod tidy
RUN go get

# Compile the Go application
RUN go build -o /autodock-be

# Install Nginx
RUN apt-get update && apt-get install -y nginx
# Set the PATH environment variable to include /usr/sbin
RUN apt-get update && \
    apt-get install -y python3-certbot-nginx

ENV PATH="/usr/sbin:$PATH"

RUN apt-get update && apt-get install -y certbot


# Install Docker CLI and Docker Compose
RUN apt-get update && \
    apt-get install -y curl && \
    curl -fsSL https://get.docker.com -o get-docker.sh && \
    sh get-docker.sh && \
    curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" \
    -o /usr/local/bin/docker-compose && \
    chmod +x /usr/local/bin/docker-compose && \
    rm get-docker.sh

# Expose the application port
EXPOSE 8888

# Command to run the application
CMD [ "/autodock-be" ]
