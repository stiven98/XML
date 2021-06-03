package service

import (
	"agent-service/model"
	"agent-service/repository"
)

type CampaignService struct {
	CampaignRepo *repository.CampaignRepository
}

func (service *CampaignService) Create(campaign *model.Campaign) error {
	service.CampaignRepo.Create(campaign)
	return nil
}

func (service *CampaignService) GetAll() []model.Campaign {
	return  service.CampaignRepo.GetAll()
}
func (service *CampaignService) Delete(id string) error {
	 service.CampaignRepo.Delete(id)
	 return nil
}


