package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses/photo"
	"github.com/alifdwt/haiwangram/models"
)

type photoMapper struct {
}

func NewPhotoMapper() *photoMapper {
	return &photoMapper{}
}

func (m *photoMapper) ToPhotoResponse(request *models.Photo) *photo.PhotoResponse {
	return &photo.PhotoResponse{
		ID:       request.ID,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   request.UserID,
	}
}

func (m *photoMapper) ToPhotoWithRelationResponse(request *models.Photo) *photo.PhotoWithRelationResponse {
	return &photo.PhotoWithRelationResponse{
		ID:       request.ID,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   request.UserID,
		User:     *NewUserMapper().ToUserResponse(&request.User),
	}
}

func (m *photoMapper) ToPhotoWithRelationResponses(requests *[]models.Photo) []photo.PhotoWithRelationResponse {
	var responses []photo.PhotoWithRelationResponse

	for _, request := range *requests {
		response := m.ToPhotoWithRelationResponse(&request)
		responses = append(responses, *response)
	}

	if len(responses) > 0 {
		return responses
	}

	return []photo.PhotoWithRelationResponse{}
}
