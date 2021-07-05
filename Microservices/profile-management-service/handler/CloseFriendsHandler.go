package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"profile-management-service/model"
	"profile-management-service/service"
)

type CloseFriendHandler struct {
	CloseFriendService *service.CloseFriendsService
}

func (h CloseFriendHandler) AddCloseFriend(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	userId := vars["userId"]
	friendId := vars["friendId"]
	friend := model.CloseFriends{
		UserID:   uuid.MustParse(userId),
		FriendID: uuid.MustParse(friendId),
	}
	err := h.CloseFriendService.AddCloseFriend(&friend)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}


func (h CloseFriendHandler) RemoveCloseFriend(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	userId := vars["userId"]
	friendId := vars["friendId"]
	friend := model.CloseFriends{
		UserID:   uuid.MustParse(userId),
		FriendID: uuid.MustParse(friendId),
	}
	err := h.CloseFriendService.RemoveCloseFriend(&friend)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

}

func (h CloseFriendHandler)  GetAllCloseFriend(writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	friends, err := h.CloseFriendService.GetAllCloseFriend(vars["id"])

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	renderJSON(writer,&friends)
}