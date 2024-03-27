package models

import "time"

type Comment struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	UserID    int            `json:"user_id" gorm:"not null"`
	PhotoID   int            `json:"photo_id" gorm:"not null"`
	Message   string         `json:"message" gorm:"size:200;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Replies   []CommentReply `json:"replies"`
	User      User           `json:"user"`
	Photo     Photo          `json:"photo"`
}
