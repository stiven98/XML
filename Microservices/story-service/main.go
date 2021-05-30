package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"storyservice/handler"
	"storyservice/repository"
	"storyservice/service"
)
func initRepo (database *redis.Client) *repository.StoriesRepository {
	return &repository.StoriesRepository{Database: database}
}
func initService (storyRepo *repository.StoriesRepository) *service.StoriesService {
	return &service.StoriesService{StoriesRepo: storyRepo}
}
func initHandler (storyService *service.StoriesService) *handler.StoriesHandler {
	return &handler.StoriesHandler{Service: storyService}
}
func handleFunc(storiesHandler *handler.StoriesHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/stories/create", storiesHandler.Create).Methods("POST")
	router.HandleFunc("/stories/getByKey/{key}", storiesHandler.GetByKey).Methods("GET")
	router.HandleFunc("/stories/get", storiesHandler.GetAca).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8083"), router))

}

func main() {
	fmt.Println("Aca")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	storiesRepo := initRepo(client)
	storiesService := initService(storiesRepo)
	storiesHandler := initHandler(storiesService)
	handleFunc(storiesHandler)
}


