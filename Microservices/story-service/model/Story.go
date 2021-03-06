package model

import (
	"github.com/google/uuid"
	"time"
)

type StoryType string
const (
	STORY StoryType = "story"
	ALBUM StoryType = "album"
)

type StoryItem struct {
	ID uuid.UUID `json:"itemid"`
	PATH string `json:"path"`
}

type Feed struct {
	UserId uuid.UUID `json:"userid"`
	StoryId uuid.UUID `json:"storyid"`
}

type Highlight struct {
	UserId uuid.UUID `json:"userid"`
	StoryId uuid.UUID `json:"storyid"`
}

type Page struct {
	Stories []Story `json:"stories"`
	TotalCount int `json:"total_count"`
}

type Story struct {
	ID uuid.UUID `json:"id"`
	USERID uuid.UUID `json:"userid"`
	TIMESTAMP time.Time `json:"timestamp"`
	ITEMS []StoryItem `json:"items"`
	LOCATION string `json:"location"`
	HASHTAG string `json:"hashtag"`
	TYPE string `json:"type"`
}

func (story *Story) BeforeCreate() error {
	story.ID = uuid.New()
	return nil
}
