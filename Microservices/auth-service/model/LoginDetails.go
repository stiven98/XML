package model

import "github.com/google/uuid"

type LoginDetails struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email" gorm:"unique; not null"`
	Password string    `json:"password" gorm:"not null"`
}