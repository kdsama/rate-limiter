package repository

import "github.com/kdsama/rate-limiter/entity"

type UserRepo interface {
	SaveUser(*entity.User) error
	GetUserByEmail(string) (*entity.User, error)
	GetUserByID(string) (*entity.User, error)
	CountUsersFromIDs([]string) (int64, error)
	GetUserNamesByIDs([]string) ([]string, error)
}
