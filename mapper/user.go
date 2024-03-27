package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses/user"
	"github.com/alifdwt/haiwangram/models"
)

type userMapper struct {
}

func NewUserMapper() *userMapper {
	return &userMapper{}
}

func (m *userMapper) ToUserResponse(request *models.User) *user.UserResponse {
	return &user.UserResponse{
		ID:              request.ID,
		Email:           request.Email,
		Username:        request.Username,
		FullName:        request.FullName,
		ProfileImageURL: request.ProfileImageURL,
		Description:     request.Description,
		// BirthDate:       request.BirthDate,
	}
}
