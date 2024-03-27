package service

import (
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/dotenv"
	"github.com/alifdwt/haiwangram/pkg/hashing"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/alifdwt/haiwangram/repository"
)

type Service struct {
	Auth    AuthService
	User    UserService
	Photo   PhotoService
	Comment CommentService
}

type Deps struct {
	Config     dotenv.Config
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	TokenMaker token.Maker
	Logger     logger.Logger
	Mapper     mapper.Mapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:    NewAuthService(deps.Config, deps.Repository.User, deps.Hashing, deps.Logger, deps.TokenMaker, deps.Mapper.UserMapper),
		User:    NewUserService(deps.Repository.User, deps.Hashing, deps.Logger, deps.Mapper.UserMapper),
		Photo:   NewPhotoService(deps.Repository.Photo, deps.Logger, deps.Mapper.PhotoMapper),
		Comment: NewCommentService(deps.Repository.Comment, deps.Logger, deps.Mapper.CommentMapper),
	}
}
