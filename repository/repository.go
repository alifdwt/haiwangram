package repository

import "gorm.io/gorm"

type Repositories struct {
	User  UserRepository
	Photo PhotoRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:  NewUserRepository(db),
		Photo: NewPhotoRepository(db),
	}
}
