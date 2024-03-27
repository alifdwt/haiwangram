package models

import "time"

type CommentReply struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id" gorm:"not null"`
	CommentID int       `json:"comment_id" gorm:"not null"`
	Message   string    `json:"message" gorm:"size:200;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
	Comment   Comment   `json:"comment"`
}
