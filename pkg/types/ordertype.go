package types

// OrderType denotes whether an order is a limit or market order.
type OrderType int

const (
	OrderTypeLimit OrderType = iota + 1
	OrderTypeMarket
)

func (o OrderType) String() string {
	switch o {
	case OrderTypeLimit:
		return "limit"
	case OrderTypeMarket:
		return "market"
	default:
		return "unknown"
	}
}
