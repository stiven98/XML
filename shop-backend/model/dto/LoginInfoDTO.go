package dto

import "github.com/google/uuid"

type LoginInfoDTO struct {
	UserID   uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
