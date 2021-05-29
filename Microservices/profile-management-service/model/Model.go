package model

import "github.com/google/uuid"

type MutedUsers struct {
	//ID uuid.UUID `json:"id"`
	MutedByID uuid.UUID `json:"muted_by_id" gorm:"primaryKey"`
	MutedID uuid.UUID `json:"muted_id" gorm:"primaryKey"`
}

type CloseFriends struct {
	//ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id" gorm:"primaryKey"`
	FriendID uuid.UUID `json:"muted_id" gorm:"primaryKey"`
}

type BlockedUsers struct {
	//ID uuid.UUID `json:"id"`
	BlockedByID uuid.UUID `json:"blocked_by_id" gorm:"primaryKey"`
	BlockedID uuid.UUID `json:"blocked_id" gorm:"primaryKey"`
}



