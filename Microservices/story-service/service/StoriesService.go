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

func (service *StoriesService) GetMyStories(id string) [] model.Story {
	return  service.StoriesRepo.GetMyStories(id)
}

func (service *StoriesService) AddToHighlights(highlight model.Highlight) error {
	return service.StoriesRepo.AddToHighlights(highlight)
}

func (service *StoriesService) RemoveFromHighlights(highlight model.Highlight) error {
	return service.StoriesRepo.RemoveFromHighlights(highlight)
}

func (service *StoriesService) GetHighlights(id string) []model.Story {
	return  service.StoriesRepo.GetHighlights(id)
}