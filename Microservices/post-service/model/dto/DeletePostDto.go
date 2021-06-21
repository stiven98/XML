package dto

import "github.com/google/uuid"

type DeletePostDto struct {
	OWNERID uuid.UUID `json:"ownerId"`
	POSTID uuid.UUID `json:"postId"`
}