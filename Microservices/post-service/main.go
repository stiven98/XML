package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"post_service/handler"
	"post_service/repository"
	"post_service/service"
)

func initRepo (database *redis.Client) (*repository.CommentsRepository, *repository.PostsRepository) {
	return &repository.CommentsRepository{Database: database}, &repository.PostsRepository{Database: database}
}
func initService (commentRepo *repository.CommentsRepository, postRepo *repository.PostsRepository) (*service.CommentsService, *service.PostsService) {
	return &service.CommentsService{CommentsRepo: commentRepo}, &service.PostsService{PostsRepo: postRepo}
}
func initHandler (commentService *service.CommentsService, postService *service.PostsService) (*handler.CommentsHandler, *handler.PostsHandler) {
	return &handler.CommentsHandler{Service: commentService}, &handler.PostsHandler{Service: postService}
}
func handleFunc(commentsHandler *handler.CommentsHandler, postsHandler *handler.PostsHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/comments/create", commentsHandler.Create).Methods("POST")
	router.HandleFunc("/comments/getByKey/{key}", commentsHandler.GetByKey).Methods("GET")
	router.HandleFunc("/posts/create", postsHandler.Create).Methods("POST")
	router.HandleFunc("/posts/getByKey/{key}", postsHandler.GetByKey).Methods("GET")
	router.HandleFunc("/posts/feed/{id}", postsHandler.GetFeed).Methods("GET")
	router.HandleFunc("/posts/public", postsHandler.GetPublic).Methods("GET")
	router.HandleFunc("/posts/liked/{id}", postsHandler.GetLiked).Methods("GET")
	router.HandleFunc("/posts/disliked/{id}", postsHandler.GetDisliked).Methods("GET")
	router.HandleFunc("/posts/reported", postsHandler.GetReported).Methods("GET")
	router.HandleFunc("/posts/delete", postsHandler.DeletePost).Methods("POST")
	router.HandleFunc("/posts/public-tags", postsHandler.GetAllTagsPublic).Methods("GET")
	router.HandleFunc("/posts/public-locations", postsHandler.GetAllLocationsPublic).Methods("GET")
	router.HandleFunc("/posts/signed-in-tags/{id}", postsHandler.GetAllTagsSignedIn).Methods("GET")
	router.HandleFunc("/posts/signed-in-locations/{id}", postsHandler.GetAllLocationsSignedIn).Methods("GET")
	router.HandleFunc("/upload", postsHandler.UploadFile).Methods("POST")
	router.HandleFunc("/like-post", postsHandler.LikePost).Methods("POST")
	router.HandleFunc("/dislike-post", postsHandler.DislikePost).Methods("POST")
	router.HandleFunc("/report-post", postsHandler.ReportPost).Methods("POST")
	router.Handle("/images/{rest}",
		http.StripPrefix("/images/", http.FileServer(http.Dir("./user_posts/"))))
	headers := handlers.AllowedHeaders([] string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([] string{"GET", "POST", "PUT"})
	origins := handlers.AllowedOrigins([] string{"*"})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8086"), handlers.CORS(headers, methods, origins) (router)))

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
		DB: 1,
	})
	pong, err := client.Ping().Result()
	fmt.Println(client.MGet())
	fmt.Println(pong, err)

	commentRepo, postRepo := initRepo(client)
	commentService, postService := initService(commentRepo,postRepo)
	commentHandler, postHandler := initHandler(commentService, postService)
	handleFunc(commentHandler, postHandler)
}