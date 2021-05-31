package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"post_service/model"
	"post_service/service"
)

type PostsHandler struct {
	Service *service.PostsService
}


func (handler *PostsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	fmt.Println(json.NewDecoder(r.Body).Decode(&post))
	err := json.NewDecoder(r.Body).Decode(&post)

	err = handler.Service.Create(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *PostsHandler) GetByKey(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	post :=handler.Service.GetByKey(vars["key"])

	renderJSON(w, &post)
}
