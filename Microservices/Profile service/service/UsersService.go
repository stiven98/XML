package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type UsersService struct {
	Repo *repository.UsersRepository
}

func (service *UsersService) Create(user *model.SystemUser) error {
	service.Repo.Create(user)
	return nil
}
func (service *UsersService) Update(user *model.SystemUser) error {
	service.Repo.Update(user)
	return nil
}
func (service *UsersService) GetAll() []model.SystemUser {
	return  service.Repo.GetAll()
}