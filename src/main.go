package main

import (
	"log"
	"net"
	"os"

	"github.com/flightlogteam/userservice/grpc/userservice"
	"github.com/flightlogteam/userservice/src/server"
	"github.com/flightlogteam/userservice/src/user"

	"google.golang.org/grpc"
)

func main() {
	rcpPort := ":61226"
	config := getConfiguration()

	if !config.IsValidConfiguration() {
		log.Fatalf("No valid database configuration present \n")
	}

	repo, err := user.NewRepository(config.username, config.password, config.hostname, config.port, config.database)

	if err != nil {
		log.Fatalf("cannot create connection to database: %v, on hostname %v, with user %v. Config gave following error:\n %v \n", config.database, config.hostname, config.username, err)
		return
	}

	log.Printf("Connected to database %v, with user %v \n", config.database, config.username)


	listener, err := net.Listen("tcp", rcpPort)

	if err != nil {
		log.Fatalf("Unable to listen to spcified port: %v \n", err)
	}

	userService := user.NewUserService(repo)

	userserver := server.NewGrpcServer(userService)
	grpcServer := grpc.NewServer()
	userservice.RegisterUserServiceServer(grpcServer, &userserver)

	log.Printf("starting server on port %v...", rcpPort)

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
