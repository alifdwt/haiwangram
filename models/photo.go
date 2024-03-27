package models

import "time"

type Photo struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"size:100;not null"`
	Caption   string    `json:"caption" gorm:"size:200"`
	PhotoURL  string    `json:"photo_url" gorm:"type:text;not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
