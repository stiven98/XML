package service

import (
	"admin-service/model"
	"admin-service/repository"
)

type AdministratorsService struct {
	AdministratorRepo *repository.AdministratorsRepository
}

func (service AdministratorsService) Update(admin *model.Administrator) error {
	err := service.AdministratorRepo.Update(admin)
	if err != nil {
		return err
	}
	return nil
}
func (service *AdministratorsService) Create(admin *model.Administrator) error {
	service.AdministratorRepo.Create(admin)
	return nil
}

func (service *AdministratorsService) GetAll() []model.Administrator {
	return  service.AdministratorRepo.GetAll()
}


