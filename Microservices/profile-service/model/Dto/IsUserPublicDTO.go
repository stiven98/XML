package Dto


import "github.com/google/uuid"

type IsUserPublicDTO struct {
	ID uuid.UUID `json:"id"`
	IsPublic bool	`json:"is_public"`
}
