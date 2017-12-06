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

type Belt struct {
	Attribute int
	DefaultItem
}

// String implements the Item interface
func (b *Belt) String() string {
	if b.Attribute == 0 {
		return "a belt of striking"
	} else if b.Attribute > 0 {
		return fmt.Sprintf("a belt of striking + %v", b.Attribute)
	} else {
		return fmt.Sprintf("a belt of striking - %v", int(math.Abs(float64(b.Attribute))))
	}
}

// Rune implements the io.Runeable interface
func (b *Belt) Rune() rune {
	return beltRune
}

// Log implements the Loggable interface
func (b *Belt) Log() string {
	return b.String()
}

// PickUp implements the Item interface
func (b *Belt) PickUp(s *stats.Stats) {
	s.Wc += (2 + (b.Attribute << 1))
}

// Drop implements the Item interface
func (b *Belt) Drop(s *stats.Stats) {
	s.Wc -= (2 + (b.Attribute << 1))
}

// Wield implements the Weapon interface
func (b *Belt) Wield(s *stats.Stats) {
	s.Wc += (beltBase + b.Attribute)
}

// Disarm implements the Weapon interface
func (b *Belt) Disarm(s *stats.Stats) {
	s.Wc -= (beltBase + b.Attribute)
}
