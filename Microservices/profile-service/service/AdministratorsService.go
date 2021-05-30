package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type AdministratorsService struct {
	AdministratorRepo *repository.AdministratorsRepository
	SystemUserRepo *repository.SystemUsersRepository
}

func (service AdministratorsService) Update(admin *model.Administrator) error {
	err := service.SystemUserRepo.Update(&admin.SystemUser)
	if err != nil {
		return err
	}
	err = service.AdministratorRepo.Update(admin)
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


