package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type AgentsService struct {
	AgentsRepo *repository.AgentsRepository
	SystemUserRepo *repository.SystemUsersRepository
}

func (service AgentsService) Update(agent *model.Agent) error {
	err := service.SystemUserRepo.Update(&agent.SystemUser)
	if err != nil {
		return err
	}
	err = service.AgentsRepo.Update(agent)
	if err != nil {
		return err
	}
	return nil
}
func (service *AgentsService) Create(agent *model.Agent) error {
	service.AgentsRepo.Create(agent)
	return nil
}

func (service *AgentsService) GetAll() []model.Agent {
	return  service.AgentsRepo.GetAll()
}


