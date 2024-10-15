package main

import (
	"context"
	"fmt"
	"github.com/Chigvero/auth/internal/api"
	"github.com/Chigvero/auth/internal/repository"
	"github.com/Chigvero/auth/internal/service"
	"github.com/jackc/pgx/v5"
	"log"
	"net"

	desc "github.com/Chigvero/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcHost = "localhost:"
	grpcPort = 50051
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "host=localhost port=54331 user=user dbname=user-db password=user-password sslmode=disable")
	repos := repository.NewRepository(conn)
	services := service.NewService(repos)
	server := api.NewImplementation(services)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%d", grpcHost, grpcPort))
	if err != nil {
		log.Fatalf("Error with listening port:%v\n", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, server)
	log.Println("Server started on port:")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve:%v", err)
	}

}
