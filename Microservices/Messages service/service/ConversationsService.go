package service

import (
	"profileservice/model"
	"profileservice/repository"
)

type ConversationsService struct {
	ConversationsRepo *repository.ConversationRepository
}

func (service *ConversationsService) Create(conversation *model.Conversation) error {
	service.ConversationsRepo.Create(conversation)
	return nil
}

func (service *ConversationsService) GetAll() []model.Conversation {
	return  service.ConversationsRepo.GetAll()
}