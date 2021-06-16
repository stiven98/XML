package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"profileservice/model"
	"profileservice/model/Dto"
	"profileservice/service"
	"strings"
)

type UsersHandler struct {
	Service *service.UsersService
}

func (handler UsersHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.Service.Update(&user)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusExpectationFailed)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.SystemUser.TypeOfUser = model.USER
	fmt.Println(user)
	err = handler.Service.Create(&user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) GetById(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	user, _ :=handler.Service.GetById(vars["id"])
	renderJSON(w, &user)
}

func (handler *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request){
	users:=handler.Service.GetAll()
	renderJSON(w, &users)
}
func (handler *UsersHandler) ChangeWhetherIsPublic(w http.ResponseWriter, r *http.Request) {
	var dto Dto.ChangeWhetherIsPublicDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.ChangeWhetherIsPublic(dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UsersHandler) ChangeAllowedTags(w http.ResponseWriter, r *http.Request) {
	var dto Dto.ChangeAllowedTagsDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.Service.ChangeAllowedTags(dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler UsersHandler) IsPublic(writer http.ResponseWriter, request *http.Request) {
	tokens := strings.Split(request.URL.Path, "/")
	ID := tokens[int(len(tokens))-1]

	fmt.Println("ID")
	user, err := handler.Service.GetById(ID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(user)
	res := Dto.IsUserPublicDTO {
		ID: user.UserID,
		IsPublic: user.IsPublic,
	}
	renderJSON(writer, res)
	writer.WriteHeader(http.StatusOK)
}

func (handler *UsersHandler) GetPublicUsersIds(w http.ResponseWriter, r *http.Request){
	var ids Dto.PublicUsersIdsDto
	users := handler.Service.GetAllPublic()
	for i := range users {
		ids.KEYS = append(ids.KEYS, users[i].UserID.String())
	}
	renderJSON(w, &ids)
}



func (handler *UsersHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	var name string
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
			filePath = filepath.Join("profile_picture", resourceName)
		}else{
			filePath = filepath.Join("profile_picture", fh.Filename)
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
		name = resourceName
		//var postItem model.PostItem
		//postItem.PATH = resourceName
		//postItem.ID = uuid.New()
		//postItems = append(postItems, postItem)
		defer f.Close()

	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	renderJSON(w,name)
}