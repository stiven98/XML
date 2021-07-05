package dto

import (
	"github.com/google/uuid"
	"storyservice/model"
	"time"
)

type NewStory struct {
	ID uuid.UUID `json:"id"`
	USERID uuid.UUID `json:"userid"`
	TIMESTAMP time.Time `json:"timestamp"`
	ITEMS []model.StoryItem `json:"items"`
	LOCATION string `json:"location"`
	HASHTAG string `json:"hashtag"`
	TYPE string `json:"type"`
	CLOSEFRIENDS bool `json:"closefriends"`
}
