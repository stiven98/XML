package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"profileservice/model"
	"profileservice/service"
)

type AgentsHandler struct {
	Service *service.AgentsService
}

func (handler AgentsHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var agent model.Agent
	err := json.NewDecoder(request.Body).Decode(&agent)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(&agent)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *AgentsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request model.AgentRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var agent model.Agent
	ID := uuid.New()
	agent.SystemUser.ID = ID
	agent.UserID = ID
	agent.SystemUser.FirstName = request.FirstName
	agent.SystemUser.LastName = request.LastName
	agent.SystemUser.Username = request.Username
	agent.SystemUser.Email = request.Email
	agent.SystemUser.Password, _ = HashPassword(request.Password)
	agent.SystemUser.Gender = request.Gender
	agent.SystemUser.TypeOfUser = model.AGENT
	agent.SystemUser.DateOfBirth = request.DateOfBirth
	agent.WebsiteLink = request.WebsiteLink
	request.IsApproved = true;
	err = handler.Service.Create(&agent)
	handler.Service.UpdateRequest(&request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
func (handler *AgentsHandler) DeclineRegistrationRequest(w http.ResponseWriter, r *http.Request) {
	var request model.AgentRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.DeclineRegistrationRequest(&request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *AgentsHandler) CerateRegistrationRequest(w http.ResponseWriter, r *http.Request) {
	var request model.AgentRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request.ID = uuid.New()
	request.UserID = uuid.New()
	request.TypeOfUser = model.AGENT
	request.IsApproved = false
	//fmt.Println(user)
	err = handler.Service.CreateRegistrationRequest(&request)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
func (handler *AgentsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	agents:=handler.Service.GetAll()
	renderJSON(w, &agents)
}
func (handler *AgentsHandler) GetAllRequests(w http.ResponseWriter, r *http.Request){
	requests:=handler.Service.GetAllRequests()
	renderJSON(w, requests)
}