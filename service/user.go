package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/user"
	"github.com/alifdwt/haiwangram/domain/responses"
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

func (s *userService) GetUserByEmail(email string) (*responses.UserResponse, error) {
	res, err := s.Repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) GetUserByUsername(username string) (*responses.UserResponse, error) {
	res, err := s.Repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) UpdateUserById(id int, request *user.UpdateUserRequest) (*responses.UserResponse, error) {
	res, err := s.Repository.UpdateUserById(id, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) DeleteUserById(id int) (*responses.UserResponse, error) {
	res, err := s.Repository.DeleteUserById(id)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(res)

	return mapper, nil
}

func (s *userService) GetRandomUser(count int) (*[]responses.UserResponse, error) {
	res, err := s.Repository.GetRandomUser(count)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponses(res)

	return &mapper, nil
}
