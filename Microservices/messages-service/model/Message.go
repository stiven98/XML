package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Message struct {
	ID uuid.UUID `json:"id"`
	USERID uuid.UUID `json:"userid"`
	TIMESTAMP time.Time `json:"timestamp"`
	VALUE string `json:"value"`
}


func (message *Message) BeforeCreate(scope *mongo.Database) error {
	message.ID = uuid.New()
	return nil
}