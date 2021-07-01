package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"profile-management-service/model"
	"profile-management-service/service"
)

type SubscribeAccHandler struct {
	SubscriberAccService *service.SubscribeAccService
}

func (h SubscribeAccHandler) Subscribe(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	subscribedById := vars["subscribedById"]
	subscribedId := vars["subscribedId"]
	sub := model.SubscribeAcc{
		uuid.MustParse(subscribedById),
		uuid.MustParse(subscribedId),
	}

	err := h.SubscriberAccService.Subscribe(&sub)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

}

func (h SubscribeAccHandler) UnSubscribe(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	subscribedById := vars["subscribedById"]
	subscribedId := vars["subscribedId"]
	sub := model.SubscribeAcc{
		uuid.MustParse(subscribedById),
		uuid.MustParse(subscribedId),
	}

	err := h.SubscriberAccService.UnSubscribe(&sub)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusNoContent)
	writer.Header().Set("Content-Type", "application/json")

}

func (h SubscribeAccHandler) GetAllSubscribers(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	subscriber, err := h.SubscriberAccService.GetAllSubscribers(vars["id"])
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	renderJSON(writer,&subscriber)
}