package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Gender string
const (
	male Gender = "male"
	female Gender = "female"
)

type UserType string
const (
	admin UserType = "admin"
	user UserType = "user"
)
type SystemUser struct {
	ID   uuid.UUID `json:"id"`
	FirstName string    `json:"firstName" gorm:"not null"`
	LastName string    `json:"lastName" gorm:"not null"`
	Username string    `json:"username" gorm:"unique; not null"`
	Email string    `json:"email" gorm:"unique; not null"`
	Password string    `json:"password" gorm:"not null"`
	Gender string 	`json:"gender" gorm:"not null"`
	DateOfBirth time.Time  `json:"dateOfBirth" gorm:"not null"`
}

type User struct {
	SystemUser SystemUser
	IsPublic bool `json:"isPublic" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
	WebSite string `json:"webSite" gorm:"not null"`
	Biography string `json:"biography"`
	AllowedTags bool `json:"allowedTags"`
	IsBlocked bool `json:"isBlocked"`
}
type Administrator struct {
	SystemUser SystemUser
}

type Agent struct {
	SystemUser SystemUser
}


//TODO check if needs to be done only for the first insert, not modifications
func (sysUser *SystemUser) BeforeCreate(scope *gorm.DB) error {
	sysUser.ID = uuid.New()
	return nil
}
