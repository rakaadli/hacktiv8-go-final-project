package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Title     string    `gorm:"not null;" json:"title" valid:"required~ Your title is required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null;" json:"photo_url" valid:"required~ Your Photo URL is required"`
	UserID    int
	User      *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	return errCreate
}
