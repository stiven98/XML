package handler

import (
	"auth-service/model"
	"auth-service/service"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type LoginDetailsHandler struct {
	Service *service.LoginDetailsService
}

func (handler *LoginDetailsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var loginDetails model.LoginDetails
	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&loginDetails)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LoginDetailsHandler) Update(w http.ResponseWriter, r *http.Request) {
	var loginDetails model.LoginDetails
	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(&loginDetails)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *LoginDetailsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	loginDetailsList:=handler.Service.GetAll()
	renderJSON(w, &loginDetailsList)
}

func (handler *LoginDetailsHandler) GetByEmail(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	email, ok:=vars["email"]
	if !ok {
		fmt.Println("Email is missing in parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	loginDetailsList := handler.Service.GetByEmail(email)
	renderJSON(w, &loginDetailsList)
}
