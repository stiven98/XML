package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"shop-backend/model"
	"shop-backend/service"
)

type OrdersHandler struct {
	OrdersService *service.OrdersService
	ProductsService *service.ProductsService
}

func (h OrdersHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var order model.Order
	err := json.NewDecoder(request.Body).Decode(&order)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductsService.GetProductById(order.ProductID.String())
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.Quantity < order.Quantity {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	product.Quantity = product.Quantity - order.Quantity
	h.ProductsService.Update(product)

	order.OrderID = uuid.New()

	err = h.OrdersService.Create(order)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)


}

func (h OrdersHandler) GetOrdersByUserId(writer http.ResponseWriter, request *http.Request) {
	vars :=mux.Vars(request)

	res := h.OrdersService.GetOrdersByUserId(vars["id"])

	for i := range res {
		response, _ := h.ProductsService.GetProductById(res[i].ProductID.String())
		res[i].Product = response
	}

	fmt.Println(res)

	writer.WriteHeader(http.StatusOK)
	renderJSON(writer, &res)
}
