package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type SystemUsersService struct {
	Repo *repository.SystemUsersRepository
}

func (service *SystemUsersService) Create(user *model.SystemUser) error {
	service.Repo.Create(user)
	return nil
}
func (service *SystemUsersService) Update(user *model.SystemUser) error {
	service.Repo.Update(user)
	return nil
}
func (service *SystemUsersService) GetAll() []model.SystemUser {
	return  service.Repo.GetAll()
}
