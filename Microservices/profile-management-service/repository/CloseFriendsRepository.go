package repository

import "gorm.io/gorm"

type CloseFriendsRepository struct {
	Database *gorm.DB
}
