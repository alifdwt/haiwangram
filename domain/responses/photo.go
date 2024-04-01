package responses

import "time"

type PhotoResponse struct {
	ID        int       `json:"id"`
	Caption   string    `json:"caption"`
	Title     string    `json:"title"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoWithRelationResponse struct {
	ID        int               `json:"id"`
	Caption   string            `json:"caption"`
	Title     string            `json:"title"`
	PhotoURL  string            `json:"photo_url"`
	UserID    int               `json:"user_id"`
	User      UserResponse      `json:"user"`
	Comments  []CommentResponse `json:"comments"`
	Likes     []LikeResponse    `json:"likes"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
