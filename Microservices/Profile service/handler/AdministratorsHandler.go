package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"profileservice/model"
	"profileservice/service"
)

type AdministratorsHandler struct {
	Service *service.AdministratorsService
}

func (handler AdministratorsHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var user model.Administrator
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Update(&user)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}
