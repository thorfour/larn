package items

import (
	"strconv"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

const shieldRune = ']'

const (
	shieldWC = 8
	ShieldAC = 2
)

type Shield struct {
	Attribute int
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (s *Shield) Rune() rune {
	return shieldRune
}

// Log implements the Loggable interface
func (s *Shield) Log() string {
	return "You have found a " + s.String()
}

// String implements the Item interface
func (s *Shield) String() string {
	if s.Attribute < 0 {
		return "shield " + strconv.Itoa(s.Attribute)
	} else if s.Attribute > 0 {
		return "shield" + " +" + strconv.Itoa(s.Attribute)
	}

	return "shield"
}

// Wield implements the Weapon interface
func (s *Shield) Wield(c *stats.Stats) {
	c.Wc += (shieldWC + s.Attribute)
}

// Disarm implements the Weapon interface
func (s *Shield) Disarm(c *stats.Stats) {
	c.Wc -= (shieldWC + s.Attribute)
}

// Wear implements the Armor interface
func (s *Shield) Wear(c *stats.Stats) {
	c.Ac += (ShieldAC + s.Attribute)
}

// TakeOff implements the Armor interface
func (s *Shield) TakeOff(c *stats.Stats) {
	c.Ac -= (ShieldAC + s.Attribute)
}
