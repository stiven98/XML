package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"storyservice/model"
	"storyservice/service"
)
type StoriesHandler struct {
	Service *service.StoriesService
}

func (handler *StoriesHandler) GetAca (w http.ResponseWriter, r *http.Request){

	var story model.Story
	err := json.NewDecoder(r.Body).Decode(&story)
	fmt.Println(err)
	renderJSON(w, &story)

}


func (handler *StoriesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var story model.Story
	fmt.Println(json.NewDecoder(r.Body).Decode(&story))
	err := json.NewDecoder(r.Body).Decode(&story)

	err = handler.Service.Create(&story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *StoriesHandler) GetByKey(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	story:=handler.Service.GetByKey(vars["key"])

	renderJSON(w, &story)
}
