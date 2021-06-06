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
	POST  PostType = "post"
	ALBUM PostType = "album"
)
type Like struct {
	UserID uuid.UUID `json:"userid"`
}

type Dislike struct {
	UserID uuid.UUID `json:"userid"`
}


type Post struct {
	ID uuid.UUID 	`json:"id"`
	USERNAME string `json:"username"`
	TIMESTAMP time.Time `json:"timestamp"`
	ITEMS []PostItem `json:"items"`
	DESCRIPTION string `json:"description"`
	LOCATION string `json:"location"`
	HASHTAG string `json:"hashtag"`
	COMMENTS []Comment `json:"comments"`
	TYPE string `json:"type"`
	LIKES []Like `json:"likes"`
	DISLIKES []Dislike `json:"dislikes"`
}


func (post *Post) BeforeCreate() error {
	post.ID = uuid.New()
	return nil
}
