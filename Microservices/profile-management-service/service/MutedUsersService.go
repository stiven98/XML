package service

import (
	"github.com/google/uuid"
	"profile-management-service/model"
	"profile-management-service/repository"
)

type MutedUsersService struct {
	MutedUsersRepository *repository.MutedUsersRepository
}


func (service MutedUsersService) IsMuted(m *model.MutedUsers) (bool, error) {
	var muted model.MutedUsers
	muted, err := service.MutedUsersRepository.IsMuted(m)

	if muted.MutedByID != m.MutedByID{
		return false,err
	}
	return true,err
}

func (service MutedUsersService) MutedUserByUser(b *model.MutedUsers) error {
	return service.MutedUsersRepository.MutedUserByUser(b)
}

func (service MutedUsersService) UnMutedUserByUser(b *model.MutedUsers) error {
	return service.MutedUsersRepository.UnMutedUserByUser(b)
}

func (service MutedUsersService) GetAllMutedByUserId(id string) ([]uuid.UUID,error) {
	var muted []model.MutedUsers
	muted, err := service.MutedUsersRepository.GetAllMutedByUserId(id)

	var list []uuid.UUID
	for i:=range muted {
		list = append(list, muted[i].MutedID)
	}

	return list , err
}