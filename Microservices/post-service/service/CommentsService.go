package service

import (
	"post_service/model"
	"post_service/repository"
)

type CommentsService struct {
	CommentsRepo *repository.CommentsRepository
}

func (service *CommentsService) Create(comment *model.Comment) error {
	service.CommentsRepo.Create(comment)
	return nil
}

func (service *CommentsService) GetByKey(key string) *model.Comment {
	return  service.CommentsRepo.GetByKey(key)
}
