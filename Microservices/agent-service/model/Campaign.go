package model

import (
	"github.com/google/uuid"
	"runtime"
	"time"
)

type CampaignType string
const (
	STORY CampaignType = "story"
	POST CampaignType = "post"
)
type Comment struct {
	ID uuid.UUID `json:"id"`
	USERID uuid.UUID `json:"userid"`
	TIMESTAMP time.Time `json:"timestamp"`
	VALUE string 	`json:"value"`
}
type Like struct {
	UserID uuid.UUID `json:"userid"`
}

type Dislike struct {
	UserID uuid.UUID `json:"userid"`
}

type Campaign struct {
	ID uuid.UUID `json:"id"`
	AGENTID uuid.UUID `json:"agentid"`
	ADLIST []Advertisment `json:"adlist"`
	COMMENTS []Comment	`json:"comments"`
	LIKES []Like	`json:"likes"`
	DISLIKES []Dislike	`json:"dislikes"`
	TIMESPLACED int 	`json:"timesplaced"`
	TIMESCLICKED int	`json:"timesclicked"`
	TYPE CampaignType	`json:"type"`
}

type SingleCampaign struct {
	CAMPAIGN Campaign `json:"campaign"`
	TIMESTAMP time.Time `json:"timestamp"`
}

type RepetableCampaign struct {
	CAMPAIGN Campaign	`json:"campaign"`
	STARTDATE time.Time	`json:"startdate"`
	ENDDATE time.Time 	`json:"enddate"`
	DAILYQOTA int 		`json:"dailyqota"`
}

func (campaign *Campaign) BeforeCreate() runtime.Error {
	campaign.ID = uuid.New()
	return nil
}