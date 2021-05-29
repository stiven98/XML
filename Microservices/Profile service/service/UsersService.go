package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type UsersService struct {
	UsersRepo *repository.UsersRepository
	SystemUserRepo *repository.SystemUsersRepository
}

func (service UsersService) Update(user *model.User) error {
	err := service.SystemUserRepo.Update(&user.SystemUser)
	if err != nil {
		return err
	}
	err = service.UsersRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}
func (service *UsersService) Create(user *model.User) error {
	service.UsersRepo.Create(user)
	return nil
}

func (service *UsersService) GetAll() []model.User {
	return  service.UsersRepo.GetAll()
}

