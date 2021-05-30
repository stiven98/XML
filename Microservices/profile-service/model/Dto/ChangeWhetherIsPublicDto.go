package Dto

import "github.com/google/uuid"

type ChangeWhetherIsPublicDto struct {
	USERID uuid.UUID `json:"userid"`
	FLAG bool	`json:"flag"`
}
