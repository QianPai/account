package impl

import (
	"context"
	"github.com/QianPai/account/model"
	"github.com/QianPai/account/repository"
	"time"
)

type userUsecase struct {
	userRepo repository.User
	contextTimeout time.Duration
}

func NewUserUsecase(repo repository.User, timeout time.Duration) *userUsecase {
	return &userUsecase{
		userRepo: repo,
		contextTimeout: timeout,
	}
}

func (u userUsecase) GetById(c context.Context, id interface{}) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res,err := u.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (u userUsecase) GetByName(c context.Context, name string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res,err := u.userRepo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (u userUsecase) GetByPhone(c context.Context, phone string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res,err := u.userRepo.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (u userUsecase) Store(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err := u.userRepo.Store(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u userUsecase) Update(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err := u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
