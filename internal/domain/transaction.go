package domain

import (
	"time"
)

type Transaction struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
