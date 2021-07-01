package service

import (
	"github.com/google/uuid"
	"profile-management-service/model"
	"profile-management-service/repository"
)

type CloseFriendsService struct {
	CloseFriendsService *repository.CloseFriendsRepository
}

func (service CloseFriendsService) AddCloseFriend(friend *model.CloseFriends) error{
	return service.CloseFriendsService.AddCloseFriend(friend)
}

func (service CloseFriendsService) RemoveCloseFriend(friend *model.CloseFriends) error{
	return service.CloseFriendsService.RemoveCloseFriend(friend)
}

func (service CloseFriendsService) GetAllCloseFriend(id string)([]uuid.UUID,error){
	var friend []model.CloseFriends
	friend, err := service.CloseFriendsService.GetAllCloseFriend(id)

	var list []uuid.UUID
	for i:=range friend {
		list = append(list, friend[i].FriendID)
	}
	return list , err
}
