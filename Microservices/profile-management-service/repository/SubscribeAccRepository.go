package repository

import (
	"gorm.io/gorm"
	"profile-management-service/model"
)

type SubscribeAccRepository struct {
	DataBase *gorm.DB
}

func (r SubscribeAccRepository) Subscribe(sub *model.SubscribeAcc) error {
	response := r.DataBase.Create(sub)
	return  response.Error
}

func (r SubscribeAccRepository) UnSubscribe(sub *model.SubscribeAcc) error {
	response := r.DataBase.Delete(sub)
	return  response.Error
}

//vraca sve oni koji su pretplaceni na moj profil da dobiju notifikaciju
func (r SubscribeAccRepository) GetAllSubscribers(id string) ([]model.SubscribeAcc, error){
	var subscriber []model.SubscribeAcc
	response := r.DataBase.Model(model.SubscribeAcc{}).Where("subscribe_id = ?", id).Find(&subscriber)
	return  subscriber, response.Error
}

func (r SubscribeAccRepository) IsSubscribed(subscriber *model.SubscribeAcc) (model.SubscribeAcc, error){
	var sub model.SubscribeAcc
	response := r.DataBase.Model(model.SubscribeAcc{}).Where(" subscribe_by_id = ? AND subscribe_id = ?", subscriber.SubscribeByID,subscriber.SubscribeID).Find(&sub)
	return  sub, response.Error
}