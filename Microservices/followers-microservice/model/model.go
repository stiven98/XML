package model

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID
}

type Followers struct {
	KEYS[] string `json:"keys"`
}
