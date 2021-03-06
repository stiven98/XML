package service

import (
	"github.com/google/uuid"
	"profile-management-service/model"
	"profile-management-service/repository"
)

type BlockedUsersService struct {
	BlockedUsersRepository *repository.BlockedUsersRepository
}

func (service BlockedUsersService) BlockUserByUser(b *model.BlockedUsers) error {
	return service.BlockedUsersRepository.BlockUserByUser(b)
}

func (service BlockedUsersService) GetAllBlockedByUserId(id string) ([]uuid.UUID,error){
	var blocked []model.BlockedUsers
	blocked, err := service.BlockedUsersRepository.GetAllBlockedByUserId(id)

	var list []uuid.UUID
	for i:=range blocked {
		list = append(list, blocked[i].BlockedID)
	}

	return list , err
}

func (service BlockedUsersService) IsBlocked(b *model.BlockedUsers) (bool, error) {
	var blocked model.BlockedUsers
	blocked, err := service.BlockedUsersRepository.IsBlocked(b)

	if blocked.BlockedByID != b.BlockedByID{
		return false,err
	}
	return true,err
}

func (service BlockedUsersService) UnBlockUserByUser(b *model.BlockedUsers) error {
	return service.BlockedUsersRepository.UnBlockUserByUser(b)
}
