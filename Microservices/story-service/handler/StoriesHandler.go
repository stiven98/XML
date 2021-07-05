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
	"strconv"
	"strings"
	"time"
)
type StoriesHandler struct {
	Service *service.StoriesService
}

func (handler *StoriesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var storyDto dto.NewStory
	err := json.NewDecoder(r.Body).Decode(&storyDto)
	story := model.Story{
		ID:        uuid.New(),
		USERID:    storyDto.USERID,
		TIMESTAMP: time.Now(),
		ITEMS:     storyDto.ITEMS,
		LOCATION:  storyDto.LOCATION,
		HASHTAG:   storyDto.HASHTAG,
		TYPE:      storyDto.TYPE,
	}

	err = handler.Service.Create(&story)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("\n\n\n\n\naaaaaaaaaaaaaaa\n\n\n\n")
	fmt.Println(storyDto.CLOSEFRIENDS)
	fmt.Println("\n\n\n\n\naaaaaaaaaaaaaaa\n\n\n\n")
	if storyDto.CLOSEFRIENDS {
		restCF, errCF := http.Get("http://localhost:8087/users/closeFriend/" +  story.USERID.String())
		if errCF != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var closeFriends []string
		err = json.NewDecoder(restCF.Body).Decode(&closeFriends)
		fmt.Println(closeFriends)
		err = handler.Service.AddStoryToFeed(closeFriends, &story)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
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
			return
		}
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

func (handler *StoriesHandler) GetPagedFeed(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := query["id"]
	pageNumber, _ := query["pageNumber"]
	neededResults, _ := query["neededResults"]
	var stories []model.Story
	neededResultsInt := 0
	pageNumberInt := 0
	if len(id) > 0 && len(pageNumber) > 0 && len(neededResults) > 0 {
		 stories = handler.Service.GetFeed(id[0])
		neededResultsInt, _ = strconv.Atoi(neededResults[0])
		pageNumberInt, _ = strconv.Atoi(pageNumber[0])
	}
	var page model.Page
	var neededStories []model.Story
	firstIndex := (pageNumberInt - 1) * neededResultsInt
	lastIndex := firstIndex + neededResultsInt
	for i := firstIndex; i < lastIndex; i++{
		if len(stories) > i {
			neededStories = append(neededStories, stories[i])
		}
	}
	page.Stories = neededStories
	page.TotalCount = len(stories)
	renderJSON(w, &page)
}

func (handler *StoriesHandler) GetMyPagedStories(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := query["id"]
	pageNumber, _ := query["pageNumber"]
	neededResults, _ := query["neededResults"]
	var stories []model.Story
	neededResultsInt := 0
	pageNumberInt := 0
	if len(id) > 0 && len(pageNumber) > 0 && len(neededResults) > 0 {
		stories = handler.Service.GetMyStories(id[0])
		neededResultsInt, _ = strconv.Atoi(neededResults[0])
		pageNumberInt, _ = strconv.Atoi(pageNumber[0])
	}
	var page model.Page
	var neededStories []model.Story
	firstIndex := (pageNumberInt - 1) * neededResultsInt
	lastIndex := firstIndex + neededResultsInt
	for i := firstIndex; i < lastIndex; i++{
		if len(stories) > i {
			neededStories = append(neededStories, stories[i])
		}
	}
	page.Stories = neededStories
	page.TotalCount = len(stories)
	renderJSON(w, &page)
}

func (handler *StoriesHandler) AddToHighlights(writer http.ResponseWriter, request *http.Request) {
	var highlightDto model.Highlight
	err := json.NewDecoder(request.Body).Decode(&highlightDto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.AddToHighlights(highlightDto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	return
}

func (handler *StoriesHandler) RemoveFromHighlights(writer http.ResponseWriter, request *http.Request) {
	var highlightDto model.Highlight
	err := json.NewDecoder(request.Body).Decode(&highlightDto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.RemoveFromHighlights(highlightDto)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	return
}

func (handler *StoriesHandler) GetPagedHighlights(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id, _ := query["id"]
	pageNumber, _ := query["pageNumber"]
	neededResults, _ := query["neededResults"]
	var stories []model.Story
	neededResultsInt := 0
	pageNumberInt := 0
	if len(id) > 0 && len(pageNumber) > 0 && len(neededResults) > 0 {
		stories = handler.Service.GetHighlights(id[0])
		neededResultsInt, _ = strconv.Atoi(neededResults[0])
		pageNumberInt, _ = strconv.Atoi(pageNumber[0])
	}
	var page model.Page
	var neededStories []model.Story
	firstIndex := (pageNumberInt - 1) * neededResultsInt
	lastIndex := firstIndex + neededResultsInt
	for i := firstIndex; i < lastIndex; i++{
		if len(stories) > i {
			neededStories = append(neededStories, stories[i])
		}
	}
	page.Stories = neededStories
	page.TotalCount = len(stories)
	renderJSON(w, &page)
}