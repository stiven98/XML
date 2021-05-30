package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profileservice/model"
	"profileservice/service"
)

type MessagesHandler struct {
	Service *service.MessagesService
}


func (handler *MessagesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var message model.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&message)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *MessagesHandler) GetAll(w http.ResponseWriter, r *http.Request){
	conversations:=handler.Service.GetAll()
	renderJSON(w, &conversations)
}