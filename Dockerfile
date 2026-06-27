# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git protobuf-dev

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Generate protobuf files
RUN protoc --go_out=. --go-grpc_out=. proto/*.proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o flim cmd/server/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/flim .

EXPOSE 50051

CMD ["./flim"]
