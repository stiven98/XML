package service

import (
	_ "github.com/google/uuid"
	"profileservice/model"
	"profileservice/model/Dto"
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
	return service.UsersRepo.Create(user)
}

func (service *UsersService) GetAll() []model.User {
	return  service.UsersRepo.GetAll()
}

func (service *UsersService) ChangeWhetherIsPublic(dto Dto.ChangeWhetherIsPublicDto) error {
	return  service.UsersRepo.ChangeWhetherIsPublic(&dto)
}
func (service *UsersService) ChangeAllowedTags(dto Dto.ChangeAllowedTagsDto) error {
	return  service.UsersRepo.ChangeAllowedTags(&dto)
}

func (service UsersService) GetById(id string) (model.User,error) {
	return service.UsersRepo.GetById(id)

}




