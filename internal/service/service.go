package service

import (
	"context"
	"github.com/Chigvero/auth/internal/entities"
	"github.com/Chigvero/auth/internal/repository"
	"github.com/Chigvero/auth/internal/service/user_v1"
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserService interface {
	Create(ctx context.Context, request *entities.CreateRequest) (int64, error) //returning id
	Get(ctx context.Context, id int64) (*desc.GetResponse, error)
	Update(ctx context.Context, request *entities.UpdateRequest) (*empty.Empty, error)
	Delete(ctx context.Context, id int64) (*empty.Empty, error)
}

type Service struct {
	UserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService: user_v1.NewUserService(repository.UserRepository),
	}
}
