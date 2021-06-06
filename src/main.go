package main

import (
	"log"
	"net"
	"os"
	"userservice/server"
	"userservice/user"
	"userservice/usergrpc/userservice"

	"google.golang.org/grpc"
)

func main() {
	config := getConfiguration()

	if !config.IsValidConfiguration() {
		log.Fatalf("No valid database configuration present")
	}

	repo, err := user.NewRepository(config.username, config.password, config.hostname, config.port, config.database)

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

func getConfiguration() databaseConfiguration {
	return databaseConfiguration{
		password: os.Getenv("DATABASE_PASSWORD"),
		username: os.Getenv("DATABASE_USERNAME"),
		port:     os.Getenv("DATABASE_PORT"),
		hostname: os.Getenv("DATABASE_HOSTNAME"),
		database: os.Getenv("DATABASE_DATABASE"),
	}
}

type databaseConfiguration struct {
	password string
	username string
	port     string
	hostname string
	database string
}

func (c *databaseConfiguration) IsValidConfiguration() bool {
	if len(c.password) > 0 && len(c.username) > 0 {
		return true
	}
	return false
}
