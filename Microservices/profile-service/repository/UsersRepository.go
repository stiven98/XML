package repository

import (
	"fmt"
	"gorm.io/gorm"
	"profileservice/model"
	"profileservice/model/Dto"
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
	return result.Error
}
func (repo *UsersRepository) ChangeWhetherIsPublic(dto *Dto.ChangeWhetherIsPublicDto) error {
	result := repo.Database.Model(model.User{}).Where("user_id = ?", dto.USERID).UpdateColumn("is_public", dto.FLAG)
	return result.Error
}

func (repo *UsersRepository) ChangeAllowedTags(dto *Dto.ChangeAllowedTagsDto) error {
	result := repo.Database.Model(model.User{}).Where("user_id = ?", dto.USERID).UpdateColumn("allowed_tags", dto.FLAG)
	return result.Error
}