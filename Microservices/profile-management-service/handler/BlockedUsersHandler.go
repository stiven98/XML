package handler

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"profile-management-service/model"
	"profile-management-service/service"
	"strings"
)

type BlockedUsersHandler struct {
	BlockedUsersService *service.BlockedUsersService
}

func (h BlockedUsersHandler) GetAllBlockedBy(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	id := tokens[int(len(tokens))-1]
	fmt.Println("Id of user: " + id)




	//fmt.Println(id)
	//h.BlockedUsersService()
}

func (h BlockedUsersHandler) BlockUserByUser(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	blockedByID := tokens[int(len(tokens))-2]
	blockedID := tokens[int(len(tokens))-1]
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
