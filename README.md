# Device Service

## Overview

Device Service is a microservice that provides management and tracking functionalities for IoT devices as part of the BerryTracer project. It uses gRPC for inter-service communication.

## Features

- Device registration, updates, and management.
- Real-time device tracking and status updates.
- Secure gRPC communication.
- Efficient MongoDB queries with dynamic indexing.

## Prerequisites

- Go (version 1.15 or later recommended)
- A MongoDB instance
- The Auth Service running and accessible (specified in the `AUTH_SERVICE_URL` environment variable)

## Installation

Clone the repository to your local machine:

```bash
git clone https://github.com/BerryTracer/device-service.git
cd device-service
```

```.env
MONGODB_URI=mongodb://root:password@localhost:27017/berrytracer
AUTH_SERVICE_URL=localhost:50052
```

## Running the Service

To run the service locally:

```bash
go run main.go
```

## Docker Compose

To start MongoDB using Docker Compose:

```bash
docker-compose -f docker-compose.dev.db.yml up -d
```

Ensure the MongoDB URI in the .env file matches the configuration in the Docker Compose file.

## Project Structure

- /grpc: gRPC service definitions and protocol buffers.
- /model: Data models for the service.
- /repository: Data access layer for database operations.
- /service: Business logic and service handlers.
- main.go: Entry point of the service.

## Development

Build the project with:

```bash
go build -o device-service
```

Run tests using:

```bash
go test ./...
```
