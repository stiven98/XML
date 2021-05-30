package Dto

import "github.com/google/uuid"

type ChangeAllowedTagsDto struct {
	USERID uuid.UUID `json:"userid"`
	FLAG bool	`json:"flag"`
}