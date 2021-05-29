package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type AdministratorsService struct {
	AdministratorRepo *repository.AdministratorRepository
	SystemUserRepo *repository.SystemUsersRepository
}

func (service AdministratorsService) Update(m *model.Administrator) error {
	err := service.SystemUserRepo.Update(&m.SystemUser)
	if err != nil {
		return err
	}
	err = service.AdministratorRepo.Update(m)
	if err != nil {
		return err
	}
	return nil
}


