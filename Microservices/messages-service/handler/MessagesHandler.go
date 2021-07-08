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
	"profileservice/model"
	"profileservice/service"
	"strings"
	"time"
)

type MessagesHandler struct {
	MessageService *service.MessagesService
	ConversationService *service.ConversationsService
}


func (handler *MessagesHandler) Add(w http.ResponseWriter, r *http.Request) {
	var message model.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	message.ID = uuid.New()
	message.Timestamp = time.Now()

	vars :=mux.Vars(r)
	ret := handler.ConversationService.GetConversation(vars["user1"], vars["user2"])


	ret.Messages = append(ret.Messages, message)

	handler.ConversationService.Update(ret)

	//err = handler..Create(&message)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	renderJSON(w, &ret)
}


func (handler *MessagesHandler) GetAll(w http.ResponseWriter, r *http.Request){
	conversations:= handler.MessageService.GetAll()
	renderJSON(w, &conversations)
}


func (handler *MessagesHandler) ImageUpload(writer http.ResponseWriter, request *http.Request) {

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
		if len(fileName) >= 2 {
			resourceName = uuid.NewString() +  "." + fileName[1]
			filePath = filepath.Join("images", resourceName)
		} else {
			filePath = filepath.Join("images", fh.Filename)
		}
		dst, err := os.Create(filePath)
		defer dst.Close()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf(filePath)

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, f); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()
		writer.WriteHeader(http.StatusCreated)
		writer.Header().Set("Content-Type", "application/json")
		renderJSON(writer, &resourceName)

	}
}