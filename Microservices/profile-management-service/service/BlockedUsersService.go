package service

import (
	"profile-management-service/model"
	"profile-management-service/repository"
)

type BlockedUsersService struct {
	BlockedUsersRepository *repository.BlockedUsersRepository
}

func (s BlockedUsersService) BlockUserByUser(b *model.BlockedUsers) error {
	return s.BlockedUsersRepository.BlockUserByUser(b)
}