package repository

import (
	"fmt"
	"gorm.io/gorm"
	"profileservice/model"
)

type UsersRepository struct {
	Database *gorm.DB
}

func (repo *UsersRepository) Update(user *model.User) error {
	result := repo.Database.Model(model.User{}).Where("user_id = ?", user.UserID).Updates(user)
	return result.Error
}
func(repo *UsersRepository) GetAll() []model.User{
	var users []model.User
	repo.Database.Preload("SystemUser").Find(&users)
	return users
}

func (repo *UsersRepository) Create(user *model.User) error {
	result := repo.Database.Create(user)
	fmt.Println(result.RowsAffected)
	return nil
}
