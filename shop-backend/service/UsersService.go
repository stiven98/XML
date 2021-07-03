package service

import (
	"shop-backend/model"
	"shop-backend/model/dto"
	"shop-backend/repository"
)

type UsersService struct {
	UsersRepository *repository.UsersRepository
}

func (s UsersService) Create(user model.User) error {
	return s.UsersRepository.Create(user)
}

func (s UsersService) Login(info dto.LoginInfoDTO) ( model.User, int64) {
	return s.UsersRepository.Login(info)
}