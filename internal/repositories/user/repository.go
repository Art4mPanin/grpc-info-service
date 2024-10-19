package user

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *Repository {
	return &Repository{DB: DB}
}
