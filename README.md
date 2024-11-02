# autodock-be

`autodock-be` is a backend service designed for automating Docker Compose deployments. This tool enables users to visually create and configure Docker Compose files, link their repositories, and automate deployment processes. By integrating with GitHub, the tool can detect updates in specified repositories through webhooks, triggering automated build and deployment sequences.

## Table of Contents
- [Features](#features)
- [Endpoints](#endpoints)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Docker Deployment](#docker-deployment)
- [Configuration](#configuration)
- [Usage](#usage)
- [License](#license)

## Features
- **GitHub Webhook Integration**: Set up webhooks to trigger automated deployments on code changes.
- **Docker Management**: View and manage all Docker images and containers on the server.
- **Deployment Automation**: Deploy services using Docker Compose directly from the API, minimizing manual intervention.
- **File Generation**: Dynamically generate Docker Compose files based on user configurations.
- **Nginx Configuration Generator**: Automatically generate and manage Nginx configurations to route traffic to deployed containers.
- **Log Management**: Retrieve and manage logs for both Docker containers and the application to simplify debugging and monitoring.

### Planned Features
- **Advanced Webhook Management**: Configure multiple webhooks for different repositories or branches.
- **Container Health Monitoring**: Real-time status updates for container health and resource usage.
- **Access Control & Authentication**: Implement API key-based access for secure usage.
- **Advanced Configuration Management**: Manage environment variables and secrets used in Docker Compose files.
- **Resource Usage Reports**: Generate usage and performance reports for managed services.

## Endpoints

| Endpoint                                   | Method | Description                                |
|--------------------------------------------|--------|--------------------------------------------|
| `/autodock-be/api/CreateWebHook`           | POST   | Creates a webhook in a specified GitHub repository. |
| `/autodock-be/api/webhook`                 | POST   | Handles incoming GitHub webhook events.    |
| `/autodock-be/api/GetAllDockerImages`      | GET    | Retrieves a list of all Docker images on the server. |
| `/autodock-be/api/GetAllDockerContainers`  | GET    | Retrieves a list of all Docker containers on the server. |
| `/autodock-be/api/Deploy`                  | POST   | Deploys a service using a specified Docker Compose file. |
| `/autodock-be/api/generate/file/docker-compose` | POST | Generates a Docker Compose file based on user input. |
| `/autodock-be/api/GetLogs`                 | GET    | Retrieves logs for a specific container (Planned). |
| `/autodock-be/api/generate/nginx-config`   | POST   | Generates Nginx configuration for routing (Planned). |
| `/autodock-be/api/MonitorContainerHealth`  | GET    | Retrieves the health status of containers (Planned). |
| `/autodock-be/api/AccessControl`           | POST   | Manages API access control for secure usage (Planned). |
| `/`                                        | GET    | Returns a status message confirming the service is running. |

## Technologies Used
- **Go (Golang)**: Backend API development.
- **Fiber**: Web framework for API routing and middleware.
- **Docker**: Containerization for service management.
- **GitHub Webhooks**: Triggers for automated deployments based on repository updates.
- **Nginx**: Reverse proxy configuration generator for routing to containerized services.

## Getting Started

### Quick Start with Docker

To run `autodock-be` without cloning the repository, you can pull and run the Docker image directly:

```bash
docker run -d \
  --name autodock-be-container \
  -p 8001:8888 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v /home/admin:/app/docker-compose \
  chamixth/autodock-be
