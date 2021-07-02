package dto

import "github.com/google/uuid"

type SavePostDto struct {
	USERID uuid.UUID `json:"userid"`
	POSTID uuid.UUID `json:"postid"`
	OWNERID uuid.UUID `json:"ownerid"`
}