package repository

import (
	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/requests/comment"
	commentreply "github.com/alifdwt/haiwangram/domain/requests/comment_reply"
	"github.com/alifdwt/haiwangram/domain/requests/like"
	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	CreateUser(registerReq *auth.RegisterRequest) (*models.User, error)
	UpdateUserById(id int, updatedUser *user.UpdateUserRequest) (*models.User, error)
	DeleteUserById(id int) (*models.User, error)
	GetRandomUser(count int) (*[]models.User, error)
}

type PhotoRepository interface {
	GetPhotoAll(limit int) (*[]models.Photo, error)
	GetPhotoById(photoId int) (*models.Photo, error)
	CreatePhoto(userId int, request photo.CreatePhotoRequest) (*models.Photo, error)
	UpdatePhoto(userId int, photoId int, updatedPhoto photo.UpdatePhotoRequest) (*models.Photo, error)
	DeletePhoto(photoId int) (*models.Photo, error)
}

type CommentRepository interface {
	CreateComment(userId int, request comment.CreateCommentRequest) (*models.Comment, error)
	GetCommentAll() (*[]models.Comment, error)
	GetCommentById(commentId int) (*models.Comment, error)
	UpdateComment(commentId int, updatedComment comment.UpdateCommentRequest) (*models.Comment, error)
	DeleteComment(commentId int) (*models.Comment, error)
}

type CommentReplyRepository interface {
	CreateCommentReply(userId int, request commentreply.CreateCommentReplyRequest) (*models.CommentReply, error)
	GetCommentReplyAll() (*[]models.CommentReply, error)
	GetCommentReplyById(commentReplyId int) (*models.CommentReply, error)
	UpdateCommentReply(commentReplyId int, updatedCommentReply commentreply.UpdateCommentReplyRequest) (*models.CommentReply, error)
	DeleteCommentReply(commentReplyId int) (*models.CommentReply, error)
}

type LikeRepository interface {
	CreateLike(userId int, request like.CreateLikeRequest) (*models.Like, error)
	// GetLikeAll() (*[]models.Like, error)
	GetLikeById(likeId int) (*models.Like, error)
	DeleteLike(photoId int, userId int) (*models.Like, error)
}
