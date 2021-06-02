package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)


type Agent struct {
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

type AgentRegistrationRequest struct {
	ID   uuid.UUID `json:"id"`
	LINK string    `json:"link" gorm:"not null"`
	USER string    `json:"agent" gorm:"not null"`
	STATUS RequestStatus `json:"status" gorm:"not null"`
}

func (request *AgentRegistrationRequest) BeforeCreate(scope *gorm.DB) error  {
	request.ID = uuid.New()
	return nil
}