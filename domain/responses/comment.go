package responses

type CommentResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
	UserID  int    `json:"user_id"`
}

type CommentWithRelationResponse struct {
	ID      int                    `json:"id"`
	Message string                 `json:"message"`
	PhotoID int                    `json:"photo_id"`
	UserID  int                    `json:"user_id"`
	User    UserResponse           `json:"user"`
	Photo   PhotoResponse          `json:"photo"`
	Replies []CommentReplyResponse `json:"replies"`
}
