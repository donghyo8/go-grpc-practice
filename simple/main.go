package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = "9000"

func main(){
	listen, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("start grpc server on %s port", port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}