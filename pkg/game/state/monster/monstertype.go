package monster

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/character"
)

type MonsterType struct {
	MonsterRune  rune   // the monsters displayable rune
	Name         string // the monsters displayable name
	Id           int
	Lvl          int
	Armor        int
	Dmg          int
	Attack       int
	Defense      int
	Intelligence int
	Gold         int
	Hitpoints    int
	Experience   int
	Visibility   bool
}

// Rune implements the io.Runeable interface
func (m *MonsterType) Rune() rune {
	// TODO
	return 'B'
}

// Bg implements the io.Runeable interface
func (m *MonsterType) Bg() termbox.Attribute { return termbox.ColorDefault }

// Fg implements the io.Runeable interface
func (m *MonsterType) Fg() termbox.Attribute { return termbox.ColorDefault }

// Displace implements the Displaceable interface
func (m *MonsterType) Displace() bool { return false } // monsters are never displaceable

// Visible implements the Visibility interface
func (m *MonsterType) Visible(v bool) { m.Visibility = v }

// MoveTowards implements the Monster interface
func (m *MonsterType) MoveTowards(c *character.Character) {
	// TODO
}

// Damage implements the Monster interface
func (m *MonsterType) Damage(c *character.Character) {
	// TODO
}

// New returns a new monstertype from a monster id
func New(monster int) *MonsterType {
	m := new(MonsterType)
	// TODO
	return m
}
