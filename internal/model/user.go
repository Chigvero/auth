package model

import (
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	UserType  desc.Role
	CreatedAt *timestamp.Timestamp
	UpdatedAt *timestamp.Timestamp
}

type UserInfo struct {
}
