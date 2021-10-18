package main

import (
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	service "github.com/mehranghajari/Menchastic/backend/pkg/service/v2"
	gamepb "github.com/mehranghajari/Menchastic/backend/pkg/api/v2"
)


func main() {
	fmt.Println("--- SERVER APP ---")
	lis, err := net.Listen("tcp", "localhost:5400")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	gamepb.RegisterMenchasticServiceServer(grpcServer, service.NewServer())
	grpcServer.Serve(lis)
}
