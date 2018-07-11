package items

import (
	"fmt"
	"math"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

const beltRune = '{'

const (
	beltBase = 7
)

// Belt is a belt of striking
type Belt struct {
	DefaultAttribute
	DefaultItem
}

// String implements the Item interface
func (b *Belt) String() string {
	if b.Attr() == 0 {
		return "a belt of striking"
	} else if b.Attr() > 0 {
		return fmt.Sprintf("a belt of striking + %v", b.Attr())
	} else {
		return fmt.Sprintf("a belt of striking - %v", int(math.Abs(float64(b.Attr()))))
	}
}

// Rune implements the io.Runeable interface
func (b *Belt) Rune() rune {
	if b.Visibility {
		return beltRune
	}
	return invisibleRune
}

// Log implements the Loggable interface
func (b *Belt) Log() string {
	return b.String()
}

// PickUp implements the Item interface
func (b *Belt) PickUp(s *stats.Stats) {
	s.Wc += (2 + (b.Attr() << 1))
}

// Drop implements the Item interface
func (b *Belt) Drop(s *stats.Stats) {
	s.Wc -= (2 + (b.Attr() << 1))
}

// Wield implements the Weapon interface
func (b *Belt) Wield(s *stats.Stats) {
	s.Wc += (beltBase + b.Attr())
}

// Disarm implements the Weapon interface
func (b *Belt) Disarm(s *stats.Stats) {
	s.Wc -= (beltBase + b.Attr())
}
