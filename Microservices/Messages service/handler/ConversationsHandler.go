package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profileservice/model"
	"profileservice/service"
)

type ConversationsHandler struct {
	Service *service.ConversationsService
}


func (handler *ConversationsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var conversation model.Conversation
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&conversation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(conversation.MESSAGES[0].VALUE)
	err = handler.Service.Create(&conversation)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *ConversationsHandler) GetAll(w http.ResponseWriter, r *http.Request){
	conversations:=handler.Service.GetAll()
	renderJSON(w, &conversations)
}