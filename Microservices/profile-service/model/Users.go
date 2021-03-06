package model

import (
	"github.com/google/uuid"
	"time"
)

type Gender string
const (
	MALE Gender = "male"
	FEMALE Gender = "female"
)

type TypeOfUser string
const (
	ADMIN TypeOfUser = "admin"
	USER TypeOfUser = "user"
	AGENT TypeOfUser = "agent"
)
type SystemUser struct {
	ID   uuid.UUID `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName" gorm:"not null"`
	LastName string    `json:"lastName" gorm:"not null"`
	Username string    `json:"username" gorm:"unique; not null"`
	Email string    `json:"email" gorm:"unique; not null"`
	Password string    `json:"password" gorm:"not null"`
	Gender Gender 	`json:"gender" gorm:"not null"`
	TypeOfUser TypeOfUser `json:"type_of_user" gorm:"not null"`
	DateOfBirth time.Time  `json:"dateOfBirth" gorm:"not null"`
	PicturePath string  `json:"picturePath" gorm:"not null"`
}

type User struct {
	UserID   uuid.UUID `json:"id"`
	SystemUser SystemUser `json:"system_user" gorm:"foreignKey:UserID"`
	IsPublic bool `json:"isPublic" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
	WebSite string `json:"webSite" gorm:"not null"`
	Biography string `json:"biography"`
	AllowedTags bool `json:"allowedTags"`
	IsBlocked bool `json:"isBlocked"`
	IsVerified bool `json:"isVerified"`
	AcceptMessagesFromNotFollowProfile bool `json:"acceptMessagesFromNotFollowProfiles"`
	NotifyLike bool `json:"notifyLike"`
	NotifyMessages bool `json:"notifyMessages"`
	NotifyDislike bool `json:"notifyDislike" `
	NotifyComments bool `json:"notifyComments"`
	NotifyLikeFromNotFollowProfile bool `json:"notifyLikeFromNotFollowProfile""`
	NotifyDislikeFromNotFollowProfile bool `json:"notifyDislikeFromNotFollowProfile""`
	NotifyCommentFromNotFollowProfile bool `json:"notifyCommentFromNotFollowProfile""`
	NotifyMessageFromNotFollowProfile bool `json:"notifyMessageFromNotFollowProfile""`
	IsCreate string `json:"isCreate"`
}
type Administrator struct {
	UserID   uuid.UUID `json:"id"`
	SystemUser SystemUser `json:"system_user" gorm:"foreignKey:UserID"`
}

type Agent struct {
	UserID   uuid.UUID `json:"id"`
	SystemUser SystemUser `json:"system_user" gorm:"foreignKey:UserID"`
	IsPublic bool `json:"isPublic" gorm:"not null"`
	PhoneNumber string `json:"phoneNumber" gorm:"not null"`
	WebSite string `json:"webSite" gorm:"not null"`
	Biography string `json:"biography"`
	AllowedTags bool `json:"allowedTags"`
	IsBlocked bool `json:"isBlocked"`
	IsVerified bool `json:"isVerified"`
	AcceptMessagesFromNotFollowProfile bool `json:"acceptMessagesFromNotFollowProfiles"`
	NotifyLike bool `json:"notifyLike"`
	NotifyMessages bool `json:"notifyMessages"`
	NotifyDislike bool `json:"notifyDislike" `
	NotifyComments bool `json:"notifyComments"`
	NotifyLikeFromNotFollowProfile bool `json:"notifyLikeFromNotFollowProfile""`
	NotifyDislikeFromNotFollowProfile bool `json:"notifyDislikeFromNotFollowProfile""`
	NotifyCommentFromNotFollowProfile bool `json:"notifyCommentFromNotFollowProfile""`
	NotifyMessageFromNotFollowProfile bool `json:"notifyMessageFromNotFollowProfile""`
}


//func (sysUser *SystemUser) BeforeCreate(scope *gorm.DB) error {
//	sysUser.ID = uuid.New()
//	return nil
//}
