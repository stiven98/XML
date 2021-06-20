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


func (r *BlockedUsersRepository) GetAllBlockedByUserId(id string) ([]model.BlockedUsers, error){
	var blocked []model.BlockedUsers
	response := r.DataBase.Model(model.BlockedUsers{}).Where("blocked_by_id = ?", id).Find(&blocked)
	return  blocked, response.Error
}

