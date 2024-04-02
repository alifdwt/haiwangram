package service

import (
	"github.com/alifdwt/haiwangram/domain/requests/comment"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/repository"
)

type commentService struct {
	repository repository.CommentRepository
	log        logger.Logger
	mapper     mapper.CommentMapping
}

func NewCommentService(repository repository.CommentRepository, log logger.Logger, mapper mapper.CommentMapping) *commentService {
	return &commentService{repository, log, mapper}
}

func (s *commentService) CreateComment(userId int, request comment.CreateCommentRequest) (*responses.CommentResponse, error) {
	res, err := s.repository.CreateComment(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponse(res)

	return mapper, nil
}

func (s *commentService) GetCommentAll() (*[]responses.CommentWithRelationResponse, error) {
	res, err := s.repository.GetCommentAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentWithRelationResponses(res)

	return &mapper, nil
}

func (s *commentService) GetCommentById(commentId int) (*responses.CommentWithRelationResponse, error) {
	res, err := s.repository.GetCommentById(commentId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentWithRelationResponse(res)

	return mapper, nil
}

func (s *commentService) GetCommentByPhotoId(photoId int) (*[]responses.CommentWithRelationResponse, error) {
	res, err := s.repository.GetCommentByPhotoId(photoId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentWithRelationResponses(res)

	return &mapper, nil
}

func (s *commentService) UpdateComment(commentId int, request comment.UpdateCommentRequest) (*responses.CommentResponse, error) {
	res, err := s.repository.UpdateComment(commentId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponse(res)

	return mapper, nil
}

func (s *commentService) DeleteComment(commentId int) (*responses.CommentResponse, error) {
	res, err := s.repository.DeleteComment(commentId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentResponse(res)

	return mapper, nil
}
