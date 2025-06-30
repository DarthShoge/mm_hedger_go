package risk

import (
	"testing"

	"github.com/darthshoge/mm_hedger_go/pkg/types"
)

func TestApplyTradesUpdates(t *testing.T) {
	m := NewManager(10, 100)
	trades := []*types.Trade{
		{Side: types.SideBuy, Price: 100, Quantity: 1},
		{Side: types.SideSell, Price: 110, Quantity: 1},
	}
	m.ApplyTrades(trades)
	if inv := m.Inventory(); inv != 0 {
		t.Fatalf("expected inventory 0, got %f", inv)
	}
	if pnl := m.PnL(); pnl != 10 {
		t.Fatalf("expected pnl 10, got %f", pnl)
	}
	if m.ShouldHalt() {
		t.Fatalf("should not halt")
	}
}

func TestInventoryLimitTriggersHalt(t *testing.T) {
	m := NewManager(5, 100)
	trades := []*types.Trade{{Side: types.SideBuy, Price: 100, Quantity: 6}}
	m.ApplyTrades(trades)
	if !m.ShouldHalt() {
		t.Fatalf("expected halt on inventory breach")
	}
}

func TestLossLimitTriggersHalt(t *testing.T) {
	m := NewManager(10, 50)
	trades := []*types.Trade{
		{Side: types.SideBuy, Price: 100, Quantity: 1},
		{Side: types.SideSell, Price: 40, Quantity: 1},
	}
	m.ApplyTrades(trades)
	if !m.ShouldHalt() {
		t.Fatalf("expected halt on loss breach")
	}
}

func TestPnlCalculation(t *testing.T) {
	m := NewManager(10, 100)
	trades := []*types.Trade{
		{Side: types.SideBuy, Price: 100, Quantity: 1},
		{Side: types.SideSell, Price: 120, Quantity: 1}, // this should give a profit of 20
		{Side: types.SideBuy, Price: 120, Quantity: 1}, // this should be flat 
		{Side: types.SideSell, Price: 110, Quantity: 1}, // this should give a loss of 10
	}
	m.ApplyTrades(trades)
	if m.PnL() != 10 {
		t.Fatalf("expected pnl 10, got %f", m.PnL())
	}
}

func TestInventoryCalculation(t *testing.T) {
	m := NewManager(10, 100)
	trades := []*types.Trade{
		{Side: types.SideBuy, Price: 100, Quantity: 5},
		{Side: types.SideSell, Price: 110, Quantity: 3},
	}
	m.ApplyTrades(trades)
	if inv := m.Inventory(); inv != 2 {
		t.Fatalf("expected inventory 2, got %f", inv)
	}
}
