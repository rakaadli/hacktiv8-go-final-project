package models

type Comment struct {
	AuditTrails
	UserID  int `json:"user_id"`
	User    *User
	PhotoID int `json:"photo_id"`
	Photo   *Photo
	Message string `gorm:"not null;" json:"message" valid:"required~ Message is required"`
}
