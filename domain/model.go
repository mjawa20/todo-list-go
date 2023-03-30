package domain

import (
	"time"
)

type Model struct {
	CreatedAt *time.Time `gorm:"type:timestamp; default:CURRENT_TIMESTAMP(); column:createdAt" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `gorm:"type:timestamp; default:CURRENT_TIMESTAMP(); column:updatedAt" json:"updatedAt,omitempty"`
}
