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
	"storyservice/model"
	"storyservice/model/dto"
	"storyservice/service"
	"strings"
	"time"
)
type StoriesHandler struct {
	Service *service.StoriesService
}

func (handler *StoriesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var story model.Story
	fmt.Println(json.NewDecoder(r.Body).Decode(&story))
	err := json.NewDecoder(r.Body).Decode(&story)
	story.ID = uuid.New()
	story.TIMESTAMP = time.Now()
	err = handler.Service.Create(&story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	rest, err := http.Get("http://localhost:8088/users/getFollowers/" + story.USERID.String())
	//rest, err := http.Get("https://mocki.io/v1/84324533-ee57-4eb2-8042-3f5845dcc41b")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dto dto.FollowersDto
	err = json.NewDecoder(rest.Body).Decode(&dto)
	fmt.Println(dto.KEYS)
	err = handler.Service.AddStoryToFeed(dto.KEYS, &story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}


func (handler *StoriesHandler) GetByKey(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Println(vars["key"])
	story:=handler.Service.GetByKey(vars["key"])

	renderJSON(w, &story)
}

func (handler *StoriesHandler) UploadFile(writer http.ResponseWriter, request *http.Request) {
	var storyItems[] model.StoryItem
	// Maximum upload of 10 MB files
	request.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	fhs := request.MultipartForm.File["files"]
	var i int
	i = 0
	for _, fh := range fhs {
		i = i + 1
		f, err := fh.Open()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
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
			filePath = filepath.Join("user_stories", resourceName)
		}else{
			filePath = filepath.Join("user_stories", fh.Filename)
		}
		dst, err := os.Create(filePath)
		defer dst.Close()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, f); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		var storyItem model.StoryItem
		storyItem.PATH = resourceName
		storyItem.ID = uuid.New()
		storyItems = append(storyItems, storyItem)
		defer f.Close()

	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	renderJSON(writer, &storyItems)
}

func (handler *StoriesHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	post :=handler.Service.GetFeed(vars["id"])
	renderJSON(w, &post)
}

func (handler *StoriesHandler) GetMyStories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	post :=handler.Service.GetMyStories(vars["id"])
	renderJSON(w, &post)
}
