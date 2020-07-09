package service

import (
	"errors"
	"github.com/denisqq/xsolla-test/app/dao"
	"github.com/denisqq/xsolla-test/app/model"
)

type UserService struct {
	UserRepository dao.UserRepository
}

func NewUserService(repository dao.UserRepository) UserService {
	return UserService{
		UserRepository: repository,
	}
}

func (us *UserService) CreateUser(user model.User) (model.User, error) {
	_, err := us.UserRepository.FindUserByUsername(user.Username)
	if err != nil {
		return us.UserRepository.Create(user), nil
	}

	return model.User{}, errors.New("user with this username exist")
}

func (us *UserService) FindByUsername(username string) (model.User, error) {
	return us.UserRepository.FindUserByUsername(username)
}
