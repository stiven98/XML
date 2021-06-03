package repository

import (
	"admin-service/model"
	"fmt"
	"gorm.io/gorm"
)

type AdministratorsRepository struct {
	Database *gorm.DB
}

func (repo *AdministratorsRepository) Update(administrator *model.Administrator) error {
	result := repo.Database.Model(model.Administrator{}).Where("id = ?", administrator.ID).Updates(administrator)
	return result.Error
}
func(repo *AdministratorsRepository) GetAll() []model.Administrator{
	var admins []model.Administrator
	repo.Database.Find(&admins)
	return admins
}

func (repo *AdministratorsRepository) Create(admin *model.Administrator) error  {
	result := repo.Database.Create(admin)
	fmt.Println(result.RowsAffected)
	return nil
}
