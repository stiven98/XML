package service

import (
	"post_service/model"
	"post_service/model/dto"
	"post_service/repository"
)

type PostsService struct {
	PostsRepo *repository.PostsRepository
}

func (service *PostsService) Create(post *model.Post) error {
	service.PostsRepo.Create(post)
	return nil
}

func (service *PostsService) GetByKey(key string) []model.Post {
	return  service.PostsRepo.GetByKey(key)
}

func (service *PostsService) GetFeed(id string) []model.Post {
	return  service.PostsRepo.GetFeed(id)
}

func (service *PostsService) GetPublic(keys []string) []model.Post {
	return  service.PostsRepo.GetPublic(keys)
}

func (service *PostsService) LikePost(likeReq dto.LikeDto) error {
	return service.PostsRepo.LikePost(likeReq)
}

func (service *PostsService) DislikePost(dislikeReq dto.LikeDto) error {
	return service.PostsRepo.DislikePost(dislikeReq)
}

func (service *PostsService) AddPostToFeed(keys []string, post *model.Post) error {
	return service.PostsRepo.AddPostToFeed(keys, post)
}
