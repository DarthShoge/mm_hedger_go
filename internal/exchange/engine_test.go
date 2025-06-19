package exchange

import (
	"testing"

	"github.com/darthshoge/mm_hedger_go/pkg/types"
)

func TestEngineSubmitMatch(t *testing.T) {
	e := NewEngine()

	sell := &types.Order{Side: types.SideSell, Type: types.OrderTypeLimit, Price: 101, Quantity: 1}
	if trades, _ := e.Submit(sell); len(trades) != 0 {
		t.Fatalf("expected no trades on first submit")
	}

	buy := &types.Order{Side: types.SideBuy, Type: types.OrderTypeLimit, Price: 102, Quantity: 1}
	trades, _ := e.Submit(buy)
	if len(trades) != 1 {
		t.Fatalf("expected one trade, got %d", len(trades))
	}
	if trades[0].Price != 101 {
		t.Fatalf("expected price 101, got %f", trades[0].Price)
	}
	if len(e.bids) != 0 || len(e.asks) != 0 {
		t.Fatalf("orderbook should be empty after match")
	}
}

func TestEngineCancel(t *testing.T) {
	e := NewEngine()
	ord := &types.Order{Side: types.SideSell, Type: types.OrderTypeLimit, Price: 100, Quantity: 1}
	e.Submit(ord)
	if !e.Cancel(ord.ID) {
		t.Fatalf("expected cancel to succeed")
	}
	if len(e.asks) != 0 {
		t.Fatalf("order should be removed")
	}
}

func TestEngineMarketOrder(t *testing.T) {
	e := NewEngine()
	e.Submit(&types.Order{Side: types.SideSell, Type: types.OrderTypeLimit, Price: 100, Quantity: 1})
	buy := &types.Order{Side: types.SideBuy, Type: types.OrderTypeMarket, Quantity: 1}
	trades, _ := e.Submit(buy)
	if len(trades) != 1 {
		t.Fatalf("expected one trade for market order")
	}
	if len(e.asks) != 0 {
		t.Fatalf("orderbook should be empty")
	}
}

func TestEngineOrderWithNoMatch(t *testing.T) {
	e := NewEngine()
	sell := &types.Order{Side: types.SideSell, Type: types.OrderTypeLimit, Price: 100, Quantity: 1}
	if trades, _ := e.Submit(sell); len(trades) != 0 {
		t.Fatalf("expected no trades on first submit")
	}

	buy := &types.Order{Side: types.SideBuy, Type: types.OrderTypeLimit, Price: 99, Quantity: 1}
	if trades, _ := e.Submit(buy); len(trades) != 0 {
		t.Fatalf("expected no trades for unmatched order")
	}
	if len(e.bids) != 1 || len(e.asks) != 1 {
		t.Fatalf("orderbook should contain one bid and one ask")
	}
}
