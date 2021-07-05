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

type SavedPost struct {
	USERID uuid.UUID `json:"userid"`
	POSTID uuid.UUID `json:"postid"`
	OWNERID uuid.UUID `json:"ownerid"`
	COLLECTION PostCollection `json:"collection"`
}

type PostCollection struct {
	NAME string `json:"name"`
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


type AD struct {
	ID uuid.UUID `json:"id"`
	PATH string `json:"path"`
	LINK string `json:"link"'`
}

type Campaign struct {
	ID uuid.UUID `json:"id"`
	USERID uuid.UUID `json:"userid"`
	INFLUENCERS []uuid.UUID `json:"influencers"`
	ADS []AD `json:"ads"`
	TYPE string `json:"type"`
	DESCRIPTION string `json:"description"`
	ISMULTIPLE bool `json:"ismultiple,omitempty"`
	STARTDAY time.Time `json:"startday"`
	ENDDAY 	time.Time `json:"endday"`
	TIMESTOPLACE int `json:"timestoplace"`
	WHENTOPLACE string `json:"whentoplace"`
	COMMENTS []Comment `json:"comments"`
	LIKES []Like `json:"likes"`
	DISLIKES []Dislike `json:"dislikes"`
	TIMESPLACED int `json:"timesplaced"`
	TIMESCLICKED int `json:"timesclicked"`
	SHOWTOMEN	bool	`json:"showtomen"`
	SHOWTOWOMEN bool `json:"showtowomen"`
	SHOWUNDER18 bool `json:"showunder18"`
	SHOW18TO24 bool `json:"show18to24"`
	SHOW24TO35 bool	`json:"show24to35"`
	SHOWOVER35 bool `json:"showover35"`
}

func (post *Post) BeforeCreate() error {
	post.ID = uuid.New()
	return nil
}
