package repository

import (
	"fmt"
	"gorm.io/gorm"
	"profileservice/model"
)

type AgentsRepository struct {
	Database *gorm.DB
}

func (repo *AgentsRepository) Update(agent *model.Agent) error {
	result := repo.Database.Model(model.Agent{}).Where("user_id = ?", agent.UserID).Updates(agent)
	return result.Error
}
func(repo *AgentsRepository) GetAll() []model.Agent{
	var agents []model.Agent
	repo.Database.Preload("SystemUser").Find(&agents)
	return agents
}

func (repo *AgentsRepository) Create(agent *model.Agent) error {
	result := repo.Database.Create(agent)
	fmt.Println(result.RowsAffected)
	return nil
}
