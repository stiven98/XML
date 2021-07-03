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
	"shop-backend/model"
	"shop-backend/service"
	"strings"
)

type ProductsHandler struct {
	 ProductsService *service.ProductsService
}

func (h ProductsHandler) GetProductByUser(writer http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)
	res := h.ProductsService.GetProductsByUser(vars["id"])
	writer.WriteHeader(http.StatusOK)
	renderJSON(writer, &res)

}

func (h ProductsHandler) GetAllProducts(writer http.ResponseWriter, request *http.Request) {
	res := h.ProductsService.GetAllProducts()
	writer.WriteHeader(http.StatusOK)
	renderJSON(writer, &res)

}

func (h ProductsHandler) ImageUpload(writer http.ResponseWriter, request *http.Request) {

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
			filePath = filepath.Join("products_images", resourceName)
		} else {
			filePath = filepath.Join("products_images", fh.Filename)
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

		defer f.Close()
		writer.WriteHeader(http.StatusCreated)
		writer.Header().Set("Content-Type", "application/json")
		renderJSON(writer, &resourceName)

	}
}

func (h ProductsHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var product model.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	product.ID = uuid.New()
	product.Deleted = false
	fmt.Println(product)

	err = h.ProductsService.Create(product)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

func (h ProductsHandler) GetProductById(writer http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)

	res, err := h.ProductsService.GetProductById(vars["id"])
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	renderJSON(writer, &res)
}

func (h ProductsHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var product model.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductsService.Update(product)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)

}

func (h ProductsHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)

	err := h.ProductsService.Delete(vars["id"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
