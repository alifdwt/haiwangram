package responses

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
	User     UserResponse      `json:"user"`
	Comments []CommentResponse `json:"comments"`
}
