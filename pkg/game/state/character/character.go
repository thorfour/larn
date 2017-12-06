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
	characterFG   = termbox.ColorRed
	characterBG   = termbox.ColorRed
	characterRune = '&'
)

type Character struct {
	loc       Coordinate
	armor     []items.Armor  // Currently worn armor
	weapon    []items.Weapon // Currently wielded weapon(s)
	inventory []items.Item
	//knownSpells []Spells
	Stats *stats.Stats
}

type Coordinate struct {
	X int
	Y int
}

func (c *Character) Init() {
	c.Stats = new(stats.Stats)
	c.Stats.Level = 1
	c.Stats.Title = titles[c.Stats.Level-1]
	c.Stats.MaxSpells = 1
	c.Stats.Spells = 1
	c.Stats.MaxHP = 5
	c.Stats.Hp = 5
	c.Stats.Cha = 12
	c.Stats.Str = 12
	c.Stats.Intelligence = 12
	c.Stats.Wisdom = 12
	c.Stats.Con = 12
	c.Stats.Dex = 12
	c.Stats.Cha = 12
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

// MoveCharacter the character in the given direction 1 space
func (c *Character) MoveCharacter(d Direction) Coordinate {
	c.loc.Move(d)
	return c.loc
}

func (c *Character) Location() Coordinate {
	return c.loc
}

// Teleport places a character at location l
func (c *Character) Teleport(x, y int) {
	c.loc.X = x
	c.loc.Y = y
}

func (c *Coordinate) Move(d Direction) {
	switch d {
	case Up:
		c.Y--
	case Down:
		c.Y++
	case Left:
		c.X--
	case Right:
		c.X++
	case UpLeft:
		c.Y--
		c.X--
	case UpRight:
		c.Y--
		c.X++
	case DownLeft:
		c.Y++
		c.X--
	case DownRight:
		c.Y++
		c.X++
	case Here:
	}
}

// AddItem adds an item to the players inventory
func (c *Character) AddItem(i items.Item) {
	c.inventory = append(c.inventory, i)
}
