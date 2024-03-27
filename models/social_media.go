package models

import "time"

type SocialMedia struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"size:50"`
	SocialMediaURL string    `json:"social_media_url" gorm:"type:text;not null"`
	UserID         uint      `json:"user_id" gorm:"not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"type:timestamp"`
	User           User
}

func (SocialMedia) TableName() string {
	return "social_medias"
}
