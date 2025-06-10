package types

import "testing"

func TestOrderTypeString(t *testing.T) {
	if OrderTypeLimit.String() != "limit" {
		t.Fatalf("expected limit, got %s", OrderTypeLimit.String())
	}
	if OrderTypeMarket.String() != "market" {
		t.Fatalf("expected market, got %s", OrderTypeMarket.String())
	}
	if OrderType(100).String() != "unknown" {
		t.Fatalf("expected unknown for invalid order type")
	}
}
