package service

import (
	"github.com/google/uuid"
	"profile-management-service/model"
	"profile-management-service/repository"
)

type SubscribeAccService struct {
	SubscriberAccRepository *repository.SubscribeAccRepository
}

func (service SubscribeAccService) Subscribe(sub *model.SubscribeAcc) error{
	return service.SubscriberAccRepository.Subscribe(sub)
}


func (service SubscribeAccService) UnSubscribe(sub *model.SubscribeAcc) error{
	return service.SubscriberAccRepository.UnSubscribe(sub)
}

func (service SubscribeAccService) GetAllSubscribers(id string)([]uuid.UUID,error){
	var subscriber []model.SubscribeAcc
	subscriber, err := service.SubscriberAccRepository.GetAllSubscribers(id)

	var list []uuid.UUID
	for i:=range subscriber {
		list = append(list, subscriber[i].SubscribeByID)
	}
	return list , err
}

func (service SubscribeAccService) IsSubscribed(s *model.SubscribeAcc) (bool,error) {
	var sub model.SubscribeAcc
	sub, err := service.SubscriberAccRepository.IsSubscribed(s)

	if s.SubscribeID != sub.SubscribeID{
		return false,err
	}

	return true,err
}


