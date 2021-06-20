package repository

import (
	"admin-service/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerificationRequestRepository struct {
	Database *gorm.DB
}

func (repo *VerificationRequestRepository) GetAll() [] model.VerificationRequest {
	var ret []model.VerificationRequest
	repo.Database.Model(model.VerificationRequest{}).Find(&ret)
	return ret
}

func (repo *VerificationRequestRepository) Create(request *model.VerificationRequest) error {
	result := repo.Database.Model(model.VerificationRequest{}).Create(request)
	return result.Error

}

func (repo *VerificationRequestRepository) Accept(id uuid.UUID) error {
	result := repo.Database.Model(model.VerificationRequest{}).Where("id = ?", id).Update("status", model.APPROVED)
	return result.Error
}

func (repo *VerificationRequestRepository) Decline(id uuid.UUID) error {
	result := repo.Database.Model(model.VerificationRequest{}).Where("id = ?", id).Update("status", model.DENIED)
	return result.Error
}