package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"shop-backend/model"
	"shop-backend/model/dto"
	"shop-backend/service"
)

type UsersHandler struct {
	UsersService *service.UsersService
}

func (h UsersHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user.UserID = uuid.New()
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UsersService.Create(user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (h UsersHandler) Login(writer http.ResponseWriter, request *http.Request) {
	var loginInfo dto.LoginInfoDTO
	err := json.NewDecoder(request.Body).Decode(&loginInfo)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var user model.User
	user, count := h.UsersService.Login(loginInfo)

	if count != 1 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	loginInfo.UserID = user.UserID
	renderJSON(writer, &loginInfo)


}
