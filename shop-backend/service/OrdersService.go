package service

import (
	"shop-backend/model"
	"shop-backend/repository"
)

type OrdersService struct {
	OrdersRepository *repository.OrdersRepository
}

func (s OrdersService) GetOrdersByUserId(id string) [] model.Order {
	return s.OrdersRepository.GetOrdersByUserId(id)
}

func (s OrdersService) Create(order model.Order) error {
	return s.OrdersRepository.Create(order)
}
