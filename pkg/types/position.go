package types

// Position represents a net position in a traded instrument.
type Position struct {
	Symbol   string
	Quantity float64
	AvgPrice float64
}
