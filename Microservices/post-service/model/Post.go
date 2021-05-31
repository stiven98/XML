package model

import (
	"github.com/google/uuid"
	"time"
)

type PostItem struct {
	ID uuid.UUID `json:"id"`
	PATH string `json:"path"`
}
type Favourites struct {
	USERID uuid.UUID `json:"userid"`
	POSTID uuid.UUID 	`json:"postid"`
}
type PostType string
const (
	STORY PostType = "story"
	ALBUM PostType = "album"
)

type Post struct {
	ID uuid.UUID 	`json:"id"`
	USERID uuid.UUID `json:"userid"`
	TIMESTAMP time.Time `json:"timestamp"`
	ITEMS []PostItem `json:"items"`
	LOCATION string `json:"location"`
	HASHTAG string `json:"hashtag"`
	COMMENTS []Comment `json:"comments"`
	TYPE string `json:"type"`
}

func (post *Post) BeforeCreate() error {
	post.ID = uuid.New()
	return nil
}
