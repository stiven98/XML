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

func (service *ConversationsService) GetConversation(s string, s2 string) model.Conversation {
	return service.ConversationsRepo.GetConversation(s, s2)
}

func (service *ConversationsService) Update(ret model.Conversation) {
	service.ConversationsRepo.Update(ret)
}