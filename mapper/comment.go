package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/models"
)

type commentMapper struct {
}

func NewCommentMapper() *commentMapper {
	return &commentMapper{}
}

func (m *commentMapper) ToCommentResponse(request *models.Comment) *responses.CommentResponse {
	return &responses.CommentResponse{
		ID:      request.ID,
		Message: request.Message,
		UserID:  request.UserID,
		PhotoID: request.PhotoID,
	}
}

func (m *commentMapper) ToCommentWithRelationResponse(request *models.Comment) *responses.CommentWithRelationResponse {
	var commentReplies []responses.CommentReplyResponse
	if request.Replies != nil {
		for _, reply := range request.Replies {
			commentReplies = append(commentReplies, *NewCommentReplyMapper().ToCommentReplyResponse(&reply))
		}
	} else {
		commentReplies = []responses.CommentReplyResponse{}
	}

	return &responses.CommentWithRelationResponse{
		ID:      request.ID,
		Message: request.Message,
		UserID:  request.UserID,
		PhotoID: request.PhotoID,
		Photo:   *NewPhotoMapper().ToPhotoResponse(&request.Photo),
		User:    *NewUserMapper().ToUserResponse(&request.User),
		Replies: commentReplies,
	}
}

func (m *commentMapper) ToCommentWithRelationResponses(requests *[]models.Comment) []responses.CommentWithRelationResponse {
	var commentResponses []responses.CommentWithRelationResponse
	for _, request := range *requests {
		commentResponses = append(commentResponses, *m.ToCommentWithRelationResponse(&request))
	}

	if len(commentResponses) > 0 {
		return commentResponses
	}

	return []responses.CommentWithRelationResponse{}
}
