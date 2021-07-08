package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func (repo *ConversationRepository)GetConversation(user1 string, user2 string) model.Conversation {
	var ret model.Conversation
	collection := repo.Database.Collection("conversations")
	ctx := context.TODO()
	cursor, _ := collection.Find(ctx, bson.M{})


	for cursor.Next(ctx) {
		var conversation model.Conversation
		cursor.Decode(&conversation)
		if conversation.FirstUser.String() == user1 && conversation.SecondUser.String() == user2 {
			return conversation
		}
		if conversation.FirstUser.String() == user2 && conversation.SecondUser.String() == user1 {
			return conversation
		}
	}
	ret.ID = uuid.New()
	ret.FirstUser = uuid.MustParse(user1)
	ret.SecondUser = uuid.MustParse(user2)
	ret.Messages = [] model.Message {}
	collection.InsertOne(ctx, ret)
	return ret

}

func (repo *ConversationRepository) Update(ret model.Conversation) {
	collection := repo.Database.Collection("conversations")

	// find the document for which the _id field matches id and set the email to "newemail@example.com"
	// specify the Upsert option to insert a new document if a document matching the filter isn't found
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", ret.ID}}
	update := bson.D{{"$set", bson.D{{"messages", ret.Messages}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return
	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
}
