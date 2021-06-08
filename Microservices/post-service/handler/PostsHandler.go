package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"post_service/model"
	"post_service/model/dto"
	"post_service/service"
	"strings"
	"time"
)

type PostsHandler struct {
	Service *service.PostsService
}


func (handler *PostsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	fmt.Println(json.NewDecoder(r.Body).Decode(&post))
	err := json.NewDecoder(r.Body).Decode(&post)
	post.ID = uuid.New()
	post.TIMESTAMP = time.Now()
	post.COMMENTS = []model.Comment{}
	post.LIKES = []model.Like{}
	post.DISLIKES = []model.Dislike{}
	err = handler.Service.Create(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	rest, err := http.Get("http://localhost:8088/users/getFollowers/" + post.USERID.String())
	//rest, err := http.Get("https://mocki.io/v1/84324533-ee57-4eb2-8042-3f5845dcc41b")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto dto.FollowersDto
	err = json.NewDecoder(rest.Body).Decode(&dto)
	fmt.Println(dto.KEYS)
	err = handler.Service.AddPostToFeed(dto.KEYS, &post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (postsHandler *PostsHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	var postItems[] model.PostItem
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	fhs := r.MultipartForm.File["files"]
	var i int
	i = 0
	for _, fh := range fhs {
		i = i + 1
		f, err := fh.Open()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Printf("Uploaded File: %+v\n", fh.Filename)
		fmt.Printf("File Size: %+v\n", fh.Size)
		fmt.Printf("MIME Header: %+v\n", fh.Header)
		fileName := strings.Split(fh.Filename, ".")
		var filePath string
		var resourceName string
		if(len(fileName) >= 2){
			resourceName = uuid.NewString() +  "." + fileName[1]
			filePath = filepath.Join("user_posts", resourceName)
		}else{
			filePath = filepath.Join("user_posts", fh.Filename)
		}
		dst, err := os.Create(filePath)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var postItem model.PostItem
		postItem.PATH = resourceName
		postItem.ID = uuid.New()
		postItems = append(postItems, postItem)
		defer f.Close()

	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	renderJSON(w, &postItems)
}

func (handler *PostsHandler) GetByKey(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	posts :=handler.Service.GetByKey(vars["key"])
	renderJSON(w, &posts)
}

func (handler *PostsHandler) LikePost(w http.ResponseWriter, r *http.Request){
	var likeReq dto.LikeDto
	err := json.NewDecoder(r.Body).Decode(&likeReq)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	err = handler.Service.LikePost(likeReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostsHandler) DislikePost(w http.ResponseWriter, r *http.Request){
	var dislikeReq dto.LikeDto
	err := json.NewDecoder(r.Body).Decode(&dislikeReq)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	err = handler.Service.DislikePost(dislikeReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *PostsHandler) GetFeed(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	post :=handler.Service.GetFeed(vars["id"])
	renderJSON(w, &post)
}

func (handler *PostsHandler) GetPublic(w http.ResponseWriter, r *http.Request){
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dto.KEYS)
	posts := handler.Service.GetPublic(dto.KEYS)
	renderJSON(w, &posts)
}

func (handler *PostsHandler) GetAllTagsPublic(w http.ResponseWriter, r *http.Request){
	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)
	posts := handler.Service.GetPublic(dtoIds.KEYS)

	var tags dto.KeyValueListDto

	for i := range posts {
		tags.KEYS = append(tags.KEYS, posts[i].HASHTAG)
	}

	renderJSON(w, &tags)
}

func (handler *PostsHandler) GetAllTagsSignedIn(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	fmt.Println(vars["id"])

	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)


	restFollowers, _ := http.Get("http://localhost:8088/users/getFollowers/" + vars["id"])
	var dtoFollowers dto.FollowersDto
	err = json.NewDecoder(restFollowers.Body).Decode(&dtoFollowers)
	fmt.Println(dtoFollowers.KEYS)

	var idsList []string

	idsList = append(idsList, dtoIds.KEYS...)
	idsList = append(idsList, dtoFollowers.KEYS...)

	posts := handler.Service.GetPublic(idsList)

	var tags dto.KeyValueListDto

	for i := range posts {
		tags.KEYS = append(tags.KEYS, posts[i].HASHTAG)
	}

	renderJSON(w, &tags)
}

func (handler *PostsHandler) GetAllLocationsPublic(w http.ResponseWriter, r *http.Request){
	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)
	posts := handler.Service.GetPublic(dtoIds.KEYS)

	var locations dto.KeyValueListDto

	for i := range posts {
		locations.KEYS = append(locations.KEYS, posts[i].LOCATION)
	}

	renderJSON(w, &locations)
}

func (handler *PostsHandler) GetAllLocationsSignedIn(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	fmt.Println(vars["id"])

	//public users ids
	rest, err := http.Get("http://localhost:8085/users/public-ids")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dtoIds dto.KeyValueListDto
	err = json.NewDecoder(rest.Body).Decode(&dtoIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(dtoIds.KEYS)


	restFollowers, _ := http.Get("http://localhost:8088/users/getFollowers/" + vars["id"])
	var dtoFollowers dto.FollowersDto
	err = json.NewDecoder(restFollowers.Body).Decode(&dtoFollowers)
	fmt.Println(dtoFollowers.KEYS)

	var idsList []string

	idsList = append(idsList, dtoIds.KEYS...)
	idsList = append(idsList, dtoFollowers.KEYS...)

	posts := handler.Service.GetPublic(idsList)

	var locations dto.KeyValueListDto

	for i := range posts {
		locations.KEYS = append(locations.KEYS, posts[i].LOCATION)
	}

	renderJSON(w, &locations)
}