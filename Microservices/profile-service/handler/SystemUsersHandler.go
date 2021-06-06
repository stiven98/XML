package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
	"profileservice/model"
	"profileservice/service"
)

type SystemUsersHandler struct {
	Service *service.SystemUsersService
}

func (handler *SystemUsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.SystemUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SystemUsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user model.SystemUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *SystemUsersHandler) GetAll(w http.ResponseWriter, r *http.Request){
	users:=handler.Service.GetAll()
	renderJSON(w, &users)
}

func (handler *SystemUsersHandler) GetAllUsernames(w http.ResponseWriter, r *http.Request){
	usernames:=handler.Service.GetAllUsernames()
	renderJSON(w, &usernames)
}
func (handler *SystemUsersHandler) GetUserId(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	id:=handler.Service.GetUserId(vars["username"])
	renderJSON(w, &id)
}
