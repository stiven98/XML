package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"post_service/handler"
	"post_service/repository"
	"post_service/service"
)

func initRepo (database *redis.Client) *repository.CommentsRepository {
	return &repository.CommentsRepository{Database: database}
}
func initService (commentRepo *repository.CommentsRepository) *service.CommentsService {
	return &service.CommentsService{CommentsRepo: commentRepo}
}
func initHandler (commentService *service.CommentsService) *handler.CommentsHandler {
	return &handler.CommentsHandler{Service: commentService}
}
func handleFunc(commentsHandler *handler.CommentsHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/comments/create", commentsHandler.Create).Methods("POST")
	router.HandleFunc("/comments/getByKey/{key}", commentsHandler.GetByKey).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8087"), router))

}

func main() {
	fmt.Println("Aca")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 1,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	commentRepo := initRepo(client)
	commentService := initService(commentRepo)
	commentHandler := initHandler(commentService)
	handleFunc(commentHandler)
}