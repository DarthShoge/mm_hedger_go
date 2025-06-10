package types

import "testing"

func TestSideString(t *testing.T) {
	if SideBuy.String() != "buy" {
		t.Fatalf("expected buy, got %s", SideBuy.String())
	}
	if SideSell.String() != "sell" {
		t.Fatalf("expected sell, got %s", SideSell.String())
	}
	if Side(100).String() != "unknown" {
		t.Fatalf("expected unknown for invalid side")
	}
}
