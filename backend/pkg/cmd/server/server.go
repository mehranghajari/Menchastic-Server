package main

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	service "github.com/mehranghajari/Menchastic/backend/pkg/service/v1"
	chatpb "github.com/mehranghajari/Menchastic/backend/pkg/api/v1"
)


func main() {
	fmt.Println("--- SERVER APP ---")
	lis, err := net.Listen("tcp", "localhost:5400")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	chatpb.RegisterChatServiceServer(grpcServer, service.NewServer())
	grpcServer.Serve(lis)
}
