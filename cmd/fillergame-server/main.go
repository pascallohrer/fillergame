package main

import (
	"log"
	"net"

	"github.com/pascallohrer/fillergame/api"
	"github.com/pascallohrer/fillergame/server"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50123")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	fillerServer := &server.Server{}
	api.RegisterFillerGameServer(grpcServer, fillerServer)
	go fillerServer.PrintDebug()
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
