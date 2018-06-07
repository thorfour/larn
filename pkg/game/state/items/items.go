package items

import (
	"math/rand"

	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state/stats"
	"github.com/thorfour/larn/pkg/io"
)

const (
	invisibleRune = ' '
)

// Item is the generis *item interface
type Item interface {
	PickUp(s *stats.Stats)
	Drop(s *stats.Stats)
	String() string
	io.Runeable
}

// Food for edible Items (fortune s *ookies)
type Food interface {
	Item
	Eat(s *stats.Stats)
}

// Quaffable for anything that s *an be quaffed
type Quaffable interface {
	Item
	Quaff(s *stats.Stats)
}

// Weapon for anything that s *an be wielded
type Weapon interface {
	Item
	Wield(s *stats.Stats)
	Disarm(s *stats.Stats)
}

// Armor interface for anything that s *an be used as armor
type Armor interface {
	Item
	Wear(s *stats.Stats)
	TakeOff(s *stats.Stats)
}

// Readable interface for any items that the user can read
type Readable interface {
	Read(s *stats.Stats) []string
}

// DefaultItem provide default Fg and Bg functions
type DefaultItem struct {
	Visibility bool
}

// Fg for implementing the io.Runeable interface
func (d *DefaultItem) Fg() termbox.Attribute { return termbox.ColorDefault | termbox.AttrBold }

// Bg for implementing the io.Runeable interface
func (d *DefaultItem) Bg() termbox.Attribute { return termbox.ColorDefault | termbox.AttrBold }

// Visible implements the visibility interface
func (d *DefaultItem) Visible(v bool) { d.Visibility = v }

// Displace implements the displaceable interface
func (d *DefaultItem) Displace() bool { return true }

// NoStats provides empty PickUp and Drop functions
type NoStats struct{}

// PickUp implements the Item interface
func (n *NoStats) PickUp(_ *stats.Stats) {}

// Drop implements the Item interface
func (n *NoStats) Drop(_ *stats.Stats) {}

// CreateItems creates a random item based on the given level
func CreateItems(l int) []Item {
	itemCount := 1
	for i := rand.Intn(101); i < 8; i = rand.Intn(101) { // Chance to create multiple items
		itemCount++
	}

	var created []Item
	for i := 0; i < itemCount; i++ {
		// TODO create the item
		tmp := 33
		if l > 6 {
			tmp = 41
		} else if l > 4 {
			tmp = 39
		}
		tmp = rand.Intn(tmp)
		switch {
		case tmp < 4: // scroll
			created = append(created, NewScroll())
		case tmp < 8: // potion
			created = append(created, NewPotion())
		case tmp < 12: // gold
			created = append(created, &GoldPile{})
		case tmp < 16: // book
		case tmp < 19: // dagger
		case tmp < 22: // leather armor
		case tmp < 25: // regen ring, shield, 2 hand sword
		case tmp < 27: // prot ring, dex ring
		case tmp < 28: // energy ring
		case tmp < 30: // str ring, cleverness ring
		case tmp < 32: // ring mail, flail
		case tmp < 34: // spear, battleaxe
		case tmp < 37: // belt, studded leather, splint
		case tmp < 38: // fortune cookie
		case tmp < 39: // chain mail
		case tmp < 40: // plate mail
		case tmp < 41: // longsword

		}
	}
	return created
}
