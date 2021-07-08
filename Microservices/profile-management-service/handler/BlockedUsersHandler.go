package handler

import (
	"fmt"
	"net/http"
	"profile-management-service/model"
	"profile-management-service/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type BlockedUsersHandler struct {
	BlockedUsersService  *service.BlockedUsersService
	SubscriberAccService *service.SubscribeAccService
	MutedUsersService    *service.MutedUsersService
	CloseFriendService   *service.CloseFriendsService
}

func (h BlockedUsersHandler) GetAllBlockedBy(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	blocked, err := h.BlockedUsersService.GetAllBlockedByUserId(vars["id"])

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	renderJSON(writer, &blocked)
}

func (h BlockedUsersHandler) BlockUserByUser(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	blockedByID := vars["blockedById"]
	blockedID := vars["blockedId"]
	blockedUser := model.BlockedUsers{
		BlockedByID: uuid.MustParse(blockedByID),
		BlockedID:   uuid.MustParse(blockedID),
	}

	// Check User exist in Profile service and return bad request if don't

	err := h.BlockedUsersService.BlockUserByUser(&blockedUser)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	subscribeUser := model.SubscribeAcc{SubscribeByID: uuid.MustParse(blockedByID), SubscribeID: uuid.MustParse(blockedID)}
	_ = h.SubscriberAccService.UnSubscribe(&subscribeUser)

	mutedUser := model.MutedUsers{MutedByID: uuid.MustParse(blockedByID), MutedID: uuid.MustParse(blockedID)}
	_ = h.MutedUsersService.UnMutedUserByUser(&mutedUser)

	closeFriend := model.CloseFriends{UserID: uuid.MustParse(blockedByID), FriendID: uuid.MustParse(blockedID)}
	_ = h.CloseFriendService.RemoveCloseFriend(&closeFriend)

	_, _ = http.Post("http://followers-microservice:8088/users/unfollow/"+blockedByID+"/"+blockedID, "application/json", nil)

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

}

func (h BlockedUsersHandler) IsBlocked(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	blockedByID := vars["blockedById"]
	blockedID := vars["blockedId"]
	blockedUser := model.BlockedUsers{
		BlockedByID: uuid.MustParse(blockedByID),
		BlockedID:   uuid.MustParse(blockedID),
	}

	isBlocked, err := h.BlockedUsersService.IsBlocked(&blockedUser)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	renderJSON(writer, isBlocked)
}

func (h BlockedUsersHandler) UnBlockUserByUser(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	blockedByID := vars["blockedById"]
	blockedID := vars["blockedId"]
	blockedUser := model.BlockedUsers{
		BlockedByID: uuid.MustParse(blockedByID),
		BlockedID:   uuid.MustParse(blockedID),
	}

	// Check User exist in Profile service and return bad request if don't

	err := h.BlockedUsersService.UnBlockUserByUser(&blockedUser)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

}
