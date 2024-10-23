package main

import (
	"context"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pbChat "my-messenger/chats/api"
	pbProfile "my-messenger/profile/api"
	"net/http"
)

type server struct {
	pbChat.UnimplementedChatServiceServer
	pbProfile.UnimplementedProfileServiceServer
}

func main() {
	grpcServer := grpc.NewServer()
	pbChat.RegisterChatServiceServer(grpcServer, &server{})
	pbProfile.RegisterProfileServiceServer(grpcServer, &server{})

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	mux := runtime.NewServeMux()
	err = pbChat.RegisterChatServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("Failed to register ChatService handler: %v", err)
	}

	err = pbProfile.RegisterProfileServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("Failed to register ProfileService handler: %v", err)
	}

	// Запуск HTTP-сервера для gRPC-Gateway
	http.Handle("/", mux)
	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
