package model

import (
	"github.com/google/uuid"
)

type Conversation struct {
	ID uuid.UUID `json:"id"`
	FIRSTUSER uuid.UUID `json:"firstuser"`
	SECONDUSER uuid.UUID `json:"seconduser"`
	MESSAGES []Message `json:"messages"`
}

func (conversation *Conversation) BeforeCreate() error {
	conversation.ID = uuid.New()
	return nil
}