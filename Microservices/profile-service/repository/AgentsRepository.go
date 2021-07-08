package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"profileservice/model"
	"profileservice/model/Dto"

	"gorm.io/gorm"
)

type AgentsRepository struct {
	Database *gorm.DB
}

func (repo *AgentsRepository) Update(agent *model.Agent) error {
	result := repo.Database.Model(model.Agent{}).Where("user_id = ?", agent.UserID).Updates(agent)
	return result.Error
}
func (repo *AgentsRepository) UpdateRequest(request *model.AgentRegistrationRequest) error {
	result := repo.Database.Model(model.AgentRegistrationRequest{}).Where("id = ?", request.ID).Updates(request)
	return result.Error
}
func (repo *AgentsRepository) GetAll() []model.Agent {
	var agents []model.Agent
	repo.Database.Preload("SystemUser").Find(&agents)
	return agents
}

func (repo *AgentsRepository) Create(agent *model.User) error {
	result := repo.Database.Create(agent)
	var dto = Dto.CreateUserDTO{
		ID:       agent.UserID,
		USERNAME: agent.SystemUser.Username,
		PASSWORD: agent.SystemUser.Password,
		ACTIVE:   true,
		ROLE:     "ROLE_AGENT",
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(dto)
	payloadBuf1 := new(bytes.Buffer)
	json.NewEncoder(payloadBuf1).Encode(dto.ID)
	_, err := http.Post("http://auth-service:8080/api/createUser", "application/json", payloadBuf)
	_, err1 := http.Post("http://followers-microservice:8088/users/addNode/"+dto.ID.String(), "application/json", payloadBuf1)

	if err != nil {
		fmt.Println(err)
		return err
	}
	if err1 != nil {
		fmt.Println(err1)
		return err1
	}
	fmt.Println(result.RowsAffected)
	return nil
}

func (repo *AgentsRepository) CreateRegistrationRequest(request *model.AgentRegistrationRequest) error {
	result := repo.Database.Create(request)
	fmt.Println(result.RowsAffected)
	return nil
}
func (repo *AgentsRepository) GetAllRequests() []model.AgentRegistrationRequest {
	var requests []model.AgentRegistrationRequest
	repo.Database.Find(&requests)
	return requests
}

func (repo *AgentsRepository) DeclineRegistrationRequest(request *model.AgentRegistrationRequest) error {
	result := repo.Database.Delete(request)
	fmt.Println(result.RowsAffected)
	return nil
}
