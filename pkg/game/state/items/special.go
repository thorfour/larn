package items

import "github.com/thorfour/larn/pkg/game/state/stats"

type SpecialType int

var specialRunes = map[SpecialType]rune{
	Orb:    'o',
	Scarab: ':',
	Cube:   '@',
	Device: '.',
}

var specialString = map[SpecialType]string{
	Orb:    "an orb of dragon slaying",
	Scarab: "a scarab of negate spirit",
	Cube:   "a cube of undead control",
	Device: "a device of theft prevention",
}

const (
	Orb SpecialType = iota
	Scarab
	Cube
	Device
)

// Special is a special item that don't offer stats but an in-game effect
type Special struct {
	Type SpecialType
	DefaultItem
}

// Log implements the Loggable interface
func (s *Special) Log() string {
	return s.String()
}

// Rune implements the io.Runeable interface
func (s *Special) Rune() rune {
	return specialRunes[s.Type]
}

// String implements the Item interface
func (s *Special) String() string {
	return specialString[s.Type]
}

// PickUp implements the Item interface
func (s *Special) PickUp(t *stats.Stats) {
	t.Special[int(s.Type)] = true
}

// Drop implements the Item interface
func (s *Special) Drop(t *stats.Stats) {
	t.Special[int(s.Type)] = false
}
