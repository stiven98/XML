package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"profileservice/model"
	"profileservice/service"
)

type  NotifyHandler struct {
	NotifyService *service.NotifyService
}

func (handler NotifyHandler) Create(writer http.ResponseWriter, request *http.Request){
	var notify model.Notify
	err := json.NewDecoder(request.Body).Decode(&notify)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	notify.ID, _ = uuid.NewUUID()
	err = handler.NotifyService.Create(&notify)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}


func (handler NotifyHandler) GetAllNotifyByUserId(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	fmt.Println(vars["id"])
	notify ,err := handler.NotifyService.GetAllNotifyByUserId(vars["id"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")


	renderJSON(writer, &notify)
}