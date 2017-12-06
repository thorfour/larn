package items

import (
	"fmt"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

const (
	goldPileRune = '*'
)

type GoldPile struct {
	Amount int
	DefaultItem
}

// Rune implements the io.Runeable interface
func (g *GoldPile) Rune() rune {
	if g.Visibility {
		return goldPileRune
	} else {
		return invisibleRune
	}
}

// Log implements the Disaplceable interface
func (g *GoldPile) Log() string {
	return fmt.Sprintf("You have found some gold worth %v", g.Amount)
}

// PickUp implements the item interface
func (g *GoldPile) PickUp(s *stats.Stats) {
	s.Gold += uint(g.Amount)
}

// Drop implements the item interface
func (g *GoldPile) Drop(s *stats.Stats) {}
