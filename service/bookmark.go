package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/bookmark"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/repository"
)

type bookmarkService struct {
	repository repository.BookmarkRepository
	log        logger.Logger
	mapper     mapper.BookmarkMapping
}

func NewBookmarkService(repository repository.BookmarkRepository, log logger.Logger, mapper mapper.BookmarkMapping) *bookmarkService {
	return &bookmarkService{
		repository: repository,
		log:        log,
		mapper:     mapper,
	}
}

func (s *bookmarkService) CreateBookmark(userId int, request bookmark.CreateBookmarkRequest) (*responses.BookmarkResponse, error) {
	res, err := s.repository.CreateBookmark(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToBookmarkResponse(res)

	return mapper, nil
}

func (s *bookmarkService) GetBookmarkById(bookmarkId int) (*responses.BookmarkWithRelationResponse, error) {
	res, err := s.repository.GetBookmarkById(bookmarkId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToBookmarkWithRelationResponse(res)

	return mapper, nil
}

func (s *bookmarkService) GetBookmarkByUserId(userId int) (*[]responses.BookmarkWithRelationResponse, error) {
	res, err := s.repository.GetBookmarkByUserId(userId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToBookmarkWithRelationResponses(res)

	return &mapper, nil
}

func (s *bookmarkService) DeleteBookmark(photoId int, userId int) (*responses.BookmarkResponse, error) {
	res, err := s.repository.DeleteBookmark(photoId, userId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToBookmarkResponse(res)

	return mapper, nil
}
