package model

import (
	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID uuid.UUID `json:"userID"`
	User User `gorm:"foreignKey:UserID"`
	Name string `json:"name" gorm:"not null"`
	Price float64 `json:"price" gorm:"not null"`
	PicturePath string `json:"picturePath"`
	Quantity int `json:"quantity" gorm:"not null"`
	Deleted bool `json:"deleted" gorm:"not null"`
}
