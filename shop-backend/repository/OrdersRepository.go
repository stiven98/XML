package repository

import (
	"gorm.io/gorm"
	"shop-backend/model"
)

type OrdersRepository struct {
	Database *gorm.DB
}

func (r OrdersRepository) GetOrdersByUserId(id string) []model.Order {
	var orders [] model.Order
	r.Database.Model(&model.Order{}).Where("user_id = ?", id).Find(&orders)
	return orders
}

func (r OrdersRepository) Create(order model.Order) error {
	return r.Database.Model(&model.Order{}).Create(order).Error
}



