package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/photo"
	photoRes "github.com/alifdwt/haiwangram/domain/responses/photo"
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

func (s *photoService) CreatePhoto(userId int, request photo.CreatePhotoRequest) (*photoRes.PhotoResponse, error) {
	res, err := s.repository.CreatePhoto(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponse(res)

	return mapper, nil
}

func (s *photoService) GetPhotoAll() (*[]photoRes.PhotoWithRelationResponse, error) {
	res, err := s.repository.GetPhotoAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoWithRelationResponses(res)

	return &mapper, nil
}

func (s *photoService) GetPhotoById(photoId int) (*photoRes.PhotoWithRelationResponse, error) {
	res, err := s.repository.GetPhotoById(photoId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoWithRelationResponse(res)

	return mapper, nil
}

func (s *photoService) UpdatePhoto(userId int, photoId int, request photo.UpdatePhotoRequest) (*photoRes.PhotoResponse, error) {
	res, err := s.repository.UpdatePhoto(userId, photoId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponse(res)

	return mapper, nil
}

func (s *photoService) DeletePhoto(photoId int) (*photoRes.PhotoResponse, error) {
	res, err := s.repository.DeletePhoto(photoId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToPhotoResponse(res)

	return mapper, nil
}
