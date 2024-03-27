package comment

import (
	"github.com/alifdwt/haiwangram/domain/responses/photo"
	"github.com/alifdwt/haiwangram/domain/responses/user"
)

type CommentResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
}

type CommentWithRelationResponse struct {
	ID      int                 `json:"id"`
	Message string              `json:"message"`
	PhotoID int                 `json:"photo_id"`
	UserID  int                 `json:"user_id"`
	User    user.UserResponse   `json:"user"`
	Photo   photo.PhotoResponse `json:"photo"`
}
