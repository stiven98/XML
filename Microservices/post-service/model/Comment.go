package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID uuid.UUID `json:"id"`
	USERID uuid.UUID `json:"userid"`
	TIMESTAMP time.Time `json:"timestamp"`
	VALUE string 	`json:"value"`
}

func (comment *Comment) BeforeCreate() error {
	comment.ID = uuid.New()
	return nil
}