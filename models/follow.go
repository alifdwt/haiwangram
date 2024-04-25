package models

import "time"

type Follow struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	FollowerID int       `json:"follower_id" gorm:"not null"`
	FollowedID int       `json:"followed_id" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Follower   User      `json:"follower" gorm:"foreignKey:FollowerID;references:ID"`
	Followed   User      `json:"followed" gorm:"foreignKey:FollowedID;references:ID"`
}
