package repository

import (
	"gorm.io/gorm"
	"shop-backend/model"
	"shop-backend/model/dto"
)

type UsersRepository struct {
	Database *gorm.DB
}

func (r UsersRepository) Create(user model.User) error {
	result := r.Database.Create(user)
	return result.Error
}

func (r UsersRepository) Login(loginInfo dto.LoginInfoDTO) (model.User, int64)  {
	var user model.User
	var count int64
	r.Database.Model(model.User{}).Find(&user, "username = ?", loginInfo.Username, "password = ?", loginInfo.UserID).Count(&count)

	return user, count
}


