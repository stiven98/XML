package handler

import (
	"admin-service/model"
	"admin-service/service"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type VerificationRequestHandler struct {
	VerificationRequestService *service.VerificationRequestService
}

func (h VerificationRequestHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	requests:= h.VerificationRequestService.GetAll()

	for i := range requests {
		fmt.Println(requests[i].UserID)

	}
	fmt.Print(requests)

	renderJSON(writer, &requests)
}

func (h VerificationRequestHandler) CreateVerificationRequest(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(32 << 20)
	fhs := request.MultipartForm.File["document"]
	i := 0
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
			filePath = filepath.Join("documents", resourceName)
		}else{
			filePath = filepath.Join("documents", fh.Filename)
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
		var verificationRequest model.VerificationRequest
		verificationRequest.DocumentPath = resourceName
		verificationRequest.Status = model.SUBMITTED
		verificationRequest.ID = uuid.New()
		defer f.Close()

		tokens := strings.Split(request.URL.Path, "/")
		ID := tokens[int(len(tokens))-1]
		verificationRequest.UserID = uuid.MustParse(ID)

		err = h.VerificationRequestService.Create(&verificationRequest)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	writer.WriteHeader(http.StatusOK)

	fmt.Println(fhs)
}

func (h VerificationRequestHandler) Accept(writer http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)
	err := h.VerificationRequestService.Accept(uuid.MustParse(vars["id"]))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h VerificationRequestHandler) Decline(writer http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)
	err := h.VerificationRequestService.Decline(uuid.MustParse(vars["id"]))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusOK)
}


