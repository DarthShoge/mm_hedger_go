package risk

import (
	"errors"
	"math"
	"sync"

	"github.com/darthshoge/mm_hedger_go/pkg/types"
)

// ErrHalted indicates trading should halt due to risk limits being breached.
var ErrHalted = errors.New("trading halted")

// Manager tracks inventory and PnL and triggers a kill switch when limits are exceeded.
type Manager struct {
	mu sync.Mutex

	MaxInventory float64
	MaxLoss      float64

	inventory float64
	pnl       float64
	halted    bool
}

// NewManager returns a Manager with the specified limits.
func NewManager(maxInv, maxLoss float64) *Manager {
	return &Manager{MaxInventory: maxInv, MaxLoss: maxLoss}
}

// ApplyTrades updates inventory and PnL based on the given trades and checks limits.
func (m *Manager) ApplyTrades(trades []*types.Trade) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, t := range trades {
		if t.Side == types.SideBuy {
			m.inventory += t.Quantity
			m.pnl -= t.Price * t.Quantity
		} else {
			m.inventory -= t.Quantity
			m.pnl += t.Price * t.Quantity
		}
	}

	if math.Abs(m.inventory) > m.MaxInventory || m.pnl <= -m.MaxLoss {
		m.halted = true
	}
}

// Inventory returns the current net inventory.
func (m *Manager) Inventory() float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.inventory
}

// PnL returns the current profit and loss.
func (m *Manager) PnL() float64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.pnl
}

// ShouldHalt returns true if risk limits have been breached.
func (m *Manager) ShouldHalt() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.halted
}

// Reset clears state and allows trading again.
func (m *Manager) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.inventory = 0
	m.pnl = 0
	m.halted = false
}
