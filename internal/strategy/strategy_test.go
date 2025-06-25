package strategy

import "testing"

func TestGenerateQuotesZeroInventory(t *testing.T) {
	strat := NewQuoteStrategy(1, 0.5, 0.1)
	bid, ask := strat.GenerateQuotes(100, 0)
	if bid.Price != 99.5 {
		t.Fatalf("bid price wrong: %f", bid.Price)
	}
	if ask.Price != 100.5 {
		t.Fatalf("ask price wrong: %f", ask.Price)
	}
}

func TestGenerateQuotesInventorySkew(t *testing.T) {
	strat := NewQuoteStrategy(1, 0.5, 0.1)
	bid, ask := strat.GenerateQuotes(100, 10)
	// adjust = 1
	if bid.Price != 98.5 {
		t.Fatalf("bid price wrong with inventory: %f", bid.Price)
	}
	if ask.Price != 99.5 {
		t.Fatalf("ask price wrong with inventory: %f", ask.Price)
	}
}
