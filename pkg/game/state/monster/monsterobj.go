package monster

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/character"
)

const (
	InvisibleRune = ' '
)

type MonsterObj struct {
	Id         int  // the lookup id for the monster
	Hitpoints  int  // the remaining hitpoints for the monster
	Visibility bool // if the player can see where this monster is
}

// Rune implements the io.Runeable interface
func (m *MonsterObj) Rune() rune {
	if m.Visibility {
		return monsterData[m.Id].MonsterRune
	} else {
		return InvisibleRune
	}
}

// Bg implements the io.Runeable interface
func (m *MonsterObj) Bg() termbox.Attribute { return termbox.ColorDefault }

// Fg implements the io.Runeable interface
func (m *MonsterObj) Fg() termbox.Attribute { return termbox.ColorDefault }

// Visible implements the Visibility interface
func (m *MonsterObj) Visible(v bool) { m.Visibility = v }

// MoveTowards implements the Monster interface
// makes no decisions about whether the monster should move, it only contains the logic
// of how to move it only performs the move
func (m *MonsterObj) MoveTowards(c *character.Character) {
	// TODO
}

// Damage implements the Monster interface
func (m *MonsterObj) Damage(c *character.Character) {
	// TODO
}

// New returns a new MonsterObj from a monster id
func New(monster int) *MonsterObj {
	return &MonsterObj{
		Id:        monster,
		Hitpoints: monsterData[monster].Hitpoints,
	}
}
