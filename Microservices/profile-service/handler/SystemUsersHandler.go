package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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

func (handler *SystemUsersHandler) GetById(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	user:=handler.Service.GetById(uuid.MustParse(vars["id"]))
	renderJSON(w, &user)
}

func (handler *SystemUsersHandler) UpdateVerification(writer http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)
	handler.Service.UpdateVerification(uuid.MustParse(vars["id"]))
}

