package service

import (
	"admin-service/model"
	"admin-service/repository"
)

type AgentRegistrationsService struct {
	AgentRegistrationsRepo *repository.AgentRegistrationsRepository
}

func (service AgentRegistrationsService) Update(request *model.AgentRegistrationRequest) error {
	err := service.AgentRegistrationsRepo.Update(request)
	if err != nil {
		return err
	}
	return nil
}
func (service *AgentRegistrationsService) Create(request *model.AgentRegistrationRequest) error {
	service.AgentRegistrationsRepo.Create(request)
	return nil
}

func (service *AgentRegistrationsService) GetAll() []model.AgentRegistrationRequest {
	return  service.AgentRegistrationsRepo.GetAll()
}
