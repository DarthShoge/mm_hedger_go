package hedge

import (
	"time"

	"github.com/darthshoge/mm_hedger_go/internal/exchange"
	"github.com/darthshoge/mm_hedger_go/pkg/types"
)

// Hedger mirrors trades executed on one exchange onto another with a fixed latency.
// It is not thread-safe and should be used from a single goroutine.
type Hedger struct {
	Engine  *exchange.Engine
	Latency time.Duration
}

// NewHedger returns a Hedger that sends orders to the given engine after the specified latency.
func NewHedger(engine *exchange.Engine, latency time.Duration) *Hedger {
	return &Hedger{Engine: engine, Latency: latency}
}

// Hedge submits offsetting market orders on the hedge exchange to neutralize the provided trades.
// Returns the trades generated on the hedge exchange.
func (h *Hedger) Hedge(trades []*types.Trade) ([]*types.Trade, error) {
	var hedgeTrades []*types.Trade
	for _, t := range trades {
		// opposite side of original trade
		side := types.SideSell
		if t.Side == types.SideSell {
			side = types.SideBuy
		}

		order := &types.Order{
			Side:     side,
			Type:     types.OrderTypeMarket,
			Quantity: t.Quantity,
		}

		if h.Latency > 0 {
			time.Sleep(h.Latency)
		}

		ts, err := h.Engine.Submit(order)
		if err != nil {
			return hedgeTrades, err
		}
		hedgeTrades = append(hedgeTrades, ts...)
	}
	return hedgeTrades, nil
}
