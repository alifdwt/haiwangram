package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses/photo"
	"github.com/alifdwt/haiwangram/domain/responses/user"
	"github.com/alifdwt/haiwangram/models"
)

type UserMapping interface {
	ToUserResponse(request *models.User) *user.UserResponse
}

type PhotoMapping interface {
	ToPhotoResponse(request *models.Photo) *photo.PhotoResponse
	ToPhotoWithRelationResponse(request *models.Photo) *photo.PhotoWithRelationResponse
	ToPhotoWithRelationResponses(requests *[]models.Photo) []photo.PhotoWithRelationResponse
}
