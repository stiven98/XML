package repository

import (
	"agent-service/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignRepository struct {
	Database *mongo.Database
}

func (repo *CampaignRepository) Create(campaign *model.Campaign) error {
	collection := repo.Database.Collection("campaigns")
	ctx := context.TODO()
	result, _ := collection.InsertOne(ctx, campaign)
	fmt.Println(result)
	return nil
}
func(repo *CampaignRepository) GetAll() []model.Campaign{
	var campaigns []model.Campaign
	collection := repo.Database.Collection("campaigns")
	ctx := context.TODO()
	cursor, _ := collection.Find(ctx, bson.M{})

	for cursor.Next(ctx) {
		var campaign model.Campaign
		cursor.Decode(&campaign)
		campaigns = append(campaigns, campaign)
	}
	return campaigns
}
func(repo *CampaignRepository) Delete(id string) {
	collection := repo.Database.Collection("campaigns")
	ctx := context.TODO()
	fmt.Println(id)
	collection.DeleteOne(ctx, bson.M{"timesclicked":"4"})
}

