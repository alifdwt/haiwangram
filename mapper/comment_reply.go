package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/models"
)

type commentReplyMapper struct {
}

func NewCommentReplyMapper() *commentReplyMapper {
	return &commentReplyMapper{}
}

func (m *commentReplyMapper) ToCommentReplyResponse(request *models.CommentReply) *responses.CommentReplyResponse {
	return &responses.CommentReplyResponse{
		ID:        request.ID,
		Message:   request.Message,
		UserID:    request.UserID,
		CommentID: request.CommentID,
	}
}

func (m *commentReplyMapper) ToCommentReplyWithRelationResponse(request *models.CommentReply) *responses.CommentReplyWithRelationResponse {
	return &responses.CommentReplyWithRelationResponse{
		ID:        request.ID,
		Message:   request.Message,
		UserID:    request.UserID,
		CommentID: request.CommentID,
		Comment:   *NewCommentMapper().ToCommentResponse(&request.Comment),
		User:      *NewUserMapper().ToUserResponse(&request.User),
	}
}

func (m *commentReplyMapper) ToCommentReplyWithRelationResponses(requests *[]models.CommentReply) []responses.CommentReplyWithRelationResponse {
	var commentReplyResponses []responses.CommentReplyWithRelationResponse
	for _, request := range *requests {
		commentReplyResponses = append(commentReplyResponses, *m.ToCommentReplyWithRelationResponse(&request))
	}

	if len(commentReplyResponses) > 0 {
		return commentReplyResponses
	}

	return []responses.CommentReplyWithRelationResponse{}
}
