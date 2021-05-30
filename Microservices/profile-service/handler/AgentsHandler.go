package handler

import (
	"encoding/json"
	"fmt"
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
	var agent model.Agent
	err := json.NewDecoder(r.Body).Decode(&agent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&agent)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *AgentsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	agents:=handler.Service.GetAll()
	renderJSON(w, &agents)
}
