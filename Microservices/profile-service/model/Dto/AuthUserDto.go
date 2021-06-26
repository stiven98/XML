package Dto

import "github.com/google/uuid"

type CreateUserDTO struct {
	ID uuid.UUID	`json:"id"`
	USERNAME string	`json:"username"`
	PASSWORD string	`json:"password"`
	ACTIVE bool `json:"active"`
	ROLE string	`json:"role"`
}