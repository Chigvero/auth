package main

import (
	"context"
	"flag"
	"github.com/Chigvero/auth/internal/api"
	"github.com/Chigvero/auth/internal/config"
	"github.com/Chigvero/auth/internal/config/env"
	"github.com/Chigvero/auth/internal/repository"
	"github.com/Chigvero/auth/internal/service"
	"github.com/jackc/pgx/v5"
	"log"
	"net"

	desc "github.com/Chigvero/auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	ctx := context.Background()
	//Configs
	flag.Parse()
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("error with load config file:%v", err)
	}
	dsn, err := env.NewpgConfig()
	if err != nil {
		log.Fatalf("failed to get grpcConfig:%v", err)
	}
	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpcConfig:%v", err)
	}
	//Configs
	//DB
	conn, err := pgx.Connect(ctx, dsn.DSN())
	if err != nil {
		log.Fatalf("failed to connect DB:%v", err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("failed to ping:%v", err)
	}
	//DB
	//Слои
	repos := repository.NewRepository(conn)
	services := service.NewService(repos)
	server := api.NewImplementation(services)
	//Слои

	//
	lis, err := net.Listen("tcp", grpcConfig.Address())
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
	//
}
