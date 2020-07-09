package dao

import (
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (r *UserRepository) Create(user model.User) model.User {
	r.DB.Create(&user)
	return user
}

func (r *UserRepository) FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.Where(&model.User{Username: username}).Find(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
