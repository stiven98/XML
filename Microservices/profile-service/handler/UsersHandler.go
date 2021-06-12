package handler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"profileservice/model"
	"profileservice/model/Dto"
	"profileservice/service"
	"strings"
)

type UsersHandler struct {
	Service *service.UsersService
}

func (handler UsersHandler) Update(writer http.ResponseWriter, request *http.Request) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", "http://localhost:8080/auth/is-authenticated", nil)
	req.Header.Add("Authorization", request.Header.Get("Authorization"))
	_, errAuth := client.Do(req)

	if errAuth != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": request.RequestURI,
		}).Error("Failed decode request body")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Update(&user)
	if err != nil {
		fmt.Println(err)
		log.WithFields(log.Fields{
			"user": user.UserID,
			"handler": request.RequestURI,
		}).Error("Error during update")
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	log.WithFields(log.Fields{
		"user": user.UserID,
		"handler": request.RequestURI,
	}).Info("Profile successfully updated")
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		log.WithFields(log.Fields{
			"handler": r.RequestURI,
		}).Error("Failed decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userCreds := &Dto.UserCredentialsDto{
		USERNAME: user.SystemUser.Username,
		PASSWORD: user.SystemUser.Password,
		EMAIL:    user.SystemUser.Email,
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(userCreds)

	fmt.Println(payloadBuf)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", "https://localhost:8080/api/login-details/create", payloadBuf)
	req.Header.Set("Content-Type","application/json; charset=UTF-8")
	res, errAuth := client.Do(req)

	fmt.Println(res)

	if errAuth != nil {
		fmt.Println(errAuth)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.SystemUser.TypeOfUser = model.USER
	fmt.Println(user)
	err = handler.Service.Create(&user)
	if err != nil {
		fmt.Println(err)
		log.WithFields(log.Fields{
			"user": user.UserID,
			"handler": r.RequestURI,
		}).Error("Error during creation")
		w.WriteHeader(http.StatusBadRequest)
	}
	log.WithFields(log.Fields{
		"user": user.UserID,
		"handler": r.RequestURI,
	}).Info("User successfully created")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) GetById(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	user, _ :=handler.Service.GetById(vars["id"])
	log.WithFields(log.Fields{
		"user": user.UserID,
		"handler": r.RequestURI,
	}).Info("User successfully get by id")
	renderJSON(w, &user)
}

func (handler *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request){
	users:=handler.Service.GetAll()
	log.WithFields(log.Fields{
		"handler": r.RequestURI,
	}).Info("Get all users successfully")
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

func (handler UsersHandler) IsPublic(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	ID := tokens[int(len(tokens))-1]

	fmt.Println("ID")
	user, err := handler.Service.GetById(ID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(user)
	res := Dto.IsUserPublicDTO {
		ID: user.UserID,
		IsPublic: user.IsPublic,
	}
	renderJSON(writer, res)
	writer.WriteHeader(http.StatusOK)
}

func (handler *UsersHandler) GetPublicUsersIds(w http.ResponseWriter, r *http.Request){
	var ids Dto.PublicUsersIdsDto
	users := handler.Service.GetAllPublic()
	for i := range users {
		ids.KEYS = append(ids.KEYS, users[i].UserID.String())
	}
	renderJSON(w, &ids)
}