package repository

import (
	"admin-service/model"
	"fmt"
	"gorm.io/gorm"
)

type AgentRegistrationsRepository struct {
	Database *gorm.DB
}

func (repo *AgentRegistrationsRepository) Update(request *model.AgentRegistrationRequest) error {
	result := repo.Database.Model(model.AgentRegistrationRequest{}).Where("id = ?", request.ID).Updates(request)
	return result.Error
}
func(repo *AgentRegistrationsRepository) GetAll() []model.AgentRegistrationRequest{
	var requests []model.AgentRegistrationRequest
	repo.Database.Find(&requests)
	return requests
}

func (repo *AgentRegistrationsRepository) Create(request *model.AgentRegistrationRequest) error  {
	result := repo.Database.Create(request)
	fmt.Println(result.RowsAffected)
	return nil
}
