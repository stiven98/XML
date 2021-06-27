package repository

import (
	"gorm.io/gorm"
	"profile-management-service/model"
)

type CloseFriendsRepository struct {
	Database *gorm.DB
}

func (r CloseFriendsRepository) AddCloseFriend(friend *model.CloseFriends) error{
	response := r.Database.Create(friend)
	return response.Error
}

func (r CloseFriendsRepository)  RemoveCloseFriend(friend *model.CloseFriends) error{
	response := r.Database.Delete(friend)
	return response.Error
}

func (r CloseFriendsRepository) GetAllCloseFriend(id string) ([]model.CloseFriends, error)  {
	 var friends []model.CloseFriends
	 response := r.Database.Model(model.CloseFriends{}).Where("user_id = ?", id).Find(&friends)
	 return  friends, response.Error
}
