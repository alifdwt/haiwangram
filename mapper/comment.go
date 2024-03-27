package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses/comment"
	"github.com/alifdwt/haiwangram/models"
)

type commentMapper struct {
}

func NewCommentMapper() *commentMapper {
	return &commentMapper{}
}

func (m *commentMapper) ToCommentResponse(request *models.Comment) *comment.CommentResponse {
	return &comment.CommentResponse{
		ID:      request.ID,
		Message: request.Message,
		UserID:  request.UserID,
		PhotoID: request.PhotoID,
	}
}

func (m *commentMapper) ToCommentWithRelationResponse(request *models.Comment) *comment.CommentWithRelationResponse {
	return &comment.CommentWithRelationResponse{
		ID:      request.ID,
		Message: request.Message,
		UserID:  request.UserID,
		PhotoID: request.PhotoID,
		Photo:   *NewPhotoMapper().ToPhotoResponse(&request.Photo),
		User:    *NewUserMapper().ToUserResponse(&request.User),
	}
}

func (m *commentMapper) ToCommentWithRelationResponses(requests *[]models.Comment) []comment.CommentWithRelationResponse {
	var responses []comment.CommentWithRelationResponse
	for _, request := range *requests {
		responses = append(responses, *m.ToCommentWithRelationResponse(&request))
	}

	if len(responses) > 0 {
		return responses
	}

	return []comment.CommentWithRelationResponse{}
}
