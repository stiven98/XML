package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"profileservice/model"
)

type ConversationRepository struct {
	Database *mongo.Database
}

func (repo *ConversationRepository) Create(conversation *model.Conversation) error {
	collection := repo.Database.Collection("conversations")
	ctx := context.TODO()
	result, _ := collection.InsertOne(ctx, conversation)
	fmt.Println(result)
	return nil
}
func(repo *ConversationRepository) GetAll() []model.Conversation{
	var conversations []model.Conversation
	collection := repo.Database.Collection("conversations")
	ctx := context.TODO()
	cursor, _ := collection.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var conversation model.Conversation
		cursor.Decode(&conversation)
		conversations = append(conversations, conversation)
	}
	return conversations
}
