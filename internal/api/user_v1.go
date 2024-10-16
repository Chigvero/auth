package api

import (
	"context"
	"github.com/Chigvero/auth/internal/converter"
	"github.com/Chigvero/auth/internal/service"
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	service *service.Service
}

func NewImplementation(service *service.Service) *Implementation {
	return &Implementation{
		UnimplementedUserV1Server: desc.UnimplementedUserV1Server{},
		service:                   service,
	}
}

func (i *Implementation) Create(ctx context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {
	user := converter.ToCreateUser(r)
	log.Println(*user)
	id, err := i.service.UserService.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (i *Implementation) Get(ctx context.Context, r *desc.GetRequest) (*desc.GetResponse, error) {
	log.Println("id:", r.GetId())
	res, err := i.service.UserService.Get(ctx, r.GetId())
	if err != nil {
		return nil, err
	}
	return res, nil
	//return &desc.GetResponse{
	//	Id:        gofakeit.Int64(),
	//	Name:      gofakeit.BeerName(),
	//	Email:     gofakeit.Email(),
	//	UserType:  desc.Role_admin,
	//	CreatedAt: timestamppb.Now(),
	//	UpdatedAt: timestamppb.Now(),
	//}, nil
}

func (i *Implementation) Update(ctx context.Context, r *desc.UpdateRequest) (*empty.Empty, error) {
	usr := converter.ToUpdateUser(r)
	return i.service.UserService.Update(ctx, usr)

}

func (i *Implementation) Delete(ctx context.Context, r *desc.DeleteRequest) (*empty.Empty, error) {
	id := r.GetId()
	return i.service.UserService.Delete(ctx, id)
}
