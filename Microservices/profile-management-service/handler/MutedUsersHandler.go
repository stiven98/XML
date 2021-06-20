package handler

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"profile-management-service/model"
	"profile-management-service/service"
)

type MutedUsersHandler struct {
	MutedUsersService *service.MutedUsersService
}

func (h MutedUsersHandler) IsMuted(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	mutedByID := vars["mutedById"]
	mutedID := vars["mutedId"]
	mutedUser := model.MutedUsers{
		MutedByID: uuid.MustParse(mutedByID),
		MutedID:   uuid.MustParse(mutedID),
	}

	isMuted, err := h.MutedUsersService.IsMuted(&mutedUser)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}

	renderJSON(writer,isMuted)
}


func (h MutedUsersHandler) MutedUserByUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	mutedById := vars["mutedById"]
	mutedId := vars["mutedId"]
	muted := model.MutedUsers{
		MutedByID: uuid.MustParse(mutedById),
		MutedID:   uuid.MustParse(mutedId),
	}

	// Check User exist in Profile service and return bad request if don't

	err := h.MutedUsersService.MutedUserByUser(&muted)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")

}


func (h MutedUsersHandler) UnMutedUserByUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	mutedById := vars["mutedById"]
	mutedId := vars["mutedId"]
	muted := model.MutedUsers{
		MutedByID: uuid.MustParse(mutedById),
		MutedID:   uuid.MustParse(mutedId),
	}

	// Check User exist in Profile service and return bad request if don't

	err := h.MutedUsersService.UnMutedUserByUser(&muted)

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	writer.WriteHeader(http.StatusNoContent)
	writer.Header().Set("Content-Type", "application/json")

}