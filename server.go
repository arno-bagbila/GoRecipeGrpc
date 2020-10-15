package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/go-recipe-grpc/recipe"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	s := recipe.Server{}

	grpcServer := grpc.NewServer()

	recipe.RegisterRecipeServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server on port 9000: %s", err)
	}
}
