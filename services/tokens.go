package services

import (
	"errors"
	"time"

	"github.com/kdsama/rate-limiter/entity"
	"github.com/kdsama/rate-limiter/repository"
	"github.com/kdsama/rate-limiter/utils"
)

var (
	Err_MaliciousIntent = errors.New("this token doesnot belong to the user but was saved here. ")
	Err_TokenExpired    = errors.New("token was expired")
)

type UserTokenService struct {
	UserTokenRepo repository.UserTokenRepo
}

func NewUserTokenService(repo repository.UserTokenRepo) *UserTokenService {
	return &UserTokenService{repo}
}

func (uts *UserTokenService) GenerateUserToken(userid string) (string, error) {
	return utils.CreateJWTToken(userid)
}

// generate saves and returns the token
func (uts *UserTokenService) GenerateAndSaveUserToken(userid string) (string, error) {

	timestamp := time.Now().Unix()
	// Generate JWT Token
	token, err := uts.GenerateUserToken(userid)
	if err != nil {
		return "", utils.Err_InvalidToken
	}
	userTokenObject := entity.NewUserToken(userid, token, timestamp)
	err = uts.UserTokenRepo.SaveUserToken(userTokenObject)
	if err != nil {
		return "", err
	}
	return userTokenObject.Token, nil
}

func (uts *UserTokenService) GetUserTokenByID(userid string) (*entity.UserToken, error) {
	return uts.UserTokenRepo.GetUserTokenByID(userid)
}
func (uts *UserTokenService) GetUserByToken(token string) (*entity.UserToken, error) {
	return uts.UserTokenRepo.GetUserByToken(token)
}

func (uts *UserTokenService) ValidateUserTokenAndGetUserID(token string) (string, error) {

	id, err := utils.VerifyJWTToken(token)
	if err != nil {
		return "", err
	}
	usertoken, err := uts.GetUserByToken(token)
	if err != nil {
		return "", err
	}
	if id != usertoken.User_ID {
		return "", Err_MaliciousIntent
	}
	if time.Now().Unix()-usertoken.UpdatedAt >= utils.OneHourInSeconds() {
		return "", Err_TokenExpired
	}
	return id, nil
}
