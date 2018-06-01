package character

import (
	"fmt"
	"math/rand"

	"github.com/golang/glog"
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/game/state/stats"
	"github.com/thorfour/larn/pkg/game/state/types"
	"github.com/thorfour/larn/pkg/io"
)

var (
	NoSpellsErr     = fmt.Errorf("You don't have any spells!")
	NothingHappened = fmt.Errorf("  Nothing Happened")
	Inexperienced   = fmt.Errorf("  Nothing happens. You seem Inexperienced at this")
	DidntWork       = fmt.Errorf("  It didn't work!")
)

type action int

const (
	DropAction action = iota
	WieldAction
	WearAction
	ReadAction
)

const (
	characterFG   = termbox.ColorRed
	characterBG   = termbox.ColorRed
	characterRune = '&'
)

type Character struct {
	loc       types.Coordinate
	armor     []items.Armor  // Currently worn armor
	weapon    []items.Weapon // Currently wielded weapon(s)
	inventory []items.Item
	Stats     *stats.Stats
	Displaced io.Runeable // object character is currently on top of
}

func (c *Character) Init() {
	c.Stats = new(stats.Stats)
	c.Stats.Special = make(map[int]bool)
	c.Stats.KnownSpells = make(map[string]bool)
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
func (c *Character) MoveCharacter(d types.Direction) types.Coordinate {
	c.loc = types.Move(c.loc, d)
	return c.loc
}

func (c *Character) Location() types.Coordinate {
	return c.loc
}

// Teleport places a character at location l
func (c *Character) Teleport(x, y int) {
	c.loc.X = x
	c.loc.Y = y
}

// Wield has the character wield a weapon
func (c *Character) Wield(e rune) error {
	_, err := c.item(e, WieldAction)
	return err
}

// AddItem adds an item to the players inventory
func (c *Character) AddItem(i items.Item) {
	c.inventory = append(c.inventory, i)
}

// DropItem removes an item from a characters inventory. Returns the item if there was no error
// FIXME this isn't a stable removal. Items need to maintain their index
func (c *Character) DropItem(e rune) (items.Item, error) {

	i, err := c.item(e, DropAction)
	if err != nil {
		return nil, err
	}

	return i, nil
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

// TakeOff removes a characters armor
func (c *Character) TakeOff() error {
	if len(c.armor) == 0 {
		return fmt.Errorf("no armor being worn")
	}

	// Move all armor into inventory
	for i := range c.armor {
		c.inventory = append(c.inventory, c.armor[i])
		c.armor = append(c.armor[:i], c.armor[i+1:]...)
	}

	return nil
}

// Wear has the character wear a weapon
func (c *Character) Wear(e rune) error {
	_, err := c.item(e, WearAction)
	return err
}

func (c *Character) Read(e rune) ([]string, error) {
	i, err := c.item(e, ReadAction)
	if err != nil {
		return nil, err
	}
	return i.(items.Readable).Read(c.Stats), nil
}

func (c *Character) item(e rune, a action) (items.Item, error) {
	label := 'a'
	for i, w := range c.weapon {
		if label == e {
			switch a {
			case ReadAction:
				return nil, fmt.Errorf("You can't read that!")
			case DropAction:
				c.weapon = append(c.weapon[:i], c.weapon[i+1:]...)
				w.Disarm(c.Stats)
				return w, nil
			case WieldAction:
				return w, nil
			case WearAction:
				if t, ok := w.(items.Armor); ok { // Ensure the item is armor
					t.Wear(c.Stats)              // wear the armor
					c.armor = append(c.armor, t) // Add item to armor
					return t, nil
				} else {
					return nil, fmt.Errorf("You can't wear that!")
				}
			}
		}
		label++
	}

	for i, ar := range c.armor {
		if label == e {
			switch a {
			case ReadAction:
				return nil, fmt.Errorf("You can't read that!")
			case DropAction:
				c.armor = append(c.armor[:i], c.armor[i+1:]...)
				ar.TakeOff(c.Stats)
				return ar, nil
			case WieldAction:
				if t, ok := ar.(items.Weapon); ok { // Ensure the item is a weapon
					t.Wield(c.Stats)               // Wield the weapon
					c.weapon = append(c.weapon, t) // Add item to weapons
					return t, nil
				} else {
					return nil, fmt.Errorf("You can't wield item %s!", string(e))
				}
			case WearAction:
				return ar, nil
			}
		}
		label++
	}

	for i, t := range c.inventory {
		if label == e {
			switch a {
			case ReadAction:
				if _, ok := t.(items.Readable); ok { // Ensure item is readable
					c.inventory = append(c.inventory[:i], c.inventory[i+1:]...) // Remove the item from inventory
					return t, nil
				}
				return nil, fmt.Errorf("You can't read that!")
			case DropAction:
				c.inventory = append(c.inventory[:i], c.inventory[i+1:]...)
				t.Drop(c.Stats)
				return t, nil
			case WieldAction:
				if it, ok := t.(items.Weapon); ok { // Ensure the item is a weapon
					it.Wield(c.Stats)                                           // Wield the weapon
					c.weapon = append(c.weapon, it)                             // Add item to weapons
					c.inventory = append(c.inventory[:i], c.inventory[i+1:]...) // Remove item from inventory
					return it, nil
				} else {
					return nil, fmt.Errorf("You can't wield item %s!", string(e))
				}
			case WearAction:
				if it, ok := t.(items.Armor); ok { // Ensure the item is armor
					it.Wear(c.Stats)                                            // wear the armor
					c.armor = append(c.armor, it)                               // Add item to armor
					c.inventory = append(c.inventory[:i], c.inventory[i+1:]...) // Remove item from inventory
					return it, nil
				} else {
					return nil, fmt.Errorf("You can't wear that!")
				}
			}
		}
		label++
	}

	return nil, fmt.Errorf("You don't have item %s!", string(e))
}

// Cast handles the bookkeeping for a character casting a spell
func (c *Character) Cast(s string) (*items.Spell, error) {
	if c.Stats.Spells == 0 { // this should never happen, there's a guard before calls to this
		glog.Error("Cast requested with no spells")
		return nil, NoSpellsErr
	}

	// lookup spell and remove available spells from caster
	spell := items.SpellFromCode(s)
	c.Stats.Spells--

	// check if caster knows this spell
	if !c.Stats.KnownSpells[s] {
		return nil, NothingHappened
	}

	// check if caster has enough intelligence also always random chance to fail
	if rand.Intn(23) == 0 || rand.Intn(18) > int(c.Stats.Intelligence) {
		return nil, DidntWork
	}

	// check if caster is high level enough to cast spell
	if int(c.Stats.Level)*3+2 < spell.Level {
		return nil, Inexperienced
	}

	// Return the spell the character cast
	return &spell, nil
}

//Heal the character up to their max hp
func (c *Character) Heal(hp int) {
	c.Stats.Hp += uint(hp)
	if c.Stats.Hp > c.Stats.MaxHP {
		c.Stats.Hp = c.Stats.MaxHP
	}
}

// Damage decreases the HP of character
func (c *Character) Damage(dmg int) bool {
	if dmg <= 0 {
		return false
	}
	if uint(dmg) > c.Stats.Hp {
		c.Stats.Hp = 0
		return true
	}
	c.Stats.Hp -= uint(dmg)
	return false
}
