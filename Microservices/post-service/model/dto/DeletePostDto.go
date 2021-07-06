package dto

import "github.com/google/uuid"

type DeletePostDto struct {
	OWNERID uuid.UUID `json:"ownerId"`
	POSTID uuid.UUID `json:"postId"`
}

type DeleteCampaignReq struct {
	OWNERID uuid.UUID `json:"ownerId"`
	CAMPAIGNREQID uuid.UUID `json:"campaignReqId"`
}