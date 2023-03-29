package domain

import (
	"time"
)

type Model struct {
	CreatedAt *time.Time `gorm:"type:timestamp" json:"createdAt,omitempty"`
}
