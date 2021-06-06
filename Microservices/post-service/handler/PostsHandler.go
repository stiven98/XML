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
	err = handler.Service.Create(&post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
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
		if(len(fileName) >= 2){
			filePath = filepath.Join("user_posts", uuid.NewString() +  "." + fileName[1])
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
		postItem.PATH = filePath
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
	post :=handler.Service.GetByKey(vars["key"])

	renderJSON(w, &post)
}
