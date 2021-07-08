package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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

	router.HandleFunc("/story", storiesHandler.Create).Methods("POST")
	router.HandleFunc("/upload", storiesHandler.UploadFile).Methods("POST")
	router.HandleFunc("/story/feed/{id}", storiesHandler.GetFeed).Methods("GET")
	router.HandleFunc("/story/paged-feed", storiesHandler.GetPagedFeed).Methods("GET")
	router.HandleFunc("/story/my-paged-stories", storiesHandler.GetMyPagedStories).Methods("GET")
	router.HandleFunc("/story/my/{id}", storiesHandler.GetMyStories).Methods("GET")
	router.HandleFunc("/story/highlight", storiesHandler.AddToHighlights).Methods("POST")
	router.HandleFunc("/story/remove-highlight", storiesHandler.RemoveFromHighlights).Methods("POST")
	router.HandleFunc("/story/paged-highlights", storiesHandler.GetPagedHighlights).Methods("GET")
	router.Handle("/images/{rest}",
		http.StripPrefix("/images/", http.FileServer(http.Dir("./user_stories/"))))
	headers := handlers.AllowedHeaders([] string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([] string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([] string{"*"})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8083"), handlers.CORS(headers, methods, origins)(router)))

}

func main() {
	hostName := os.Getenv("HOST_NAME")
	host := "localhost"
	if len(hostName) != 0 {
		host = hostName
	}
	client := redis.NewClient(&redis.Options{
		Addr: host + ":6379",
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


