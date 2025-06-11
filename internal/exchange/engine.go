package exchange

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/darthshoge/mm_hedger_go/pkg/types"
)

// Engine represents a simple order matching engine.
type Engine struct {
	mu     sync.Mutex
	bids   []*types.Order
	asks   []*types.Order
	orders map[string]*types.Order
}

// NewEngine creates a new Engine.
func NewEngine() *Engine {
	return &Engine{
		bids:   make([]*types.Order, 0),
		asks:   make([]*types.Order, 0),
		orders: make(map[string]*types.Order),
	}
}

// Submit adds an order to the book and attempts to match it.
func (e *Engine) Submit(o *types.Order) ([]*types.Trade, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	o.ID = e.randomID()
	o.Timestamp = time.Now()
	trades := e.matchOrder(o)
	if o.Quantity > 0 && o.Type == types.OrderTypeLimit {
		e.addToBook(o)
		e.orders[o.ID] = o
	}

	return trades, nil
}

// Cancel removes an existing order from the book.
func (e *Engine) Cancel(id string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	ord, ok := e.orders[id]
	if !ok {
		return false
	}
	switch ord.Side {
	case types.SideBuy:
		e.bids = removeOrder(e.bids, id)
	case types.SideSell:
		e.asks = removeOrder(e.asks, id)
	}
	delete(e.orders, id)
	return true
}

// Match matches resting orders in the book until no more trades can occur.
func (e *Engine) Match() []*types.Trade {
	e.mu.Lock()
	defer e.mu.Unlock()

	trades := []*types.Trade{}
	for {
		if len(e.bids) == 0 || len(e.asks) == 0 {
			break
		}
		bestBid := e.bids[0]
		bestAsk := e.asks[0]
		if bestBid.Price < bestAsk.Price {
			break
		}
		qty := min(bestBid.Quantity, bestAsk.Quantity)
		trade := &types.Trade{
			ID:        e.randomID(),
			OrderID:   bestBid.ID,
			Side:      types.SideBuy,
			Price:     bestAsk.Price,
			Quantity:  qty,
			Timestamp: time.Now(),
		}
		trades = append(trades, trade)
		bestBid.Quantity -= qty
		bestAsk.Quantity -= qty
		if bestBid.Quantity == 0 {
			e.bids = e.bids[1:]
			delete(e.orders, bestBid.ID)
		}
		if bestAsk.Quantity == 0 {
			e.asks = e.asks[1:]
			delete(e.orders, bestAsk.ID)
		}
	}
	return trades
}

func (e *Engine) matchOrder(o *types.Order) []*types.Trade {
	var trades []*types.Trade
	for o.Quantity > 0 {
		if o.Side == types.SideBuy {
			if len(e.asks) == 0 {
				break
			}
			best := e.asks[0]
			if o.Type == types.OrderTypeLimit && o.Price < best.Price {
				break
			}
			qty := min(o.Quantity, best.Quantity)
			trade := &types.Trade{
				ID:        e.randomID(),
				OrderID:   o.ID,
				Side:      o.Side,
				Price:     best.Price,
				Quantity:  qty,
				Timestamp: time.Now(),
			}
			trades = append(trades, trade)
			o.Quantity -= qty
			best.Quantity -= qty
			if best.Quantity == 0 {
				e.asks = e.asks[1:]
				delete(e.orders, best.ID)
			}
		} else {
			if len(e.bids) == 0 {
				break
			}
			best := e.bids[0]
			if o.Type == types.OrderTypeLimit && o.Price > best.Price {
				break
			}
			qty := min(o.Quantity, best.Quantity)
			trade := &types.Trade{
				ID:        e.randomID(),
				OrderID:   o.ID,
				Side:      o.Side,
				Price:     best.Price,
				Quantity:  qty,
				Timestamp: time.Now(),
			}
			trades = append(trades, trade)
			o.Quantity -= qty
			best.Quantity -= qty
			if best.Quantity == 0 {
				e.bids = e.bids[1:]
				delete(e.orders, best.ID)
			}
		}
	}
	return trades
}

func (e *Engine) addToBook(o *types.Order) {
	switch o.Side {
	case types.SideBuy:
		e.bids = append(e.bids, o)
		sort.SliceStable(e.bids, func(i, j int) bool {
			return e.bids[i].Price > e.bids[j].Price
		})
	case types.SideSell:
		e.asks = append(e.asks, o)
		sort.SliceStable(e.asks, func(i, j int) bool {
			return e.asks[i].Price < e.asks[j].Price
		})
	}
}

func removeOrder(list []*types.Order, id string) []*types.Order {
	for i, o := range list {
		if o.ID == id {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func (e *Engine) randomID() string {
	return fmt.Sprintf("%d", rand.Int63())
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
