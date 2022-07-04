package models

import "time"

type AuditTrails struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
