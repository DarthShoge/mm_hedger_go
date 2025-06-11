package exchange

import (
	"math/rand"
	"time"

	"github.com/darthshoge/mm_hedger_go/pkg/types"
)

// PriceFeed defines the interface for generating market quotes.
type PriceFeed interface {
	// Step advances the feed one tick and returns a new market quote.
	Step() *types.Quote
}

// RandomWalkFeed implements PriceFeed using a simple random walk model.
type RandomWalkFeed struct {
	Price float64
	Mu    float64
	Sigma float64
	rand  *rand.Rand
}

// NewRandomWalkFeed returns a random walk price feed.
func NewRandomWalkFeed(initial, mu, sigma float64) *RandomWalkFeed {
	return &RandomWalkFeed{
		Price: initial,
		Mu:    mu,
		Sigma: sigma,
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Step advances the price using a random walk and returns a new quote.
func (p *RandomWalkFeed) Step() *types.Quote {
	p.Price += p.Mu + p.Sigma*p.rand.NormFloat64()
	return &types.Quote{
		Bid:       p.Price - 0.5,
		Ask:       p.Price + 0.5,
		BidSize:   1,
		AskSize:   1,
		Timestamp: time.Now(),
	}
}

// StaticPriceFeed always returns the same price.
type StaticPriceFeed struct {
	Quote types.Quote
}

// NewStaticPriceFeed returns a StaticPriceFeed with the given price.
func NewStaticPriceFeed(price float64) *StaticPriceFeed {
	return &StaticPriceFeed{
		Quote: types.Quote{
			Bid:       price - 0.5,
			Ask:       price + 0.5,
			BidSize:   1,
			AskSize:   1,
			Timestamp: time.Now(),
		},
	}
}

// Step returns the stored quote and updates the timestamp.
func (s *StaticPriceFeed) Step() *types.Quote {
	s.Quote.Timestamp = time.Now()
	q := s.Quote
	return &q
}
