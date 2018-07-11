package items

import (
	"strconv"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

const shieldRune = ']'

const (
	shieldWC = 8
	shieldAC = 2
)

// Shield is a shield item
type Shield struct {
	DefaultAttribute
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (s *Shield) Rune() rune {
	if s.Visibility {
		return shieldRune
	}
	return invisibleRune
}

// Log implements the Loggable interface
func (s *Shield) Log() string {
	return "You have found a " + s.String()
}

// String implements the Item interface
func (s *Shield) String() string {
	if s.Attr() < 0 {
		return "shield " + strconv.Itoa(s.Attr())
	} else if s.Attr() > 0 {
		return "shield" + " +" + strconv.Itoa(s.Attr())
	}

	return "shield"
}

// Wield implements the Weapon interface
func (s *Shield) Wield(c *stats.Stats) {
	c.Wc += (shieldWC + s.Attr())
}

// Disarm implements the Weapon interface
func (s *Shield) Disarm(c *stats.Stats) {
	c.Wc -= (shieldWC + s.Attr())
}

// Wear implements the Armor interface
func (s *Shield) Wear(c *stats.Stats) {
	c.Ac += (shieldAC + s.Attr())
}

// TakeOff implements the Armor interface
func (s *Shield) TakeOff(c *stats.Stats) {
	c.Ac -= (shieldAC + s.Attr())
}
