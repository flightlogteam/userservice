package main

import (
	"log"
	"net"
	"userservice/server"
	"userservice/user"
	"userservice/usergrpc/userservice"

	"google.golang.org/grpc"
)

func main() {
	repo, err := user.NewRepository("root", "pass", "localhost", "3306", "flightloguser")

	if err != nil {
		log.Fatalf("cannot create connection %v", err)
		return
	}

	listener, err := net.Listen("tcp", ":61226")

	if err != nil {
		log.Fatalf("Unable to listen to spcified port: %v", err)
	}

	userService := user.NewUserService(repo)

	userserver := server.NewGrpcServer(userService)
	grpcServer := grpc.NewServer()
	userservice.RegisterUserServiceServer(grpcServer, &userserver)

	log.Println("starting server...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}
