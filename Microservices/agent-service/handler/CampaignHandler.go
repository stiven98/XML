package handler

import (
	"agent-service/model"
	"agent-service/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type CampaignHandler struct {
	Service *service.CampaignService
}


func (handler *CampaignHandler) Create(w http.ResponseWriter, r *http.Request) {
	var campaign model.Campaign
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&campaign)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.Create(&campaign)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *CampaignHandler) GetAll(w http.ResponseWriter, r *http.Request){
	campaigns:=handler.Service.GetAll()
	renderJSON(w, &campaigns)
}

func (handler *CampaignHandler) Delete(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	handler.Service.Delete(vars["id"])
	w.WriteHeader(http.StatusOK)
}