package usecase

import (
	"context"
	"github.com/QianPai/account/model"
	"github.com/QianPai/account/repository"
	"time"
)

type User interface {
	Login(ctx context.Context) (res model.User)
	Register(ctx context.Context) (res model.User)
	ResetPassword(ctx context.Context) (res model.User)
	ResetToken(ctx context.Context) (res model.User)
	GetSmsCheckCode(ctx context.Context) string
}

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