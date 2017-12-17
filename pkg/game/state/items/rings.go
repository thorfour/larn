package items

import (
	"strconv"

	"github.com/thorfour/larn/pkg/game/state/stats"
)

type RingType int

const ringRune = '|'

const (
	Regen RingType = iota
	ExtraRegen
	Protection
	Strength
	Clever
	Damage
	Dexterity
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

type Ring struct {
	Type      RingType // the type of ring
	Attribute int      // the ring modifier
	DefaultItem
}

// Rune implements the io.Runeable interface
func (r *Ring) Rune() rune {
	if r.Visibility {
		return ringRune
	} else {
		return invisibleRune
	}
}

// Log implements the Loggable interface
func (r *Ring) Log() string {
	return "You have found a " + r.String()
}

// String implements the Item interface
func (r *Ring) String() string {
	if r.Attribute < 0 {
		return ringName[r.Type] + " " + strconv.Itoa(r.Attribute)
	} else if r.Attribute > 0 {
		return ringName[r.Type] + " +" + strconv.Itoa(r.Attribute)
	}
	return ringName[r.Type]
}

// FIXME can you wear or wield rings?

// PickUp implements the Item interface
func (r *Ring) PickUp(s *stats.Stats) {
	switch r.Type {
	case Strength:
		s.Str += uint(1 + r.Attribute)
	case Clever:
		s.Intelligence += uint(1 + r.Attribute)
	case Dexterity:
		s.Dex += uint(1 + r.Attribute)
	}
}

// Drop implements the Item interface
func (r *Ring) Drop(s *stats.Stats) {
	switch r.Type {
	case Strength:
		s.Str -= uint(1 + r.Attribute)
	case Clever:
		s.Intelligence -= uint(1 + r.Attribute)
	case Dexterity:
		s.Dex -= uint(1 + r.Attribute)
	}
}
