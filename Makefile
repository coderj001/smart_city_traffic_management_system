# Makefile for Smart City Traffic Management System

# Project-specific settings
PROJECT_NAME := smart_city_traffic_system
DOCKER_COMPOSE_FILE := docker-compose.yaml

# Docker and Docker Compose commands
DOCKER := docker
DOCKER_COMPOSE := docker-compose -f $(DOCKER_COMPOSE_FILE)

# Default action
all: build

# Build the Docker image
build:
	$(DOCKER) build -t $(PROJECT_NAME) .

# Start the services defined in the Docker Compose file
up:
	$(DOCKER_COMPOSE) up -d

# Stop the services defined in the Docker Compose file
down:
	$(DOCKER_COMPOSE) down

# Run tests
test:
	$(DOCKER) run --rm $(PROJECT_NAME) go test ./...

# Clean up Docker images and containers
clean:
	$(DOCKER) rmi $(PROJECT_NAME)
	$(DOCKER_COMPOSE) down --volumes --remove-orphans

# Display logs from the containers
logs:
	$(DOCKER_COMPOSE) logs

# Run the application in development mode
dev:
	$(DOCKER_COMPOSE) up web

# Interactive shell inside the web container
shell:
	$(DOCKER_COMPOSE) exec web /bin/sh

.PHONY: all build up down test clean logs dev shell
