package mapper

import (
	"github.com/alifdwt/haiwangram/domain/responses/user"
	"github.com/alifdwt/haiwangram/models"
)

type UserMapping interface {
	ToUserResponse(request *models.User) *user.UserResponse
}
