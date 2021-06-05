package service

import (
	"admin-service/model"
	"admin-service/repository"
)

type AccountVerificationsService struct {
	AccountVerificationsRepo *repository.AccountVerificationsRepository
}

func (service AccountVerificationsService) Update(request *model.AccountVerificationRequest) error {
	err := service.AccountVerificationsRepo.Update(request)
	if err != nil {
		return err
	}
	return nil
}
func (service *AccountVerificationsService) Create(request *model.AccountVerificationRequest) error {
	service.AccountVerificationsRepo.Create(request)
	return nil
}

func (service *AccountVerificationsService) GetAll() []model.AccountVerificationRequest {
	return  service.AccountVerificationsRepo.GetAll()
}
