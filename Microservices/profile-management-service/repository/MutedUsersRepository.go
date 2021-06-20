package repository

import (
	"gorm.io/gorm"
	"profile-management-service/model"
)

type MutedUsersRepository struct {
	DataBase *gorm.DB
}

func (r MutedUsersRepository) IsMuted(mutedUsers *model.MutedUsers) (model.MutedUsers, error) {
	var muted model.MutedUsers
	response := r.DataBase.Model(model.MutedUsers{}).Where("muted_by_id = ? AND muted_id = ?", mutedUsers.MutedByID, mutedUsers.MutedID).Find(&muted)
	return muted, response.Error
}

func (r MutedUsersRepository) MutedUserByUser(mutedUsers *model.MutedUsers) error {
	response := r.DataBase.Create(mutedUsers)
	return response.Error
}

func (r MutedUsersRepository) UnMutedUserByUser(mutedUsers *model.MutedUsers) error {
	response := r.DataBase.Delete(mutedUsers)
	return response.Error
}
