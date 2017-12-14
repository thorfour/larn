package character

import (
	"fmt"

	"github.com/golang/glog"
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
	c.Stats.Special = make(map[int]bool)
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

// Wield has the character wield a weapon
func (c *Character) Wield(e rune) error {
	label := 'a'
	for i, item := range c.inventory {
		if label == e {
			if t, ok := item.(items.Weapon); ok { // Ensure the item is a weapon
				t.Wield(c.Stats)                                            // Wield the weapon
				c.weapon = append(c.weapon, t)                              // Add item to weapons
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...) // Remove item from inventory
				return nil
			} else {
				return fmt.Errorf("You can't wield item %s!", string(e))
			}
		}
		label++
	}
	return fmt.Errorf("You don't have item %s!", string(e))
}

// AddItem adds an item to the players inventory
func (c *Character) AddItem(i items.Item) {
	c.inventory = append(c.inventory, i)
}

// DropItem removes an item from a characters inventory. Returns the item if there was no error
// FIXME this isn't a stable removal. Items need to maintain their index
func (c *Character) DropItem(e rune) (items.Item, error) {
	label := 'a'
	for i, w := range c.weapon {
		if label == e {
			c.weapon = append(c.weapon[:i], c.weapon[i+1:]...)
			return w, nil
		}
		label++
	}

	for i, a := range c.armor {
		if label == e {
			c.armor = append(c.armor[:i], c.armor[i+1:]...)
			return a, nil
		}
		label++
	}

	for i, t := range c.inventory {
		if label == e {
			c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
			return t, nil
		}
		label++
	}

	return nil, fmt.Errorf("You don't have item %s!", string(e))
}

// Inventory returns a list of displayable inventory items
func (c *Character) Inventory() []string {
	var inv []string
	for _, i := range c.armor {
		inv = append(inv, fmt.Sprintf("%v %v", i.String(), "(being worn)"))
	}
	for _, i := range c.weapon {
		inv = append(inv, fmt.Sprintf("%v %v", i.String(), "(weapon in hand)"))
	}
	for _, i := range c.inventory {
		inv = append(inv, i.String())
	}

	glog.V(4).Infof("Inventory: %v", inv)
	return inv
}
