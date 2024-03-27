package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/domain/responses"
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
