package repository

import (
	"fmt"
	"github.com/google/uuid"
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
	repo.Database.Where("type_of_user = 'agent'").Or("type_of_user = 'user'").Find(&users)
	var usernames []string
	for _, user := range users {
		usernames = append(usernames, user.Username);
	}
	return usernames
}
func(repo *SystemUsersRepository) GetUserId(username string) uuid.UUID{
	var users []model.SystemUser
	repo.Database.Find(&users)
	var id uuid.UUID
	for _, user := range users {
		if user.Username == username {
			id = user.ID
		}
	}
	return id;
}

func(repo *SystemUsersRepository) GetById(id uuid.UUID) model.SystemUser{
	var user model.SystemUser
	repo.Database.Find(&user).Where("id = ?", id);
	return user
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

func (repo *SystemUsersRepository) UpdateVerification(id uuid.UUID) interface{} {
	result := repo.Database.Model(model.User{}).Where("user_id = ?", id).Update("is_verified", true)
	return result
}


