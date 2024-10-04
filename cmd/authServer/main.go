package main

import (
	"context"
	"fmt"
	desc "github.com/Chigvero/auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const (
	grpcHost = "localhost:"
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedAuthV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s%d", grpcHost, grpcPort))
	if err != nil {
		log.Fatalf("Error with listening port:%v\n", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

func (s *server) Create(ctx context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Println(desc.CreateRequest{
		Name:            r.GetName(),
		Email:           r.GetEmail(),
		Password:        r.GetPassword(),
		PasswordConfirm: r.GetPasswordConfirm(),
		UserType:        r.GetUserType(),
	})
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Get(ctx context.Context, r *desc.GetRequest) (*desc.GetResponse, error) {
	log.Println("id:", r.GetId())
	return &desc.GetResponse{
		Id:        gofakeit.Int64(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		UserType:  desc.Role_admin,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}, nil
}

func (s *server) Update(ctx context.Context, r *desc.UpdateRequest) (*empty.Empty, error) {
	log.Println(desc.UpdateRequest{
		Id:    r.GetId(),
		Name:  r.GetName(),
		Email: r.GetEmail(),
	})
	return &empty.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, r *desc.DeleteRequest) (*empty.Empty, error) {
	log.Println("id: ", r.GetId())
	return &empty.Empty{}, nil
}
