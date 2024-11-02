package repository

import (
	"context"
	"github.com/Chigvero/auth/internal/entities"
	"github.com/Chigvero/auth/internal/repository/postgres/user_v1"
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	Create(ctx context.Context, request *entities.CreateRequest) (int64, error)
	Update(ctx context.Context, request *entities.UpdateRequest) (*empty.Empty, error)
	Get(ctx context.Context, id int64) (*desc.GetResponse, error)
	Delete(ctx context.Context, id int64) (*empty.Empty, error)
}

type Repository struct {
	UserRepository
}

func NewRepository(pg *pgx.Conn) *Repository {
	return &Repository{
		UserRepository: user_v1.NewUserPostgres(pg),
	}
}
