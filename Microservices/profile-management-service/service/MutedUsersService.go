package service

import (
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