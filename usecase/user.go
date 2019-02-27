package usecase

import (
	"context"
	"github.com/QianPai/account/model"
)

type User interface {
	//GetById(ctx context.Context, id interface{}) (res *model.User, err error)
	GetByPhone(ctx context.Context, phone string) (res *model.User, err error)
	GetByName(ctx context.Context, name string) (res *model.User, err error)
	//Update(ctx context.Context, user *model.User) error
	Store(ctx context.Context, user *model.User) error
	//Delete(ctx context.Context, id interface{}) error
}
