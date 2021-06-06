package repository

import (
	"fmt"
	"gorm.io/gorm"
	"profileservice/model"
)

type SystemUsersRepository struct {
	Database *gorm.DB
}

func(repo *SystemUsersRepository) GetAll() []model.SystemUser{
	var users []model.SystemUser
	repo.Database.Find(&users)
	return users
}
func(repo *SystemUsersRepository) GetAllUsernames() []string{
	var users []model.SystemUser
	repo.Database.Find(&users)
	var usernames []string
	for _, user := range users {
		usernames = append(usernames, user.Username);
	}
	return usernames
}

func (repo *SystemUsersRepository) Create(user *model.SystemUser) error {
	result := repo.Database.Create(user)
	fmt.Println(result.RowsAffected)
	return nil
}
func (repo *SystemUsersRepository) Update(user *model.SystemUser) error {
	result := repo.Database.Model(model.SystemUser{}).Where("id = ?", user.ID).UpdateColumns(user)
	fmt.Println(result.RowsAffected)
	return nil
}


