package params

import "time"

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}

type AuditTrails struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
