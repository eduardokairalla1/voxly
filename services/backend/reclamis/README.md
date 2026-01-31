# Reclamis

This is Voxly's REST API service.

---

## Getting Started

Follow these steps to set up and run the project on your local machine.

### 1. Prerequisites

- Docker and Docker Compose

> [Docker install documentation](https://docs.docker.com/engine/install/)

> **Recommended:** Use the Docker development environment for the best experience with all tools pre-configured.

### 2. Development with Docker Compose (Recommended)

**Start the development environment:**

```bash
cd ../../infra/docker
docker compose -f docker-compose.yaml -f docker-compose-dev.yaml up -d
```

**Access the container via SSH:**

```bash
ssh root@localhost -p 10001
```

**Inside the container, navigate to the service and start development:**

```bash
cd /root/reclamis

# Install dependencies (first time or after go.mod changes)
go mod download

# Start development server
./scripts/dev  # Starts Air with hot reload
```

The API will be running on `http://localhost:8080`.

### 3. Local Development (Alternative)

If you prefer to develop without Docker:

**Prerequisites:**
- [Go 1.25+](https://go.dev/)
- [Air](https://github.com/air-verse/air)
- [golangci-lint](https://golangci-lint.run/)
- [gci](https://github.com/daixiang0/gci)

**Setup:**

```bash
# Copy environment file
cp .env.example .env

# Install dependencies
go mod download

# Run development server
./scripts/dev
```

> **Important:** Remember to add any new environment variables to both `.env` and the configuration in `internal/config`.

---

## Project Structure

```
reclamis/
├── cmd/reclamis/         # Application entry point
├── internal/             # Private application code
│   ├── config/           # Configuration management
│   ├── handlers/         # HTTP request handlers
│   ├── models/           # Data models
│   └── router/           # Route definitions
├── pkg/                  # Public libraries
│   └── errors/           # Custom error types
├── docker/               # Docker configuration
├── scripts/              # Build and development scripts
└── tests/                # Test files
```

## Running Tests

```bash
./scripts/test
```

The test script will:
- Validate import ordering with `gci`
- Run linters with `golangci-lint`
- Execute unit tests with race detection

> **Before submitting your PR:** Always run the test suite to ensure code quality.

---

## Deploying with Docker

This project is designed to be deployed as a Docker container.

### Building the Image

To build the Docker image, run the `build` script:

```bash
./scripts/build
```

Or with your Docker Hub username:

```bash
DOCKER_USER=your-docker-hub-user ./scripts/build
```

This will build and tag the image (e.g., `your-docker-hub-user/reclamis:0.0.1`).

---

## System API Endpoints

- `GET /system/health` - Health check endpoint
- `GET /system/info` - Service information

### Health Check

You can query the health endpoint using `curl`:

```bash
curl -fsS http://localhost:8080/system/health | jq .
```


### Service Info

```bash
curl -fsS http://localhost:8080/system/info | jq .
```
