package model

import (
	"github.com/google/uuid"
	"net"
)

type AdType string
const (
	VIDEO AdType = "video"
	PHOTO AdType = "photo"
)

type Advertisment struct {
	ID uuid.UUID `json:"id"`
	PATH string `json:"path"`
	TYPE AdType `json:"type"`
	LINK string `json:"link"`
}

func (ad *Advertisment) BeforeCreate() net.Error {
	ad.ID = uuid.New()
	return nil
}