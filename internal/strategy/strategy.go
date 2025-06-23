package strategy

import "github.com/darthshoge/mm_hedger_go/pkg/types"

// QuoteStrategy generates bid and ask quotes around a mid price.
// Spread defines the half-spread from the mid price while InvSkew
// shifts quotes based on current inventory to encourage reversion
// towards zero.
type QuoteStrategy struct {
	BaseSize float64
	Spread   float64
	InvSkew  float64
}

// NewQuoteStrategy returns a QuoteStrategy with the provided
// parameters.
func NewQuoteStrategy(baseSize, spread, invSkew float64) *QuoteStrategy {
	return &QuoteStrategy{BaseSize: baseSize, Spread: spread, InvSkew: invSkew}
}

// GenerateQuotes returns bid and ask limit orders given the current
// mid price and inventory. Positive inventory shifts both prices down
// while negative inventory shifts them up.
func (s *QuoteStrategy) GenerateQuotes(mid, inventory float64) (bid, ask *types.Order) {
	adjust := s.InvSkew * inventory
	bidPrice := mid - s.Spread - adjust
	askPrice := mid + s.Spread - adjust

	bid = &types.Order{
		Side:     types.SideBuy,
		Type:     types.OrderTypeLimit,
		Price:    bidPrice,
		Quantity: s.BaseSize,
	}
	ask = &types.Order{
		Side:     types.SideSell,
		Type:     types.OrderTypeLimit,
		Price:    askPrice,
		Quantity: s.BaseSize,
	}
	return
}
