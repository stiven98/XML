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

func initRepo (database *redis.Client) (*repository.PostsRepository) {
	return &repository.PostsRepository{Database: database}
}
func initService ( postRepo *repository.PostsRepository) ( *service.PostsService) {
	return &service.PostsService{PostsRepo: postRepo}
}
func initHandler (postService *service.PostsService) (*handler.PostsHandler) {
	return &handler.PostsHandler{Service: postService}
}
func handleFunc(postsHandler *handler.PostsHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/comments", postsHandler.LeaveComment).Methods("POST")
	router.HandleFunc("/posts/create", postsHandler.Create).Methods("POST")
	router.HandleFunc("/posts/save", postsHandler.Save).Methods("POST")
	router.HandleFunc("/posts/all-archived/{id}", postsHandler.GetAllArchived).Methods("GET")
	router.HandleFunc("/posts/edit-archived", postsHandler.EditArchived).Methods("POST")
	router.HandleFunc("/campaigns/createCampaign", postsHandler.CreateCampaign).Methods("POST")
	router.HandleFunc("/campaigns/createRequest", postsHandler.CreateCampaignRequest).Methods("POST")
	router.HandleFunc("/campaigns/updateCampaign", postsHandler.UpdateCampaign).Methods("POST")
	router.HandleFunc("/campaigns/addInfluencer", postsHandler.AddInfluencer).Methods("POST")
	router.HandleFunc("/posts/getByKey/{key}", postsHandler.GetByKey).Methods("GET")
	router.HandleFunc("/posts/getById/{userid}/{postid}", postsHandler.GetByIds).Methods("GET")
	router.HandleFunc("/campaigns/getById/{userid}/{campaignid}", postsHandler.GetCampaignsByIds).Methods("GET")
	router.HandleFunc("/campaigns/getByInfluencerId/{userid}/{campaignid}", postsHandler.GetCampaignsByInfluencerIds).Methods("GET")

	router.HandleFunc("/posts/getByUserId/{userid}", postsHandler.GetByUserId).Methods("GET")
	router.HandleFunc("/posts/feed/{id}", postsHandler.GetFeed).Methods("GET")
	router.HandleFunc("/posts/public/{id}", postsHandler.GetPublic).Methods("GET")
	router.HandleFunc("/posts/liked/{id}", postsHandler.GetLiked).Methods("GET")
	router.HandleFunc("/campaigns/getUserCampaigns/{id}", postsHandler.GetCampaigns).Methods("GET")
	router.HandleFunc("/campaigns/getInfluencerCampaigns/{id}", postsHandler.GetInfluencerCampaigns).Methods("GET")
	router.HandleFunc("/campaigns/getUserCampaignReqs/{id}", postsHandler.GetCampaignReqs).Methods("GET")
	router.HandleFunc("/campaigns/getUserTemporaryCampaigns/{id}", postsHandler.GetTemporaryCampaigns).Methods("GET")
	router.HandleFunc("/posts/disliked/{id}", postsHandler.GetDisliked).Methods("GET")
	router.HandleFunc("/posts/reported", postsHandler.GetReported).Methods("GET")
	router.HandleFunc("/posts/delete", postsHandler.DeletePost).Methods("POST")
	router.HandleFunc("/campaigns/delete", postsHandler.DeleteCampaign).Methods("POST")
	router.HandleFunc("/campaigns/deleteReq", postsHandler.DeleteCampaignReq).Methods("POST")
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

	postRepo := initRepo(client)
	postService := initService(postRepo)
	postHandler := initHandler( postService)
	handleFunc(postHandler)
}