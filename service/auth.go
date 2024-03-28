package service

import (
	"errors"
	"time"

	"github.com/alifdwt/haiwangram/domain/requests/auth"
	"github.com/alifdwt/haiwangram/domain/responses"

	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/dotenv"
	"github.com/alifdwt/haiwangram/pkg/hashing"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/pkg/token"
	"github.com/alifdwt/haiwangram/repository"
)

type authService struct {
	config dotenv.Config
	user   repository.UserRepository
	hash   hashing.Hashing
	log    logger.Logger
	token  token.Maker
	mapper mapper.UserMapping
}

func NewAuthService(config dotenv.Config, user repository.UserRepository, hash hashing.Hashing, log logger.Logger, token token.Maker, mapper mapper.UserMapping) *authService {
	return &authService{
		config: config,
		user:   user,
		hash:   hash,
		log:    log,
		token:  token,
		mapper: mapper,
	}
}

func (s *authService) Register(input *auth.RegisterRequest) (*responses.UserResponse, error) {
	hashing, err := s.hash.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input.Password = hashing

	user, err := s.user.CreateUser(input)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(user)

	return mapper, nil
}

func (s *authService) Login(input *auth.LoginRequest) (*responses.Token, error) {
	res, err := s.user.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("error while get user: " + err.Error())
	}

	err = s.hash.ComparePassword(res.Password, input.Password)
	if err != nil {
		return nil, errors.New("error while compare password: " + err.Error())
	}

	age := time.Now().Year() - res.BirthDate.Year()

	accessToken, err := s.createAccessToken(res.ID, res.Email, res.Username, age)
	if err != nil {
		return nil, err
	}

	return &responses.Token{
		Token: accessToken,
	}, nil
}

func (s *authService) createAccessToken(id int, email string, username string, age int) (string, error) {
	res, err := s.token.CreateToken(id, email, username, age, "access", s.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return "", err
	}

	return res, nil
}
