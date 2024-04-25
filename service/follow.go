package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/follow"
	"github.com/alifdwt/haiwangram/models"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/repository"
)

type followService struct {
	repository repository.FollowRepository
	log        logger.Logger
}

func NewFollowService(repository repository.FollowRepository, log logger.Logger) *followService {
	return &followService{repository: repository, log: log}
}

func (s *followService) CreateFollow(userId int, request follow.CreateFollowRequest) (*models.Follow, error) {
	res, err := s.repository.CreateFollow(userId, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *followService) GetFollowedByUserId(userId int) (*[]models.Follow, error) {
	res, err := s.repository.GetFollowedByUserId(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *followService) GetFollowerByUserId(userId int) (*[]models.Follow, error) {
	res, err := s.repository.GetFollowerByUserId(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *followService) DeleteFollow(userId int, request follow.DeleteFollowRequest) (*models.Follow, error) {
	res, err := s.repository.DeleteFollow(userId, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
