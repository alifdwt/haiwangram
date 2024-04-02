package responses

import "time"

type CommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentWithRelationResponse struct {
	ID        int                    `json:"id"`
	Message   string                 `json:"message"`
	PhotoID   int                    `json:"photo_id"`
	UserID    int                    `json:"user_id"`
	User      UserResponse           `json:"user"`
	Photo     PhotoResponse          `json:"photo"`
	Replies   []CommentReplyResponse `json:"replies"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
