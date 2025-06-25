package hedge

import (
	"testing"
	"time"

	"github.com/darthshoge/mm_hedger_go/internal/exchange"
	"github.com/darthshoge/mm_hedger_go/pkg/types"
)

func TestHedgerMirrorsBuyTrade(t *testing.T) {
	eng := exchange.NewEngine()
	// provide liquidity on both sides so market orders fill
	eng.Submit(&types.Order{Side: types.SideSell, Type: types.OrderTypeLimit, Price: 101, Quantity: 10})
	eng.Submit(&types.Order{Side: types.SideBuy, Type: types.OrderTypeLimit, Price: 99, Quantity: 10})

	h := NewHedger(eng, time.Millisecond)
	trade := &types.Trade{Side: types.SideBuy, Quantity: 1, Price: 100}

	start := time.Now()
	trades, err := h.Hedge([]*types.Trade{trade})
	if err != nil {
		t.Fatalf("hedge returned error: %v", err)
	}
	if len(trades) != 1 {
		t.Fatalf("expected 1 hedge trade, got %d", len(trades))
	}
	if trades[0].Side != types.SideSell {
		t.Fatalf("expected hedge side sell, got %v", trades[0].Side)
	}
	if time.Since(start) < time.Millisecond {
		t.Fatalf("latency not applied")
	}
}

func TestHedgerMirrorsSellTrade(t *testing.T) {
	eng := exchange.NewEngine()
	// provide liquidity on both sides so market orders fill
	eng.Submit(&types.Order{Side: types.SideSell, Type: types.OrderTypeLimit, Price: 101, Quantity: 10})
	eng.Submit(&types.Order{Side: types.SideBuy, Type: types.OrderTypeLimit, Price: 99, Quantity: 10})

	h := NewHedger(eng, time.Millisecond)
	trade := &types.Trade{Side: types.SideSell, Quantity: 1, Price: 100}

	start := time.Now()
	trades, err := h.Hedge([]*types.Trade{trade})
	if err != nil {
		t.Fatalf("hedge returned error: %v", err)
	}
	if len(trades) != 1 {
		t.Fatalf("expected 1 hedge trade, got %d", len(trades))
	}
	if trades[0].Side != types.SideBuy {
		t.Fatalf("expected hedge side buy, got %v", trades[0].Side)
	}
	if time.Since(start) < time.Millisecond {
		t.Fatalf("latency not applied")
	}
}
