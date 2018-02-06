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

// MoveTowards implements the Monster interface
// makes no decisions about whether the monster should move, it only contains the logic
// of how to move it only performs the move
func (m *Monster) MoveTowards(c *character.Character) {
	// TODO
}

// Damage implements the Monster interface
func (m *Monster) Damage(c *character.Character) {
	// TODO
}

// New returns a new Monster from a monster id
func New(monster int) *Monster {
	return &Monster{
		Id:        monster,
		Hitpoints: monsterData[monster].Hitpoints,
	}
}
