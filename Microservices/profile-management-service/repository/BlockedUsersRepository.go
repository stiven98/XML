package repository

import (
	"gorm.io/gorm"
	"profile-management-service/model"
)

type BlockedUsersRepository struct {
	DataBase *gorm.DB
}

func (r BlockedUsersRepository) BlockUserByUser(blockedUsers *model.BlockedUsers) error {
	response := r.DataBase.Create(blockedUsers)
	return response.Error
}

