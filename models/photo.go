package models

type Photo struct {
	AuditTrails
	Title    string `gorm:"not null;" json:"title" valid:"required~ Your title is required"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null;" json:"photo_url" valid:"required~ Your Photo URL is required"`
	UserID   int
	User     *User
}
