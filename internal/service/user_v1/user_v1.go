package user_v1

import (
	"context"
	"github.com/Chigvero/auth/internal/entities"
	"github.com/Chigvero/auth/internal/repository"
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

func (s *UserService) Create(ctx context.Context, request *entities.CreateRequest) (int64, error) {
	return s.repo.Create(ctx, request)
}

func (s *UserService) Get(ctx context.Context, id int64) (*desc.GetResponse, error) {
	return s.repo.Get(ctx, id)
}
func (s *UserService) Update(ctx context.Context, request *entities.UpdateRequest) (*empty.Empty, error) {
	return s.repo.Update(ctx, request)

}
func (s *UserService) Delete(ctx context.Context, id int64) (*empty.Empty, error) {
	return s.repo.Delete(ctx, id)
}
