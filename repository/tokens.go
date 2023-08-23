package repository

import "github.com/kdsama/rate-limiter/entity"

type UserTokenRepo interface {
	SaveUserToken(*entity.UserToken) error
	GetUserTokenByID(string) (*entity.UserToken, error)
	GetUserByToken(string) (*entity.UserToken, error)
}

type UserTokenRepository struct {
	UserTokenRepo
}

func NewUserTokenRepository(br UserTokenRepo) *UserTokenRepository {
	return &UserTokenRepository{br}
}
