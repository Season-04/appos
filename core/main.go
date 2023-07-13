package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/docker/docker/client"
	"github.com/staugaard/app-os/core/internal/config"
	"github.com/staugaard/app-os/core/internal/pb"
	"github.com/staugaard/app-os/core/internal/users"
	"google.golang.org/grpc"
)

func main() {
	configDirectory := os.Getenv("APP_OS_CONFIG_DIRECTORY")
	if configDirectory == "" {
		configDirectory = "/config"
	}

	cfg := config.NewConfig(configDirectory)
	err := cfg.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatalf("failed to connect to docker: %v", err)
	}

	err = cfg.Run(context.Background(), dockerClient)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUsersServiceServer(s, &users.Server{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening at %s", lis.Addr().String())

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}