package service

import (
	"post_service/model"
	"post_service/repository"
)

type PostsService struct {
	PostsRepo *repository.PostsRepository
}

func (service *PostsService) Create(post *model.Post) error {
	service.PostsRepo.Create(post)
	return nil
}

func (service *PostsService) GetByKey(key string) *model.Post {
	return  service.PostsRepo.GetByKey(key)
}

func (service *PostsService) AddPostToFeed(keys []string, post *model.Post) error {
	return service.PostsRepo.AddPostToFeed(keys, post)
}
