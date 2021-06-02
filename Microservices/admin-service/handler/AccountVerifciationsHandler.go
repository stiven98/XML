package handler

import (
	"admin-service/model"
	"admin-service/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountVerificationsHandler struct {
	Service *service.AccountVerificationsService
}

func (handler AccountVerificationsHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var verificationRequest model.AccountVerificationRequest
	err := json.NewDecoder(request.Body).Decode(&verificationRequest)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(&verificationRequest)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *AccountVerificationsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var verificationRequest model.AccountVerificationRequest
	err := json.NewDecoder(r.Body).Decode(&verificationRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&verificationRequest)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *AccountVerificationsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	verificationRequests:=handler.Service.GetAll()
	renderJSON(w, &verificationRequests)
}
