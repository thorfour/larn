package character

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/game/state/stats"
)

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
	UpLeft
	UpRight
	DownLeft
	DownRight
	Here
)

const (
	characterFG   = termbox.ColorGreen
	characterBG   = termbox.ColorGreen
	characterRune = '&'
)

type Character struct {
	loc       Coordinate
	armor     []items.Armor  // Currently worn armor
	weapon    []items.Weapon // Currently wielded weapon(s)
	inventory []items.Item
	//knownSpells []Spells
	stats.Stats
}

type Coordinate struct {
	X int
	Y int
}

func (c *Character) Rune() rune {
	return characterRune
}

func (c *Character) Fg() termbox.Attribute {
	return characterFG
}

func (c *Character) Bg() termbox.Attribute {
	return characterBG
}

// Move the character in the given direction 1 space
func (c *Character) Move(d Direction) Coordinate {
	switch d {
	case Up:
		c.loc.Y--
	case Down:
		c.loc.Y++
	case Left:
		c.loc.X--
	case Right:
		c.loc.X++
	case UpLeft:
		c.loc.Y--
		c.loc.X--
	case UpRight:
		c.loc.Y--
		c.loc.X++
	case DownLeft:
		c.loc.Y++
		c.loc.X--
	case DownRight:
		c.loc.Y++
		c.loc.X++
	case Here:
	}

	return c.loc
}

func (c *Character) Location() Coordinate {
	return c.loc
}

func (c *Character) Teleport() {}
