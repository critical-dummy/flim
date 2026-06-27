.PHONY: help proto build run test clean

help:
	@echo "FLIM - Flash Live Instant Messenger"
	@echo "Available commands:"
	@echo "  make proto   - Generate protobuf files"
	@echo "  make build   - Build the server"
	@echo "  make run     - Run the server"
	@echo "  make test    - Run tests"
	@echo "  make clean   - Clean build artifacts"

proto:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

build: proto
	go build -o bin/flim cmd/server/main.go

run: proto
	go run cmd/server/main.go

test:
	go test ./...

clean:
	rm -rf bin/
	find . -name "*.pb.go" -delete
