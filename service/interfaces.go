package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/domain/responses"
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
