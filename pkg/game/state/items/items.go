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
			created = append(created, &GoldPile{Amount: rand.Intn((l+1)*10) + l*10 + 11})
		case tmp < 16: // book
			created = append(created, &Book{Level: uint(l)})
		case tmp < 19: // dagger
			created = append(created, GetNewWeapon(Dagger, l))
		case tmp < 22: // leather armor
			created = append(created, NewArmor(Leather, l))
		case tmp < 23: // regen ring
			created = append(created, &Ring{Type: Regen, Attribute: rand.Intn(l/3 + 1)})
		case tmp < 24: // shield
			created = append(created, &Shield{Attribute: rand.Intn(l/3 + 1)})
		case tmp < 25: // 2 hand sword
			created = append(created, GetNewWeapon(TwoHandedSword, l))
		case tmp < 26: // prot ring
			created = append(created, &Ring{Type: Protection, Attribute: rand.Intn(l/4 + 1)})
		case tmp < 27: // dex ring
			created = append(created, &Ring{Type: Dexterity, Attribute: rand.Intn(l/4 + 1)})
		case tmp < 28: // energy ring
			created = append(created, &Ring{Type: Energy, Attribute: rand.Intn(l/4 + 1)})
		case tmp < 29: // str ring
			created = append(created, &Ring{Type: Strength, Attribute: rand.Intn(l/2 + 1)})
		case tmp < 30: // cleverness ring
			created = append(created, &Ring{Type: Clever, Attribute: rand.Intn(l/2 + 1)})
		case tmp < 31: // ring mail
			created = append(created, NewArmor(RingMail, l))
		case tmp < 32: // flail
			created = append(created, GetNewWeapon(Flail, l))
		case tmp < 33: // spear
			created = append(created, GetNewWeapon(Spear, l))
		case tmp < 34: // battleaxe
			created = append(created, GetNewWeapon(BattleAxe, l))
		case tmp < 35: // belt
			created = append(created, &Belt{Attribute: rand.Intn(l/2 + 1)})
		case tmp < 36: // studded leather
			created = append(created, NewArmor(StuddedLeather, l))
		case tmp < 37: // splint
			created = append(created, NewArmor(SplintMail, l))
		case tmp < 38: // fortune cookie
			created = append(created, &Cookie{})
		case tmp < 39: // chain mail
			created = append(created, NewArmor(ChainMail, l))
		case tmp < 40: // plate mail
			created = append(created, NewArmor(PlateMail, l))
		case tmp < 41: // longsword
			created = append(created, GetNewWeapon(LongSword, l))

		}
	}
	return created
}
