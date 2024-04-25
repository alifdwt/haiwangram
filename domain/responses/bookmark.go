package responses

type BookmarkResponse struct {
	ID      int `json:"id"`
	PhotoID int `json:"photo_id"`
	UserID  int `json:"user_id"`
}

type BookmarkWithRelationResponse struct {
	ID      int                       `json:"id"`
	PhotoID int                       `json:"photo_id"`
	UserID  int                       `json:"user_id"`
	Photo   PhotoWithRelationResponse `json:"photo"`
	// User    UserResponse  `json:"user"`
}
