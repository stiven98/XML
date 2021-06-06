package service

import (
	"github.com/google/uuid"
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
func (service *SystemUsersService) GetAllUsernames() []string {
	return  service.Repo.GetAllUsernames()
}
func (service *SystemUsersService) GetUserId(username string) uuid.UUID {
	return  service.Repo.GetUserId(username);
}