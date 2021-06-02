package handler

import (
	"admin-service/model"
	"admin-service/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AgentRegistrationsHandler struct {
	Service *service.AgentRegistrationsService
}

func (handler AgentRegistrationsHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var agentRegistration model.AgentRegistrationRequest
	err := json.NewDecoder(request.Body).Decode(&agentRegistration)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(&agentRegistration)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *AgentRegistrationsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var agentRegistration model.AgentRegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&agentRegistration)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&agentRegistration)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *AgentRegistrationsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	agentRegistrations:=handler.Service.GetAll()
	renderJSON(w, &agentRegistrations)
}
