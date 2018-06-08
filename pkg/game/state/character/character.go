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

const (
	// MaxLevel is the max level a character can acheive
	MaxLevel = 100
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
	TakeOffAction
)

const (
	characterFG   = termbox.ColorRed
	characterBG   = termbox.ColorRed
	characterRune = '&'
)

type Character struct {
	loc       types.Coordinate
	Stats     *stats.Stats
	inv       *Inventory
	Displaced io.Runeable // object character is currently on top of
}

// Init a character. Takes game difficulty which determines the characters starting items
func (c *Character) Init(d int) {
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
	c.inv = new(Inventory)
	c.inv.inv = make(map[rune]items.Item)

	if d <= 0 { // 0 difficulty games the plaer starts with leather armor and dagger
		w := items.GetNewWeapon(items.Dagger, 0)
		w.Attribute = 0
		if err := c.Wield(c.inv.AddItem(w, c.Stats)); err != nil {
			glog.Fatal("Uanble to wield starting weapon")
		}
		a := items.NewArmor(items.Leather, 0)
		a.Attribute = 0
		if err := c.Wear(c.inv.AddItem(a, c.Stats)); err != nil {
			glog.Fatal("Unable to wear starting armor")
		}
	}
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
func (c *Character) AddItem(i items.Item) rune {
	// TODO check if character can carry any more
	return c.inv.AddItem(i, c.Stats)
}

// DropItem removes an item from a characters inventory. Returns the item if there was no error
func (c *Character) DropItem(e rune) (items.Item, error) {
	return c.item(e, DropAction)
}

// Inventory returns a list of displayable inventory items
func (c *Character) Inventory() []string {
	return c.inv.List()
}

// TakeOff removes a characters armor
func (c *Character) TakeOff() error {
	_, err := c.item(none, TakeOffAction)
	return err
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

// item performs an item action on an item the character is carrying
func (c *Character) item(e rune, a action) (items.Item, error) {
	switch a {
	case ReadAction:
		return c.inv.Read(e, c.Stats)
	case DropAction:
		return c.inv.Drop(e, c.Stats)
	case WearAction:
		return c.inv.Wear(e, c.Stats)
	case WieldAction:
		return c.inv.Wield(e, c.Stats)
	case TakeOffAction:
		return c.inv.TakeOff(e, c.Stats)
	default:
		return nil, fmt.Errorf("Invalid action %v", a)
	}
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

// GainExperience has the character gain experience (ususally from slaying monsters)
func (c *Character) GainExperience(e int) bool {
	c.Stats.Exp += uint(e)
	levelGained := false
	for c.Stats.Exp >= uint(skill[c.Stats.Level]) && c.Stats.Level < MaxLevel {
		tmp := c.Stats.Con // TODO should take game difficulty into account
		c.Stats.Level++
		levelGained = true
		c.Stats.MaxHP += uint(rand.Intn(3) + 1 + rand.Intn(int(tmp)) + 1)
		c.Stats.MaxSpells += uint(rand.Intn(3))
		if c.Stats.Level < 7 { // - hardgame TODO
			c.Stats.MaxHP += c.Stats.Con >> 2
		}
	}

	// update player title
	c.Stats.Title = titles[c.Stats.Level-1]

	return levelGained
}
