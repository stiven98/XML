package model

import (
	"github.com/google/uuid"
)

type Conversation struct {
	ID uuid.UUID               `json:"id"`
	FirstUser uuid.UUID        `json:"firstUser"`
	SecondUser uuid.UUID       `json:"secondUser"`
	Messages []Message 			`json:"messages"`
}

func (conversation *Conversation) BeforeCreate() error {
	conversation.ID = uuid.New()
	return nil
}