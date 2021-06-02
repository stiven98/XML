package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountCategory string
const (
	INFLUENCER AccountCategory = "influencer"
	SPORTS AccountCategory = "sports"
	NEWS_MEDIA	AccountCategory = "newsMedia"
	BUSINESS	AccountCategory = "buisness"
	BRAND	AccountCategory = "brand"
	ORGANIZATION	AccountCategory = "organization"
)

type AccountVerificationRequest struct {
	ID   uuid.UUID `json:"id"`
	FIRSTNAME string    `json:"firstName" gorm:"not null"`
	LASTNAME string    `json:"lastName" gorm:"not null"`
	PHOTOPATH string `json:"photopath" gorm:"not null"`
	CATEGORY AccountCategory `json:"category" gorm:"not null"`
	STATUS RequestStatus `json:"status" gorm:"not null"`
}
func (request *AccountVerificationRequest) BeforeCreate(scope *gorm.DB) error {
	request.ID = uuid.New()
	return nil
}