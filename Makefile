.DEFAULT_GOAL := help

.PHONY: help run dev test fmt vet check docker-up docker-down

help:
	@echo "Available commands:"
	@echo "  make dev          Start Docker infrastructure and run the app"
	@echo "  make run          Run the Go server only"
	@echo "  make docker-up    Start Docker infrastructure"
	@echo "  make docker-down  Stop Docker infrastructure"
	@echo "  make test         Run tests"
	@echo "  make fmt          Format Go code"
	@echo "  make vet          Run go vet"
	@echo "  make check        Run fmt, vet and tests"

run:
	go run ./cmd/server

dev: docker-up run

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

check: fmt vet test

docker-up:
	docker compose up -d

docker-down:
	docker compose down
