package services

import (
	"errors"

	"github.com/kdsama/rate-limiter/entity"
	"github.com/kdsama/rate-limiter/repository"
	"github.com/kdsama/rate-limiter/utils"
)

var (
	Err_Invalid_Hash            = errors.New("hash provided is invalid")
	Err_User_Present            = errors.New("user already present")
	Err_IncorrectUserOrPassword = errors.New("user or password provided was incorrect")
)

type UserServicer interface {
	Get()
	Save()         // saves the user
	IsAuthorized() // is authorised
	GetLimitByUserID(userid string) (int, error)
}
type UserService struct {
	repo  repository.UserRepo
	token UserTokenService
}

func NewUserService(userrepo repository.UserRepo, tokenservice *UserTokenService) *UserService {

	return &UserService{
		repo:  userrepo,
		token: *tokenservice,
	}
}

func (us *UserService) LoginUser(email string, password string) (string, error) {

	userobject, err := us.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = utils.ComparePassword(userobject.Password, password)
	if err != nil {
		return "", Err_IncorrectUserOrPassword
	}
	return us.token.GenerateAndSaveUserToken(userobject.ID)

}

func (us *UserService) SaveUser(email string, name string, password string) (string, error) {

	_, err := us.repo.GetUserByEmail(email)

	if err == nil {
		return "", Err_User_Present
	}

	if err != nil {
		return "", err
	}

	encryptedPassword, err := utils.GenerateHashForPassword(password)
	if err != nil {
		return "", err
	}
	userObject := entity.NewUser(email, encryptedPassword)
	err = us.repo.SaveUser(userObject)
	if err != nil {
		return "", err
	}
	return us.token.GenerateAndSaveUserToken(userObject.ID)
}

func (us *UserService) GetUserByID(id string) (*entity.User, error) {
	user, err := us.repo.GetUserByID(id)

	return user, err
}
func (us *UserService) GetUserNamesByIDs(ids []string) ([]string, error) {
	user, err := us.repo.GetUserNamesByIDs(ids)
	return user, err
}
func (us *UserService) CountUsersFromIDs(user_ids []string) (int64, error) {
	user, err := us.repo.CountUsersFromIDs(user_ids)

	return user, err
}
