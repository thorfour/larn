package monster

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/character"
	"github.com/thorfour/larn/pkg/io"
)

const (
	InvisibleRune = ' '
)

type Monster struct {
	Id         int         // the lookup id for the monster
	Hitpoints  int         // the remaining hitpoints for the monster
	Visibility bool        // if the player can see where this monster is
	Displaced  io.Runeable // the object currently displaced by this monster
}

// Level returns the monsters level
func (m *Monster) Level() int {
	return 1 // TODO keep an actual level for the monster
}

// Rune implements the io.Runeable interface
func (m *Monster) Rune() rune {
	if m.Visibility {
		return monsterData[m.Id].MonsterRune
	} else {
		return InvisibleRune
	}
}

// ID implements the monster interface
func (m *Monster) ID() int { return m.Id }

// Bg implements the io.Runeable interface
func (m *Monster) Bg() termbox.Attribute { return termbox.ColorDefault }

// Fg implements the io.Runeable interface
func (m *Monster) Fg() termbox.Attribute { return termbox.ColorDefault }

// Visible implements the Visibility interface
func (m *Monster) Visible(v bool) { m.Visibility = v }

// Damage implements the Monster interface
func (m *Monster) Damage(c *character.Character) {
	// TODO
}

// Name returns the name of the monster
func (m *Monster) Name() string { return NameFromID(m.Id) }

// New returns a new Monster from a monster id
func New(monster int) *Monster {
	return &Monster{
		Id:        monster,
		Hitpoints: monsterData[monster].Hitpoints,
		Displaced: Empty{},
	}
}

// Empty represents an empty map location
type Empty struct {
	visible bool
}

func (e Empty) Visible(v bool) { e.visible = v }

// Displace implementes the Displaceable interface
func (e Empty) Displace() bool { return true }

// Rune implements the io.Runeable interface
func (e Empty) Rune() rune {
	if e.visible {
		return '.'
	} else {
		return ' '
	}
}

// Fg implements the io.Runeable interface
func (e Empty) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (e Empty) Bg() termbox.Attribute { return termbox.ColorDefault }
