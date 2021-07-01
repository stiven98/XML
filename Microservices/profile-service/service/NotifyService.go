package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type NotifyService struct {
	NotifyRepository *repository.NotifyRepository
}

func(service NotifyService) Create(notify *model.Notify) error{
	return service.NotifyRepository.Create(notify)
}


func(service NotifyService) GetAllNotifyByUserId(id string)([]model.Notify ,error){
	return service.NotifyRepository.GetAllNotifyByUserId(id)
}
