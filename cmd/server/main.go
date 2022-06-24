package main

import (
	"log"

	"github.com/safciplak/go-grpc-microservices/internal/db"
	"github.com/safciplak/go-grpc-microservices/internal/rocket"
	"github.com/safciplak/go-grpc-microservices/internal/transport/grpc"
)

func Run() error {
	// responsbile for initializing and starting
	// our gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}

	rktService := rocket.New(rocketStore)
	rktHandler := grpc.New(rktService)

	if err := rktHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
