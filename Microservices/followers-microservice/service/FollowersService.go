package service

import (
	"followers-microservice/repository"
)

type FollowersService struct {
	FollowersRepository *repository.FollowersRepository
}

func (s FollowersService) GetFollowers(id string) []string {
	return s.FollowersRepository.GetFollowers(id)
}

func (s FollowersService) UserExists(id string) interface{} {
	return s.FollowersRepository.UserExist(id)
}

func (s FollowersService) GetFollowing(id string) []string {
	return s.FollowersRepository.GetFollowing(id)
}

func (s FollowersService) CheckRelationship(userID string, targetID string) bool {
	return s.FollowersRepository.CheckRelationship(userID, targetID)
}

func (s FollowersService) Follow(userID string, targetID string) error {
	return s.FollowersRepository.Follow(userID, targetID)
}

func (s FollowersService) AddNode(id string) (interface{}, error) {
	return s.FollowersRepository.AddNode(id)
}

func (s FollowersService) Request(userID string, targetID string) interface{} {
	return s.FollowersRepository.Request(userID, targetID)
}

func (s FollowersService) GetRequests(id string) []string {
	return s.FollowersRepository.GetRequests(id)
}

func (s FollowersService) Unfollow(userID string, targetID string) interface{} {
	return s.FollowersRepository.Unfollow(userID, targetID)
}
