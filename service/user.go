package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/user"
	userRes "github.com/alifdwt/haiwangram/domain/responses/user"
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/hashing"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/repository"
)

type userService struct {
	Repository repository.UserRepository
	hash       hashing.Hashing
	log        logger.Logger
	mapper     mapper.UserMapping
}

func NewUserService(user repository.UserRepository, hash hashing.Hashing, log logger.Logger, mapper mapper.UserMapping) *userService {
	return &userService{
		Repository: user,
		hash:       hash,
		log:        log,
		mapper:     mapper,
	}
}

func (s *userService) UpdateUserById(id int, request *user.UpdateUserRequest) (*userRes.UserResponse, error) {
	res, err := s.Repository.UpdateUserById(id, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) DeleteUserById(id int) (*userRes.UserResponse, error) {
	res, err := s.Repository.DeleteUserById(id)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}
