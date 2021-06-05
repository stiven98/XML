package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Gender string
const (
	MALE Gender = "male"
	FEMALE Gender = "female"
)
type TypeOfUser string
const (
	ADMIN TypeOfUser = "admin"
	USER TypeOfUser = "user"
	AGENT TypeOfUser = "agent"
)
type Administrator struct {
	ID          uuid.UUID  `json:"id"`
	FIRSTNAME   string     `json:"firstName" gorm:"not null"`
	LASTNAME    string     `json:"lastName" gorm:"not null"`
	USERNAME    string     `json:"username" gorm:"unique; not null"`
	EMAIL       string     `json:"email" gorm:"unique; not null"`
	PASSWORD    string     `json:"password" gorm:"not null"`
	GENDER      Gender     `json:"gender" gorm:"not null"`
	TYPEOFUSER  TypeOfUser `json:"type_of_user" gorm:"not null"`
	DATEOFBIRTH time.Time  `json:"dateOfBirth" gorm:"not null"`
}

func (admin *Administrator) BeforeCreate(scope *gorm.DB) error {
	admin.ID = uuid.New()
	return nil
}