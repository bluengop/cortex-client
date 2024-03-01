SHELL := /bin/bash

.PHONY: run build docker

run:
	@go run ./cmd/main/main.go

build:
	@go build -o bin/myapp ./cmd/main/main.go

docker:
	@echo "TODO:"
	@echo "'docker build -t myapp-image .'"

