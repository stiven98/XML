package model

import "github.com/google/uuid"

type Order struct {
	OrderID uuid.UUID `json:"orderID" gorm:"primaryKey"`
	UserID uuid.UUID `json:"userID"`
	User User `gorm:"foreignKey:UserID"`
	ProductID uuid.UUID `json:"productID"`
	Product Product `gorm:"foreignKey:ProductID"`
	Quantity int `json:"quantity" gorm:"not null"`
}
