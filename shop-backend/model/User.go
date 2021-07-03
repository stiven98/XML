package model

import "github.com/google/uuid"

type User struct {
	UserID   uuid.UUID `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName" gorm:"not null"`
	LastName string    `json:"lastName" gorm:"not null"`
	Username string    `json:"username" gorm:"unique; not null"`
	Email string    `json:"email" gorm:"unique; not null"`
	Password string    `json:"password" gorm:"not null"`
}
