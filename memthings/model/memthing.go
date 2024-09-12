package model

import (
	"time"
)

// Memthing defines the memory item structure
type Memthing struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Tags           []string  `json:"tags"`
	Classification string    `json:"classification"`
	Grade          float64   `json:"grade"`
	CreatedAt      time.Time `json:"created_at"`
}
