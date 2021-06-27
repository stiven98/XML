package model

import "github.com/google/uuid"

type Notify struct {
	ID   uuid.UUID `json:"id" gorm:"primaryKey"`
	UserID uuid.UUID `json:"userId"`
	NotifyUserID uuid.UUID `json:"notify_user_id"`
	TypeOfNotify string `json:"type_of_notify"` //post lik story dislike comments
	NotifyId uuid.UUID `json:"notify_id"` // id post story
}
