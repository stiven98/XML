package Dto

type NotificationStatusesDto struct {
	NotifyLike bool `json:"notifyLike"`
	NotifyMessages bool `json:"notifyMessages"`
	NotifyDislike bool `json:"notifyDislike" `
	NotifyComments bool `json:"notifyComments"`
	NotifyLikeFromNotFollowProfile bool `json:"notifyLikeFromNotFollowProfile""`
	NotifyDislikeFromNotFollowProfile bool `json:"notifyDislikeFromNotFollowProfile""`
	NotifyCommentFromNotFollowProfile bool `json:"notifyCommentFromNotFollowProfile""`
	NotifyMessageFromNotFollowProfile bool `json:"notifyMessageFromNotFollowProfile""`
}

