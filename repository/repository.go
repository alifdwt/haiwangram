package repository

import "gorm.io/gorm"

type Repositories struct {
	User    UserRepository
	Photo   PhotoRepository
	Comment CommentRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:    NewUserRepository(db),
		Photo:   NewPhotoRepository(db),
		Comment: NewCommentRepository(db),
	}
}
