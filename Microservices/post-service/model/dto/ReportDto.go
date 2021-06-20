package dto

import "github.com/google/uuid"

type ReportDto struct {
	POSTID uuid.UUID `json:"postid"`
	USERID uuid.UUID `json:"userid"`
	OWNERID uuid.UUID `json:"ownerid"`
}
