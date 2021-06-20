package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"profile-management-service/model"
	"profile-management-service/service"
)

type BlockedUsersHandler struct {
	BlockedUsersService *service.BlockedUsersService
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
	renderJSON(writer,&blocked)
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

	renderJSON(writer,isBlocked)
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


