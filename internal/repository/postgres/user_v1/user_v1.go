package user_v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/Chigvero/auth/internal/converter"
	"github.com/Chigvero/auth/internal/entities"
	desc "github.com/Chigvero/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v5"
)

const (
	usersTable = "users"
)

type UserPostgres struct {
	db *pgx.Conn
}

func NewUserPostgres(db *pgx.Conn) *UserPostgres {
	return &UserPostgres{db: db}
}

func (ur *UserPostgres) Create(ctx context.Context, user *entities.CreateRequest) (int64, error) {
	userQuerystr := fmt.Sprintf("INSERT INTO %s (name,email ,password,password_confirm,user_type) VALUES($1,$2,$3,$4,$5) RETURNING id", usersTable)
	var id int64
	err := ur.db.QueryRow(ctx, userQuerystr, user.Name, user.Email, user.Password, user.PasswordConfirm, user.UserType).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ur *UserPostgres) Get(ctx context.Context, id int64) (*desc.GetResponse, error) {
	selectQuery := fmt.Sprintf("SELECT id,name,email,user_type,created_at,updated_at from %s WHERE id =$1", usersTable)
	var u entities.GetResponse
	var user_type string
	err := ur.db.QueryRow(ctx, selectQuery, id).Scan(&u.Id, &u.Name, &u.Email, &user_type, &u.Created_at, &u.Updated_at)
	if err != nil {
		return nil, err
	}
	switch user_type {
	case "user":
		u.Role = 1
	case "admin":
		u.Role = 2
	}
	fmt.Println(u)
	res := converter.ToDescGetUser(&u)
	return res, nil
}

func (ur *UserPostgres) Update(ctx context.Context, r *entities.UpdateRequest) (*empty.Empty, error) {
	tx, err := ur.db.Begin(ctx)
	if err != nil {
		return &empty.Empty{}, err
	}
	if r.Name != nil {
		updateQuery := fmt.Sprintf("UPDATE %s SET name=$1 where id=$2", usersTable)
		res, err := tx.Exec(ctx, updateQuery, r.Name, r.Id)
		if err != nil {
			tx.Rollback(ctx)
			return &empty.Empty{}, err
		}
		if res.RowsAffected() == 0 {
			return &empty.Empty{}, errors.New("NOT found user with same id")
		}
	}
	if r.Email != nil {
		updateQuery := fmt.Sprintf("UPDATE %s SET email=$1 where id=$2", usersTable)
		res, err := tx.Exec(ctx, updateQuery, r.Email, r.Id)
		if err != nil {
			tx.Rollback(ctx)
			return &empty.Empty{}, err
		}
		if res.RowsAffected() == 0 {
			return &empty.Empty{}, errors.New("NOT found user with same id")
		}
	}
	tx.Commit(ctx)
	return &empty.Empty{}, nil
}

func (ur *UserPostgres) Delete(ctx context.Context, id int64) (*empty.Empty, error) {
	deleteQuery := fmt.Sprintf("DELETE FROM %s where id=$1", usersTable)
	res, err := ur.db.Exec(ctx, deleteQuery, id)
	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, errors.New("Not found user with this id")
	}
	return &empty.Empty{}, nil
}
