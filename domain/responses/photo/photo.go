package photo

import "github.com/alifdwt/haiwangram/domain/responses/user"

type PhotoResponse struct {
	ID       int    `json:"id"`
	Caption  string `json:"caption"`
	Title    string `json:"title"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

type PhotoWithRelationResponse struct {
	ID       int               `json:"id"`
	Caption  string            `json:"caption"`
	Title    string            `json:"title"`
	PhotoURL string            `json:"photo_url"`
	UserID   int               `json:"user_id"`
	User     user.UserResponse `json:"user"`
}
