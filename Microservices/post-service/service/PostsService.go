package service

import (
	"github.com/google/uuid"
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
func (service *PostsService) CreateCampaign(campaign *model.Campaign) error {
	return	service.PostsRepo.CreateCampaign(campaign)
}

func (service *PostsService) CreateCampaignInf(campaign *model.Campaign) error {
	return	service.PostsRepo.CreateCampaignForInfluencer(campaign)
}

func (service *PostsService) CreateCampaignRequest(campaignReq *dto.CampaignRequestDto) error {
	return	service.PostsRepo.CreateCampaignRequest(campaignReq)
}
func (service *PostsService) CreateTemporaryCampaign(campaign *model.Campaign) error {
	return	service.PostsRepo.CreateTemporaryCampaign(campaign)
}

func (service *PostsService) GetByKey(key string) []model.Post {
	return  service.PostsRepo.GetByKey(key)
}

func (service *PostsService) GetFeed(id string) []model.Post {
	return  service.PostsRepo.GetFeed(id)
}

func (service *PostsService) GetLiked(id string) []model.Post {
	return  service.PostsRepo.GetLiked(id)
}
func (service *PostsService) GetCampaigns(id string) []model.Campaign {
	return  service.PostsRepo.GetCampaigns(id)
}
func (service *PostsService) GetInfluencerCampaigns(id string) []model.Campaign {
	return  service.PostsRepo.GetCampaignsInf(id)
}
func (service *PostsService) GetCampaignReqs(id string) []dto.CampaignRequestDto {
	return  service.PostsRepo.GetCampaignReqs(id)
}
func (service *PostsService) GetTemporaryCampaigns(id string) []model.Campaign {
	return  service.PostsRepo.GetTemporaryCampaigns(id)
}
func (service *PostsService) GetDisliked(id string) []model.Post {
	return  service.PostsRepo.GetDisliked(id)
}
func (service *PostsService) Delete(deletePost *dto.DeletePostDto) bool {
	return  service.PostsRepo.Delete(deletePost)
}
func (service *PostsService) DeleteCampaign(deletePost *dto.DeletePostDto) bool {
	return  service.PostsRepo.DeleteCampaign(deletePost)
}
func (service *PostsService) DeleteCampaignReq(deleteReq *dto.DeleteCampaignReq) bool {
	return  service.PostsRepo.DeleteCampaignReq(deleteReq)
}
func (service *PostsService) GetReported(ids []dto.UserId) ([]model.Post)  {
	  return service.PostsRepo.GetReported(ids)
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
func (service *PostsService) ReportPost(reportReq dto.ReportDto) error {
	return service.PostsRepo.ReportPost(reportReq)
}

func (service *PostsService) AddPostToFeed(keys []string, post *model.Post) error {
	return service.PostsRepo.AddPostToFeed(keys, post)
}

func (service *PostsService) LeaveComment(postId uuid.UUID, ownerId uuid.UUID, comment *model.Comment) error {
	return service.PostsRepo.LeaveComment(postId, ownerId, comment)
}

func (service *PostsService) GetByIds(userid string, postid string) interface{} {
	return service.PostsRepo.GetByIds(userid, postid)
}
func (service *PostsService) GetCampaignsByIds(userid string, campaignid string) model.Campaign {
	return service.PostsRepo.GetCampaignsByIds(userid, campaignid)
}
func (service *PostsService) GetCampaignsByInfluencerIds(userid string, campaignid string) model.Campaign {
	return service.PostsRepo.GetCampaignsByInfluencerIds(userid, campaignid)
}

func (service *PostsService) GetByUserId(userid string) interface{} {
	return service.PostsRepo.GetByUserId(userid)
}

func (service *PostsService) SavePost(post *model.SavedPost) error {
	return service.PostsRepo.SavePost(post)
}

func (service *PostsService) GetAllArchived(id string) []model.SavedPost {
	return service.PostsRepo.GetAllArchived(id)
}

func (service *PostsService) EditArchived(post model.SavedPost) error {
	return service.PostsRepo.EditArchived(post)
}
