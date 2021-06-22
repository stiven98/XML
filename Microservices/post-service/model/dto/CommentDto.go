package dto

import "github.com/google/uuid"

type CommentDto struct {
	USERID uuid.UUID `json:"userid"`
	POSTID uuid.UUID `json:"postid"`
	OWNERID uuid.UUID `json:"ownerid"`
	COMMENT string `json:"comment"`

}