package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/models"
)

type photoMapper struct {
}

func NewPhotoMapper() *photoMapper {
	return &photoMapper{}
}

func (m *photoMapper) ToPhotoResponse(request *models.Photo) *responses.PhotoResponse {
	return &responses.PhotoResponse{
		ID:        request.ID,
		Title:     request.Title,
		Caption:   request.Caption,
		PhotoURL:  request.PhotoURL,
		UserID:    request.UserID,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
}

func (m *photoMapper) ToPhotoWithRelationResponse(request *models.Photo) *responses.PhotoWithRelationResponse {
	var comments []responses.CommentResponse
	if request.Comments != nil {
		for _, comment := range request.Comments {
			comments = append(comments, *NewCommentMapper().ToCommentResponse(&comment))
		}
	} else {
		comments = []responses.CommentResponse{}
	}

	var likes []responses.LikeResponse
	if request.Likes != nil {
		for _, like := range request.Likes {
			likes = append(likes, *NewLikeMapper().ToLikeResponse(&like))
		}
	} else {
		likes = []responses.LikeResponse{}
	}

	return &responses.PhotoWithRelationResponse{
		ID:        request.ID,
		Title:     request.Title,
		Caption:   request.Caption,
		PhotoURL:  request.PhotoURL,
		UserID:    request.UserID,
		User:      *NewUserMapper().ToUserResponse(&request.User),
		Comments:  comments,
		Likes:     likes,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
}

func (m *photoMapper) ToPhotoWithRelationResponses(requests *[]models.Photo) []responses.PhotoWithRelationResponse {
	var photoResponses []responses.PhotoWithRelationResponse

	for _, request := range *requests {
		response := m.ToPhotoWithRelationResponse(&request)
		photoResponses = append(photoResponses, *response)
	}

	if len(photoResponses) > 0 {
		return photoResponses
	}

	return []responses.PhotoWithRelationResponse{}
}
