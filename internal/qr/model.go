package qr

import (
	"time"

	"github.com/google/uuid"
)

type QR struct {
	ID      uuid.UUID `json:"id" db:"id"`
	Slug    string    `json:"slug" db:"slug"`
	Content string    `json:"content" db:"content"`
	Type    string    `json:"type" db:"type"`

	ScanCount int  `json:"scan_count" db:"scan_count"`
	IsActive  bool `json:"is_active" db:"is_active"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
