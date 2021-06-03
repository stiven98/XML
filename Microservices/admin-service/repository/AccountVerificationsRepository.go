package repository

import (
	"admin-service/model"
	"fmt"
	"gorm.io/gorm"
)

type AccountVerificationsRepository struct {
	Database *gorm.DB
}

func (repo *AccountVerificationsRepository) Update(request *model.AccountVerificationRequest) error {
	result := repo.Database.Model(model.AccountVerificationRequest{}).Where("id = ?", request.ID).Updates(request)
	return result.Error
}
func(repo *AccountVerificationsRepository) GetAll() []model.AccountVerificationRequest{
	var requests []model.AccountVerificationRequest
	repo.Database.Find(&requests)
	return requests
}

func (repo *AccountVerificationsRepository) Create(request *model.AccountVerificationRequest) error  {
	result := repo.Database.Create(request)
	fmt.Println(result.RowsAffected)
	return nil
}
