package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/photo"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/repository"
)

type photoService struct {
	repository repository.PhotoRepository
	log        logger.Logger
	mapper     mapper.PhotoMapping
}

func NewPhotoService(repository repository.PhotoRepository, log logger.Logger, mapper mapper.PhotoMapping) *photoService {
	return &photoService{
		repository: repository,
		log:        log,
		mapper:     mapper,
	}
}

func (s *photoService) CreatePhoto(userId int, request photo.CreatePhotoRequest) (*responses.PhotoResponse, error) {
	res, err := s.repository.CreatePhoto(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponse(res)

	return mapper, nil
}

func (s *photoService) GetPhotoAll(limit int) (*[]responses.PhotoWithRelationResponse, error) {
	res, err := s.repository.GetPhotoAll(limit)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoWithRelationResponses(res)

	return &mapper, nil
}

func (s *photoService) GetPhotoById(photoId int) (*responses.PhotoWithRelationResponse, error) {
	res, err := s.repository.GetPhotoById(photoId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoWithRelationResponse(res)

	return mapper, nil
}

func (s *photoService) GetPhotoByUserId(userId int, limit int) (*[]responses.PhotoWithRelationResponse, error) {
	res, err := s.repository.GetPhotoByUserId(userId, limit)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoWithRelationResponses(res)

	return &mapper, nil
}

func (s *photoService) UpdatePhoto(userId int, photoId int, request photo.UpdatePhotoRequest) (*responses.PhotoResponse, error) {
	res, err := s.repository.UpdatePhoto(userId, photoId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponse(res)

	return mapper, nil
}

func (s *photoService) DeletePhoto(photoId int) (*responses.PhotoResponse, error) {
	res, err := s.repository.DeletePhoto(photoId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponse(res)

	return mapper, nil
}
