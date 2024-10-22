package converter

import (
	"github.com/Chigvero/auth/internal/entities"
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateUser(r *desc.CreateRequest) *entities.CreateRequest {
	return &entities.CreateRequest{
		Name:            r.GetName(),
		Email:           r.GetEmail(),
		Password:        r.GetPassword(),
		PasswordConfirm: r.GetPasswordConfirm(),
		UserType:        r.GetUserType(),
	}
}

func ToDescGetUser(r *entities.GetResponse) *desc.GetResponse {
	if r.Updated_at.Valid {
		return &desc.GetResponse{
			UserType:  r.Role,
			Id:        r.Id,
			CreatedAt: timestamppb.New(r.Created_at),
			UpdatedAt: timestamppb.New(r.Updated_at.Time),
			Email:     r.Email,
			Name:      r.Name,
		}
	}
	return &desc.GetResponse{
		UserType:  r.Role,
		Id:        r.Id,
		CreatedAt: timestamppb.New(r.Created_at),
		UpdatedAt: nil,
		Email:     r.Email,
		Name:      r.Name,
	}
}

func ToUpdateUser(usr *desc.UpdateRequest) *entities.UpdateRequest {
	return &entities.UpdateRequest{
		Email: usr.GetEmail(),
		Name:  usr.GetName(),
		Id:    usr.GetId(),
	}
}
