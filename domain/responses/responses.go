package responses

import "github.com/alifdwt/haiwangram/domain/responses/user"

/* --------------- Comment Responses --------------- */
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
	User    user.UserResponse      `json:"user"`
	Photo   PhotoResponse          `json:"photo"`
	Replies []CommentReplyResponse `json:"replies"`
}

/* --------------- Comment Reply Responses --------------- */
type CommentReplyResponse struct {
	ID        int    `json:"id"`
	Message   string `json:"message"`
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
}

type CommentReplyWithRelationResponse struct {
	ID        int               `json:"id"`
	Message   string            `json:"message"`
	CommentID int               `json:"comment_id"`
	UserID    int               `json:"user_id"`
	Comment   CommentResponse   `json:"comment"`
	User      user.UserResponse `json:"user"`
}

/* --------------- Photo Responses --------------- */
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
	Comments []CommentResponse `json:"comments"`
}
