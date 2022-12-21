#!/usr/bin/env make

.PHONY: build
build:
	go build -o bin/mongo-client ./cmd/mongo-client

.PHONY: run
run: build
	./bin/mongo-client