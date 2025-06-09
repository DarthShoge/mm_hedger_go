package types

import "time"

// Order represents an order in the system.
type Order struct {
	ID        string
	Side      Side
	Type      OrderType
	Price     float64
	Quantity  float64
	Timestamp time.Time
}
