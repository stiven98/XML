package repository

import (
	"auth-service/model"
	"fmt"
	"gorm.io/gorm"
)

type LoginDetailsRepository struct {
	Database *gorm.DB
}

func (repo *LoginDetailsRepository) Update(loginDetails *model.LoginDetails) error {
	result := repo.Database.Model(model.LoginDetails{}).Where("id = ?", loginDetails.ID).Updates(loginDetails)
	return result.Error
}

func(repo *LoginDetailsRepository) GetAll() []model.LoginDetails{
	var loginDetailsList []model.LoginDetails
	repo.Database.Find(&loginDetailsList)
	return loginDetailsList
}

func(repo *LoginDetailsRepository) GetByEmail(email string) model.LoginDetails{
	var loginDetails model.LoginDetails
	repo.Database.Find(&loginDetails).Where("email = ?", email)
	return loginDetails
}

func (repo *LoginDetailsRepository) Create(loginDetails *model.LoginDetails) error {
	result := repo.Database.Create(loginDetails)
	fmt.Println(result.RowsAffected)
	return nil
}