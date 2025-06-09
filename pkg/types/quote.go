package types

import "time"

// Quote represents a market quote with bid/ask prices and sizes.
type Quote struct {
	Bid       float64
	Ask       float64
	BidSize   float64
	AskSize   float64
	Timestamp time.Time
}
