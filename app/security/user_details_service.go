package security

import (
	"github.com/denisqq/xsolla-test/app/dao"
	"github.com/denisqq/xsolla-test/app/model"
)

type UserDetailsService interface {
	LoadUserDetails(username string)
}

type BasicAuthUserDetailsService struct {
	UserRepository dao.UserRepository
}

func NewBasicAuthUserDetailsService(userRepository dao.UserRepository) BasicAuthUserDetailsService {
	return BasicAuthUserDetailsService{UserRepository: userRepository}
}

func (b BasicAuthUserDetailsService) LoadUserDetails(username string) (model.User, error) {
	return b.UserRepository.FindUserByUsername(username)
}
