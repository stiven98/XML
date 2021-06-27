package repository

import (
	"gorm.io/gorm"
	"profileservice/model"
)

type NotifyRepository struct{
	Database *gorm.DB
}

func (repo *NotifyRepository) Create(notify *model.Notify) error{
	result := repo.Database.Create(notify)
	return result.Error
}


func (repo *NotifyRepository) GetAllNotifyByUserId(id string)([]model.Notify ,error){
	var notify []model.Notify
	result := repo.Database.Model(model.Notify{}).Where("notify_user_id = ?", id).Find(&notify)
	return notify,result.Error
}