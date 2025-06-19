package exchange

import (
	"testing"
	"time"
)

func TestRandomWalkFeedStep(t *testing.T) {
	pf := NewRandomWalkFeed(100, 1, 0)
	q := pf.Step()
	if q.Bid != 100+1-0.5 {
		t.Fatalf("unexpected bid price")
	}
	if pf.Price != 101 {
		t.Fatalf("price not updated")
	}
}

func TestQuoteMidPrice(t *testing.T) {
	pf := NewRandomWalkFeed(100, 1, 0)
	q := pf.Step()
	mid := q.MidPrice()
	if mid != 100+1 {
		t.Fatalf("unexpected mid price: got %f, want %f", mid, float64(100+1))
	}
	if q.Bid == 0 || q.Ask == 0 {
		t.Fatalf("bid or ask should not be zero")
	}
}

func TestStaticPriceFeed(t *testing.T) {
	pf := NewStaticPriceFeed(50)
	q1 := pf.Step()
	if q1.Bid != 49.5 || q1.Ask != 50.5 {
		t.Fatalf("static feed returned wrong quote")
	}
	time.Sleep(time.Millisecond)
	q2 := pf.Step()
	if !q2.Timestamp.After(q1.Timestamp) {
		t.Fatalf("timestamp not updated")
	}
}
