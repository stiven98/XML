package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type MessagesService struct {
	MessagesRepo *repository.MessagesRepository
}

func (service *MessagesService) Create(message *model.Message) error {
	service.MessagesRepo.Create(message)
	return nil
}

func (service *MessagesService) GetAll() []model.Message {
	return  service.MessagesRepo.GetAll()
}
