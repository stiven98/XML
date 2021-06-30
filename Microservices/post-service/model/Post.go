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

type Feed struct {
	UserId uuid.UUID `json:"userid"`
	PostId uuid.UUID `json:"postid"`
}
type ReportedBy struct {
	UserId uuid.UUID `json:"userid"`
}
type Post struct {
	ID uuid.UUID 	`json:"id"`
	USERID uuid.UUID `json:"userid"`
	TIMESTAMP time.Time `json:"timestamp"`
	ITEMS []PostItem `json:"items"`
	DESCRIPTION string `json:"description"`
	LOCATION string `json:"location"`
	HASHTAG string `json:"hashtag"`
	COMMENTS []Comment `json:"comments"`
	TYPE string `json:"type"`
	LIKES []Like `json:"likes"`
	DISLIKES []Dislike `json:"dislikes"`
	REPORTS []ReportedBy `json:"reports"`
}
type Campaign struct {
	ID uuid.UUID `json:"id"`
	ITEMS []PostItem `json:"posts"`
	DESCRIPTION string `json:"description"`
	WEBSITE string `json:"website"`
	ISMULTIPLE bool `json:"ismultiple"`
	STARTDATE time.Time `json:"startdate"`
	ENDDATE time.Time `json:"enddate"`
	TIMESTOPLACE int `json:"timestoplace"`
	TIMETOSHOW string `json:"timetoshow"`
	TARGETGROUP []string	`json:"targetgroup"`
	LIKES []Like `json:"likes"`
	DISLIKES []Dislike `json:"dislikes"`
	COMMENTS []Comment `json:"comments"`
}

func (post *Post) BeforeCreate() error {
	post.ID = uuid.New()
	return nil
}
