package items

import "github.com/thorfour/larn/pkg/game/state/stats"

const (
	invisibleRune = ' '
)

// Item is the generis *item interface
type Item interface {
	PickUp(s *stats.Stats)
	Drop(s *stats.Stats)
}

// Food for edible Items (fortune s *ookies)
type Food interface {
	Item
	Eat(s *stats.Stats)
}

// Potion for anything that s *an be quaffed
type Potion interface {
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

// DisplaceableItem is used as an embedded struct to indicate an object is displaceable
type DisplaceableItem struct{}

// Displace implements the displaceable interface
func (d *DisplaceableItem) Displace() bool { return true }
