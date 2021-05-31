package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"post_service/model"
	"post_service/service"
)
type CommentsHandler struct {
	Service *service.CommentsService
}


func (handler *CommentsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var comment model.Comment
	fmt.Println(json.NewDecoder(r.Body).Decode(&comment))
	err := json.NewDecoder(r.Body).Decode(&comment)

	err = handler.Service.Create(&comment)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *CommentsHandler) GetByKey(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	comment :=handler.Service.GetByKey(vars["key"])

	renderJSON(w, &comment)
}
