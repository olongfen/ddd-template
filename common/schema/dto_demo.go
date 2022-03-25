package schema

import (
	"time"
)

type DemoInfo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Message   string    `json:"message"`
}
