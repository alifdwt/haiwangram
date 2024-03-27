package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/requests/comment"
	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/domain/responses"
	commentRes "github.com/alifdwt/haiwangram/domain/responses/comment"
	photoRes "github.com/alifdwt/haiwangram/domain/responses/photo"
	userRes "github.com/alifdwt/haiwangram/domain/responses/user"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*userRes.UserResponse, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
}

type UserService interface {
	UpdateUserById(id int, request *user.UpdateUserRequest) (*userRes.UserResponse, error)
	DeleteUserById(id int) (*userRes.UserResponse, error)
}

type PhotoService interface {
	CreatePhoto(userId int, request photo.CreatePhotoRequest) (*photoRes.PhotoResponse, error)
	GetPhotoAll() (*[]photoRes.PhotoWithRelationResponse, error)
	GetPhotoById(photoId int) (*photoRes.PhotoWithRelationResponse, error)
	UpdatePhoto(userId int, photoId int, request photo.UpdatePhotoRequest) (*photoRes.PhotoResponse, error)
	DeletePhoto(photoId int) (*photoRes.PhotoResponse, error)
}

type CommentService interface {
	CreateComment(userId int, request comment.CreateCommentRequest) (*commentRes.CommentResponse, error)
	GetCommentAll() (*[]commentRes.CommentWithRelationResponse, error)
	GetCommentById(commentId int) (*commentRes.CommentWithRelationResponse, error)
	UpdateComment(commentId int, request comment.UpdateCommentRequest) (*commentRes.CommentResponse, error)
	DeleteComment(commentId int) (*commentRes.CommentResponse, error)
}
