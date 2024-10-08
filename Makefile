.PHONY: help rs rc

.DEFAULT_GOAL := help

help:
	@echo "Available commands:"
	@echo "  make rs     - Start the server"
	@echo "  make rc     - Start the client"

rs:
	go run ./cmd/server/main.go

rc:
	go run ./cmd/client/main.go