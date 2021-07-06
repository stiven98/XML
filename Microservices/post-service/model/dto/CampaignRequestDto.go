package dto

import "github.com/google/uuid"

type CampaignRequestDto struct {
	ID uuid.UUID `json:"id"`
	AGENTID uuid.UUID `json:"agentId"`
	INFLUENCERID uuid.UUID `json:"influencerId"`
	CAMPAIGNID uuid.UUID `json:"campaignId"`
}

type AddInfluencerDto struct {
	ID uuid.UUID `json:"id"`
	OWNERID uuid.UUID `json:"ownerId"`
	INFLUENCERID uuid.UUID `json:"influencerId"`
	CAMPAIGNID uuid.UUID `json:"campaignId"`
}