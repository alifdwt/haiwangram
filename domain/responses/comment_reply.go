package responses

type CommentReplyResponse struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
}

type CommentReplyWithRelationResponse struct {
	ID        int             `json:"id"`
	Message   string          `json:"message"`
	CommentID int             `json:"comment_id"`
	UserID    int             `json:"user_id"`
	Comment   CommentResponse `json:"comment"`
	User      UserResponse    `json:"user"`
}
