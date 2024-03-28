package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/like"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/repository"
)

type likeService struct {
	repository repository.LikeRepository
	log        logger.Logger
	mapper     mapper.LikeMapping
}

func NewLikeService(repository repository.LikeRepository, log logger.Logger, mapper mapper.LikeMapping) *likeService {
	return &likeService{
		repository: repository,
		log:        log,
		mapper:     mapper,
	}
}

func (s *likeService) CreateLike(userId int, request like.CreateLikeRequest) (*responses.LikeResponse, error) {
	res, err := s.repository.CreateLike(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToLikeResponse(res)

	return mapper, nil
}

func (s *likeService) GetLikeById(likeId int) (*responses.LikeResponse, error) {
	res, err := s.repository.GetLikeById(likeId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToLikeResponse(res)

	return mapper, nil
}

func (s *likeService) DeleteLike(likeId int) (*responses.LikeResponse, error) {
	res, err := s.repository.DeleteLike(likeId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToLikeResponse(res)

	return mapper, nil
}
