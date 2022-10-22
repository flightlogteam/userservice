package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/flightlogteam/userservice/grpc/userservice"
	"github.com/flightlogteam/userservice/src/device"
	"github.com/flightlogteam/userservice/src/presentation"
	"github.com/flightlogteam/userservice/src/server"
	"github.com/flightlogteam/userservice/src/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

func main() {
	rcpPort := ":61226"
	restPort := ":61228"
	config := getConfiguration()

	if !config.IsValidConfiguration() {
		log.Fatalf("No valid database configuration present \n")
	}

	dbContext, err := connectToDatabase(createConnectionString(config.username, config.password, config.hostname, config.port, config.database), 5, 10)

	err = dbContext.AutoMigrate(&user.User{}, &device.FlyingDevice{})

	if err != nil {
		log.Fatalf("Cannot migrate the database: %v, on hostname %v, with user %v. Config gave following error:\n %v \n", config.database, config.hostname, config.username, err)
		return
	}

	repo := user.NewRepository(dbContext)

	log.Printf("Connected to database %v, with user %v \n", config.database, config.username)

	listener, err := net.Listen("tcp", rcpPort)

	if err != nil {
		log.Fatalf("Unable to listen to spcified port: %v \n", err)
	}

	userService := user.NewUserService(repo)

	deviceRepo := device.NewRepository(dbContext)
	deviceService := device.NewDeviceService(deviceRepo)

	userserver := server.NewGrpcServer(userService, deviceService)
	grpcServer := grpc.NewServer()
	userservice.RegisterUserServiceServer(grpcServer, &userserver)

	api := presentation.NewApi(userService, deviceService)

	go func() {
		api.Start(restPort)
	}()

	log.Printf("starting server on port %v...", rcpPort)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}

func connectToDatabase(connectionString string, retryCount int, sleepTime int) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	attempt := 1

	for err != nil && attempt <= retryCount {
		log.Printf("Could not connect to the database, will wait for %v and this was attempt %v of %v", sleepTime, attempt, retryCount)
		attempt += 1
		time.Sleep(time.Duration(sleepTime) * time.Second)
		db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	}

	return db, err
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

func createConnectionString(username string, password string, hostname string, port string, database string) string {
	connectionString := fmt.Sprintf("%v:%v@", username, password)

	if len(hostname) > 0 {
		connectionString += fmt.Sprintf("tcp(%v:%v)", hostname, port)
	}

	if len(database) > 0 {
		connectionString += fmt.Sprintf("/%v", database)
	} else {
		connectionString += "/"
	}

	return connectionString + "?parseTime=true"

}

func ensureDatabaseExists(username string, password string, hostname string, port string, database string) error {
	connectionWithoutDatabase, err := gorm.Open(mysql.Open(createConnectionString(username, password, hostname, port, "")), &gorm.Config{})

	if err != nil {
		return err
	}

	tx := connectionWithoutDatabase.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", database))
	tx.Commit()

	return nil
}
