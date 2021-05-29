package repository

import "gorm.io/gorm"

type MutedUsersRepository struct {
	DataBase *gorm.DB
}
