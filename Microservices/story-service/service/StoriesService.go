package service

import (
	"storyservice/model"
	"storyservice/repository"
)

type StoriesService struct {
	StoriesRepo *repository.StoriesRepository
}

func (service *StoriesService) Create(story *model.Story) error {
	service.StoriesRepo.Create(story)
	return nil
}

func (service *StoriesService) GetByKey(key string) *model.Story {
	return  service.StoriesRepo.GetByKey(key)
}
