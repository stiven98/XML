package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"profileservice/model"
	"profileservice/model/Dto"
	"profileservice/service"

)

type UsersHandler struct {
	Service *service.UsersService
}

func (handler UsersHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Update(&user)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.SystemUser.TypeOfUser = model.USER
	fmt.Println(user)
	err = handler.Service.Create(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) GetById(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	user:=handler.Service.GetById(uuid.MustParse(vars["id"]))
	renderJSON(w, &user)
}

func (handler *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request){
	users:=handler.Service.GetAll()
	renderJSON(w, &users)
}
func (handler *UsersHandler) ChangeWhetherIsPublic(w http.ResponseWriter, r *http.Request) {
	var dto Dto.ChangeWhetherIsPublicDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.ChangeWhetherIsPublic(dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) ChangeAllowedTags(w http.ResponseWriter, r *http.Request) {
	var dto Dto.ChangeAllowedTagsDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.ChangeAllowedTags(dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}