package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/requests/comment"
	commentreply "github.com/alifdwt/haiwangram/domain/requests/comment_reply"
	"github.com/alifdwt/haiwangram/domain/requests/like"
	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/domain/responses"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*responses.UserResponse, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
}

type UserService interface {
	UpdateUserById(id int, request *user.UpdateUserRequest) (*responses.UserResponse, error)
	DeleteUserById(id int) (*responses.UserResponse, error)
}

type PhotoService interface {
	CreatePhoto(userId int, request photo.CreatePhotoRequest) (*responses.PhotoResponse, error)
	GetPhotoAll() (*[]responses.PhotoWithRelationResponse, error)
	GetPhotoById(photoId int) (*responses.PhotoWithRelationResponse, error)
	UpdatePhoto(userId int, photoId int, request photo.UpdatePhotoRequest) (*responses.PhotoResponse, error)
	DeletePhoto(photoId int) (*responses.PhotoResponse, error)
}

type CommentService interface {
	CreateComment(userId int, request comment.CreateCommentRequest) (*responses.CommentResponse, error)
	GetCommentAll() (*[]responses.CommentWithRelationResponse, error)
	GetCommentById(commentId int) (*responses.CommentWithRelationResponse, error)
	UpdateComment(commentId int, request comment.UpdateCommentRequest) (*responses.CommentResponse, error)
	DeleteComment(commentId int) (*responses.CommentResponse, error)
}

type CommentReplyService interface {
	CreateCommentReply(userId int, request commentreply.CreateCommentReplyRequest) (*responses.CommentReplyResponse, error)
	GetCommentReplyAll() (*[]responses.CommentReplyWithRelationResponse, error)
	GetCommentReplyById(commentReplyId int) (*responses.CommentReplyWithRelationResponse, error)
	UpdateCommentReply(commentReplyId int, request commentreply.UpdateCommentReplyRequest) (*responses.CommentReplyResponse, error)
	DeleteCommentReply(commentReplyId int) (*responses.CommentReplyResponse, error)
}

type LikeService interface {
	CreateLike(userId int, request like.CreateLikeRequest) (*responses.LikeResponse, error)
	GetLikeById(likeId int) (*responses.LikeResponse, error)
	DeleteLike(likeId int) (*responses.LikeResponse, error)
}
