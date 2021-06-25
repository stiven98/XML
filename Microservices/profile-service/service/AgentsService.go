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
func (service AgentsService) UpdateRequest(request *model.AgentRegistrationRequest) error {
	err := service.AgentsRepo.UpdateRequest(request)
	if err != nil {
		return err
	}
	return nil
}
func (service AgentsService) DeclineRegistrationRequest(request *model.AgentRegistrationRequest) error {
	err := service.AgentsRepo.DeclineRegistrationRequest(request)
	if err != nil {
		return err
	}
	return nil
}
func (service *AgentsService) Create(agent *model.Agent) error {
	service.AgentsRepo.Create(agent)
	return nil
}
func (service *AgentsService) CreateRegistrationRequest(request *model.AgentRegistrationRequest) error {
	return service.AgentsRepo.CreateRegistrationRequest(request)

}
func (service *AgentsService) GetAll() []model.Agent {
	return  service.AgentsRepo.GetAll()
}
func (service *AgentsService) GetAllRequests() []model.AgentRegistrationRequest {
	return  service.AgentsRepo.GetAllRequests()
}


