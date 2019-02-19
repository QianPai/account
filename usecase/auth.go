package usecase

import (
	"context"
	"github.com/QianPai/account/model"
)

type auth interface {
	Login(ctx context.Context) (res model.User)
	Register(ctx context.Context) (res model.User)
	ResetPassword(ctx context.Context) (res model.User)
	ResetToken(ctx context.Context) (res model.User)
	GetSmsCheckCode(ctx context.Context) string
}
