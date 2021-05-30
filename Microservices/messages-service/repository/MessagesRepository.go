package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"profileservice/model"
)

type MessagesRepository struct {
	Database *mongo.Database
}

func (repo *MessagesRepository) Create(message *model.Message) error {
	collection := repo.Database.Collection("messages")
	ctx := context.TODO()
	result, _ := collection.InsertOne(ctx, message)
	fmt.Println(result)
	return nil
}
func(repo *MessagesRepository) GetAll() []model.Message{
	var messages []model.Message
	collection := repo.Database.Collection("messages")
	ctx := context.TODO()
	cursor, _ := collection.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var message model.Message
		cursor.Decode(&message)
		messages = append(messages, message)
	}
	return messages
}

