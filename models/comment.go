package models

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UserID    int       `json:"user_id"`
	User      *User
	PhotoID   int `json:"photo_id"`
	Photo     *Photo
	Message   string `gorm:"not null;" json:"message" valid:"required~ Message is required"`
}
