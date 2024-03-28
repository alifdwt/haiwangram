package repository

import "gorm.io/gorm"

type Repositories struct {
	User         UserRepository
	Photo        PhotoRepository
	Comment      CommentRepository
	CommentReply CommentReplyRepository
	Like         LikeRepository
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		User:         NewUserRepository(db),
		Photo:        NewPhotoRepository(db),
		Comment:      NewCommentRepository(db),
		CommentReply: NewCommentReplyRepository(db),
		Like:         NewLikeRepository(db),
	}
}
