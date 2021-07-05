package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Message struct {
	ID uuid.UUID `json:"id"`
	Sender uuid.UUID `json:"sender"`
	Receiver uuid.UUID `json:"receiver"`
	Timestamp time.Time `json:"timestamp"`
	Content string `json:"content"`
	Type string `json:"type"`
}


func (message *Message) BeforeCreate(scope *mongo.Database) error {
	message.ID = uuid.New()
	return nil
}