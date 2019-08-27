package repository

import (
	"bookstore/model"
	"context"
)

type UserRepository interface {
	CheckLogin(context context.Context, loginReq model.LoginRequest) (model.User, error)
	Save(context context.Context, user model.User) (model.User, error)
	SelectById(context context.Context, userId string) (model.User, error)
	SelectAll(context context.Context, userId string) ([]model.User, error)
}
