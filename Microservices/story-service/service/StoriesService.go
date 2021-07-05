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

func (service *StoriesService) AddStoryToFeed(keys []string, m *model.Story) error {
	return service.StoriesRepo.AddStoryToFeed(keys, m)
}

func (service *StoriesService) GetFeed(id string) []model.Story {
	return  service.StoriesRepo.GetFeed(id)
}

func (service *StoriesService) GetMyStories(id string) interface{} {
	return  service.StoriesRepo.GetMyStories(id)
}
