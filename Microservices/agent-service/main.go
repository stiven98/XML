package main

import (
	"agent-service/handler"
	"agent-service/repository"
	"agent-service/service"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

var client *mongo.Client

func initDB() *mongo.Database{
	log.Println("Connecting to database...")
	hostName := os.Getenv("HOST_NAME")
	host := "localhost"
	if len(hostName) != 0 {
		host = hostName
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+ host + ":27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	database := client.Database("agent_service")
	log.Println("Connected to database...")

	return database
}

func initRepo (database *mongo.Database) (*repository.CampaignRepository) {
	return &repository.CampaignRepository{Database: database}
}

func initServices (campaignRepo *repository.CampaignRepository) (*service.CampaignService) {
	return &service.CampaignService{CampaignRepo: campaignRepo}
}

func initHandlers (campaignService *service.CampaignService) (*handler.CampaignHandler) {
	return &handler.CampaignHandler{Service: campaignService}
}
func handleFunc(agentHandler *handler.CampaignHandler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/campaigns/create", agentHandler.Create).Methods("POST")
	router.HandleFunc("/campaigns/getAll",agentHandler.GetAll).Methods("GET")
	router.HandleFunc("/campaigns/delete/{id}", agentHandler.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8001"), router))
}
func main() {
	database := initDB()
	campaignRepo := initRepo(database)
	campaignService := initServices(campaignRepo)
	campaignHandler := initHandlers(campaignService)
	handleFunc(campaignHandler)
	fmt.Println("aca", database)
}

