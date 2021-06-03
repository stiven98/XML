package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profileservice/model"
	"profileservice/service"
)

type AdministratorsHandler struct {
	Service *service.AdministratorsService
}

func (handler AdministratorsHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var admin model.Administrator
	err := json.NewDecoder(request.Body).Decode(&admin)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(&admin)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *AdministratorsHandler) Create(w http.ResponseWriter, r *http.Request) {

	var admin model.Administrator
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&admin)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *AdministratorsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	admins:=handler.Service.GetAll()
	renderJSON(w, &admins)
}
