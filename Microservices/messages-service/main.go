package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"profileservice/handler"
	"profileservice/repository"
	"profileservice/service"
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
	database := client.Database("messages_service")
	log.Println("Connected to database...")

	return database
}

func initRepo (database *mongo.Database) (*repository.MessagesRepository, *repository.ConversationRepository) {
	return &repository.MessagesRepository{Database: database}, &repository.ConversationRepository{Database: database}
}

func initServices (messagesRepo *repository.MessagesRepository, conversationRepo *repository.ConversationRepository) (*service.MessagesService, *service.ConversationsService) {
	return &service.MessagesService{MessagesRepo: messagesRepo}, &service.ConversationsService{ConversationsRepo: conversationRepo}
}

func initHandlers (messagesService *service.MessagesService, conversationsService *service.ConversationsService) (*handler.MessagesHandler, *handler.ConversationsHandler) {
	return &handler.MessagesHandler{Service: messagesService}, &handler.ConversationsHandler{Service: conversationsService}
}
func handleFunc(messagesHandler *handler.MessagesHandler, conversationsHandler *handler.ConversationsHandler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/messages/create", messagesHandler.Create).Methods("POST")
	router.HandleFunc("/messages/getAll",messagesHandler.GetAll).Methods("GET")
	router.HandleFunc("/conversations/create", conversationsHandler.Create).Methods("POST")
	router.HandleFunc("/conversations/getAll",conversationsHandler.GetAll).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8002"), router))
}
func main() {
	database := initDB()
	messagesRepo, conversationsRepo := initRepo(database)
	messagesService, conversationsService := initServices(messagesRepo, conversationsRepo)
	messagesHandler, conversationsHandler := initHandlers(messagesService, conversationsService)
	handleFunc(messagesHandler, conversationsHandler)
	fmt.Println("aca", database)
}
