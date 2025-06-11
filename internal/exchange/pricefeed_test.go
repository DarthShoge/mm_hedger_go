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
