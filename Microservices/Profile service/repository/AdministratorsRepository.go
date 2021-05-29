package repository

import (
	"gorm.io/gorm"
	"profileservice/model"
)

type AdministratorRepository struct {
	Database *gorm.DB
}

func (repo *AdministratorRepository) Update(administrator *model.Administrator) error {
	result := repo.Database.Model(model.Administrator{}).Where("user_id = ?", administrator.UserID).Updates(administrator)
	return result.Error
}
