package responses

type LikeResponse struct {
	ID      int `json:"id"`
	PhotoID int `json:"photo_id"`
	UserID  int `json:"user_id"`
}

type LikeWithRelationResponse struct {
	ID      int           `json:"id"`
	PhotoID int           `json:"photo_id"`
	UserID  int           `json:"user_id"`
	User    UserResponse  `json:"user"`
	Photo   PhotoResponse `json:"photo"`
}
