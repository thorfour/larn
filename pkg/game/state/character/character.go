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
)

const (
	characterFG = termbox.ColorGreen
	characterBG = termbox.ColorGreen
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
	x int
	y int
}

// Move the character in the given direction 1 space
func (c *Character) Move(d Direction) {}
func (c *Character) Teleport()        {}
