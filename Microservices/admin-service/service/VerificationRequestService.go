package service

import (
	"admin-service/model"
	"admin-service/repository"
	"github.com/google/uuid"
)

type VerificationRequestService struct {
	VerificationRequestRepository *repository.VerificationRequestRepository
}


func (service *VerificationRequestService) Create(request *model.VerificationRequest) error {
	return service.VerificationRequestRepository.Create(request)
}

func (service *VerificationRequestService) GetAll() []model.VerificationRequest {
	return  service.VerificationRequestRepository.GetAll()
}

func (service *VerificationRequestService) Accept(parse uuid.UUID) error {
	return service.VerificationRequestRepository.Accept(parse)
}

func (service *VerificationRequestService) Decline(parse uuid.UUID) error {
	return service.VerificationRequestRepository.Decline(parse)
}
