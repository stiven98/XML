package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
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
	return &handler.MessagesHandler{MessageService: messagesService, ConversationService: conversationsService}, &handler.ConversationsHandler{Service: conversationsService}
}
func handleFunc(messagesHandler *handler.MessagesHandler, conversationsHandler *handler.ConversationsHandler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/messages/add/{user1}/{user2}", messagesHandler.Add).Methods("POST")
	router.HandleFunc("/messages/getAll",messagesHandler.GetAll).Methods("GET")
	router.HandleFunc("/conversations/create", conversationsHandler.Create).Methods("POST")
	router.HandleFunc("/conversations/getAll",conversationsHandler.GetAll).Methods("GET")
	router.HandleFunc("/conversations/{user1}/{user2}", conversationsHandler.GetConversation).Methods("GET")
	router.HandleFunc("/images/upload", messagesHandler.ImageUpload).Methods("POST")
	router.Handle("/images/{rest}",
		http.StripPrefix("/images/", http.FileServer(http.Dir("./messages_images/"))))

	headers := handlers.AllowedHeaders([] string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([] string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([] string{"*"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8002"), handlers.CORS(headers, methods, origins) (router)))
}
func main() {
	database := initDB()
	messagesRepo, conversationsRepo := initRepo(database)
	messagesService, conversationsService := initServices(messagesRepo, conversationsRepo)
	messagesHandler, conversationsHandler := initHandlers(messagesService, conversationsService)
	handleFunc(messagesHandler, conversationsHandler)
}
