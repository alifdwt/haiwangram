package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses"
	"github.com/alifdwt/haiwangram/models"
)

type likeMapper struct {
}

func NewLikeMapper() *likeMapper {
	return &likeMapper{}
}

func (m *likeMapper) ToLikeResponse(request *models.Like) *responses.LikeResponse {
	return &responses.LikeResponse{
		ID:      request.ID,
		PhotoID: request.PhotoID,
		UserID:  request.UserID,
	}
}

func (m *likeMapper) ToLikeWithRelationResponse(request *models.Like) *responses.LikeWithRelationResponse {
	return &responses.LikeWithRelationResponse{
		ID:      request.ID,
		PhotoID: request.PhotoID,
		UserID:  request.UserID,
		Photo:   *NewPhotoMapper().ToPhotoResponse(&request.Photo),
		User:    *NewUserMapper().ToUserResponse(&request.User),
	}
}

func (m *likeMapper) ToLikeWithRelationResponses(requests *[]models.Like) []responses.LikeWithRelationResponse {
	var likeResponses []responses.LikeWithRelationResponse
	for _, request := range *requests {
		likeResponses = append(likeResponses, *m.ToLikeWithRelationResponse(&request))
	}

	if len(likeResponses) > 0 {
		return likeResponses
	}

	return []responses.LikeWithRelationResponse{}
}
