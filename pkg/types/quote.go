package types

import "time"

 // Quote represents a market quote with bid/ask prices and sizes.

// MidPrice returns the midpoint between Bid and Ask.
// If either Bid or Ask is zero, it returns zero.
func (q *Quote) MidPrice() float64 {
	if q.Bid == 0 || q.Ask == 0 {
		return 0
	}
	return (q.Bid + q.Ask) / 2
}
type Quote struct {
	Bid       float64
	Ask       float64
	BidSize   float64
	AskSize   float64
	Timestamp time.Time
}
