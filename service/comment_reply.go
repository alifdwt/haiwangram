package service

import (
	commentreply "github.com/alifdwt/haiwangram/domain/requests/comment_reply"
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/mapper"
	"github.com/alifdwt/haiwangram/pkg/logger"
	"github.com/alifdwt/haiwangram/repository"
)

type commentReplyService struct {
	repository repository.CommentReplyRepository
	log        logger.Logger
	mapper     mapper.CommentReplyMapping
}

func NewCommentReplyService(repository repository.CommentReplyRepository, log logger.Logger, mapper mapper.CommentReplyMapping) *commentReplyService {
	return &commentReplyService{
		repository: repository,
		log:        log,
		mapper:     mapper,
	}
}

func (s *commentReplyService) CreateCommentReply(userId int, request commentreply.CreateCommentReplyRequest) (*responses.CommentReplyResponse, error) {
	res, err := s.repository.CreateCommentReply(userId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentReplyResponse(res)

	return mapper, nil
}

func (s *commentReplyService) GetCommentReplyAll() (*[]responses.CommentReplyWithRelationResponse, error) {
	res, err := s.repository.GetCommentReplyAll()
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentReplyWithRelationResponses(res)

	return &mapper, nil
}

func (s *commentReplyService) GetCommentReplyById(commentReplyId int) (*responses.CommentReplyWithRelationResponse, error) {
	res, err := s.repository.GetCommentReplyById(commentReplyId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentReplyWithRelationResponse(res)

	return mapper, nil
}

func (s *commentReplyService) UpdateCommentReply(commentReplyId int, request commentreply.UpdateCommentReplyRequest) (*responses.CommentReplyResponse, error) {
	res, err := s.repository.UpdateCommentReply(commentReplyId, request)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentReplyResponse(res)

	return mapper, nil
}

func (s *commentReplyService) DeleteCommentReply(commentReplyId int) (*responses.CommentReplyResponse, error) {
	res, err := s.repository.DeleteCommentReply(commentReplyId)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToCommentReplyResponse(res)

	return mapper, nil
}
