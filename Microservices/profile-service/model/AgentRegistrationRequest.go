package model

import (
	"github.com/google/uuid"
	"time"
)

type AgentRegistrationRequest struct {
	ID   uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID   uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName" gorm:"not null"`
	LastName string    `json:"lastName" gorm:"not null"`
	Username string    `json:"username" gorm:"unique; not null"`
	Email string    `json:"email" gorm:"unique; not null"`
	Password string    `json:"password" gorm:"not null"`
	Gender Gender 	`json:"gender" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
	TypeOfUser TypeOfUser `json:"type_of_user" gorm:"not null"`
	DateOfBirth time.Time  `json:"dateOfBirth" gorm:"not null"`
	PicturePath string  `json:"picturePath" gorm:"not null"`
	WebsiteLink string `json:"websiteLink" gorm:"not null"`
	IsApproved bool `json:"isApproved" gorm:"not null"`
}