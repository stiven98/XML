package model

import "github.com/google/uuid"

type VerificationRequest struct {
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	DocumentPath string `json:"document_path"`
	Status RequestStatus `json:"status" gorm:"not null"`
}

