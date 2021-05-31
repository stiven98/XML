package service

import (
	"auth-service/model"
	"auth-service/repository"
)

type LoginDetailsService struct{
	LoginDetailsRepository *repository.LoginDetailsRepository
}

func (service *LoginDetailsService) Update(loginDetails *model.LoginDetails) error {
	err := service.LoginDetailsRepository.Update(loginDetails)
	if err != nil {
		return err
	}
	return nil
}
func (service *LoginDetailsService) Create(loginDetails *model.LoginDetails) error {
	service.LoginDetailsRepository.Create(loginDetails)
	return nil
}

func (service *LoginDetailsService) GetAll() []model.LoginDetails {
	return service.LoginDetailsRepository.GetAll()
}

func (service *LoginDetailsService) GetByEmail(email string) model.LoginDetails {
	return service.LoginDetailsRepository.GetByEmail(email)
}
