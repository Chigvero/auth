package entities

import (
	"database/sql"
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/wrappers"
	"time"
)

type CreateRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	UserType        desc.Role
}
type GetResponse struct {
	Id         int64
	Name       string
	Email      string
	Role       desc.Role
	Created_at time.Time
	Updated_at sql.NullTime
}

type UpdateRequest struct {
	Id    int64
	Email *wrappers.StringValue
	Name  *wrappers.StringValue
}
