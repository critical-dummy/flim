package main

import (
	"context"
	"log"
	"net"

	"github.com/critical-dummy/flim/internal/auth"
	"github.com/critical-dummy/flim/internal/db"
	"github.com/critical-dummy/flim/internal/server"
	"google.golang.org/grpc"
)

func main() {
	// MongoDB connection
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoDB, err := db.NewMongoDBClient(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoDB.Disconnect(ctx)

	// Create gRPC server
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// Initialize JWT manager
	jwtManager := auth.NewJWTManager("secret-key-change-this", 24)

	// Create and register messenger service
	messengerServer := server.NewMessengerServer(mongoDB, jwtManager)
	// TODO: Register service with s.Register()

	log.Println("FLIM Server listening on :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
