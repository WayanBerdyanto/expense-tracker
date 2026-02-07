package models

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description" binding: "required"`
	Amount      float64   `json:"amount" binding: "required,gt=0"`
	Category    string    `json:"category" binding: "required"`
	CreatedAt   time.Time `json:"created_at"`
}
