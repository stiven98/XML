package repository

import (
	"fmt"
	"gorm.io/gorm"
	"profileservice/model"
)


type AdministratorsRepository struct {
	Database *gorm.DB
}

func (repo *AdministratorsRepository) Update(administrator *model.Administrator) error {
	result := repo.Database.Model(model.Administrator{}).Where("user_id = ?", administrator.UserID).Updates(administrator)
	return result.Error
}
func(repo *AdministratorsRepository) GetAll() []model.Administrator{
	var admins []model.Administrator
	repo.Database.Preload("SystemUser").Find(&admins)
	return admins
}

func (repo *AdministratorsRepository) Create(admin *model.Administrator) error {
	result := repo.Database.Create(admin)
	fmt.Println(result.RowsAffected)
	return nil
}
