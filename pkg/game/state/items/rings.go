package items

import (
	"strconv"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

// RingType is the type of ring
type RingType int

const ringRune = '|'

const (
	// Regen ring of regeneration
	Regen RingType = iota
	// ExtraRegen ring of extra regeneration
	ExtraRegen
	// Protection ring of protection
	Protection
	// Strength ring of strength
	Strength
	// Clever ring of cleverness
	Clever
	// Damage ring of extra damage
	Damage
	// Dexterity ring of dexterity
	Dexterity
	// Energy ring
	Energy
)

var ringName = map[RingType]string{
	Regen:      "ring of regeneration",
	ExtraRegen: "ring of extra regeneration",
	Protection: "ring of protection",
	Strength:   "ring of strength",
	Clever:     "ring of cleverness",
	Damage:     "ring of damage",
	Dexterity:  "ring of dexterity",
	Energy:     "energy ring",
}

// Ring is a wearable ring item
type Ring struct {
	Type RingType // the type of ring
	DefaultAttribute
	DefaultItem
}

// Rune implements the io.Runeable interface
func (r *Ring) Rune() rune {
	if r.Visibility {
		return ringRune
	}
	return invisibleRune
}

// Log implements the Loggable interface
func (r *Ring) Log() string {
	return "You have found a " + r.String()
}

// String implements the Item interface
func (r *Ring) String() string {
	if r.Attr() < 0 {
		return ringName[r.Type] + " " + strconv.Itoa(r.Attr())
	} else if r.Attr() > 0 {
		return ringName[r.Type] + " +" + strconv.Itoa(r.Attr())
	}
	return ringName[r.Type]
}

// FIXME can you wear or wield rings?

// PickUp implements the Item interface
func (r *Ring) PickUp(s *stats.Stats) {
	switch r.Type {
	case Strength:
		s.Str += uint(1 + r.Attr())
	case Clever:
		s.Intelligence += uint(1 + r.Attr())
	case Dexterity:
		s.Dex += uint(1 + r.Attr())
	}
}

// Drop implements the Item interface
func (r *Ring) Drop(s *stats.Stats) {
	switch r.Type {
	case Strength:
		s.Str -= uint(1 + r.Attr())
	case Clever:
		s.Intelligence -= uint(1 + r.Attr())
	case Dexterity:
		s.Dex -= uint(1 + r.Attr())
	}
}
