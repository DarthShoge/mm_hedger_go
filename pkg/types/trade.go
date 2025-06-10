package types

import "time"

// Trade represents an executed trade.
type Trade struct {
	ID        string
	OrderID   string
	Side      Side
	Price     float64
	Quantity  float64
	Timestamp time.Time
}
