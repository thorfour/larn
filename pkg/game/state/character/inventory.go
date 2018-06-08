package character

import (
	"fmt"
	"sort"
	"strings"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/game/state/stats"
)

const none = '0'

// Inventory represents the characters inventory
type Inventory struct {
	shield rune
	weapon rune
	armor  rune
	inv    map[rune]items.Item
	unused []rune
}

// NewInventory returns a new initialized inventory struct
func NewInventory() *Inventory {
	i := new(Inventory)
	i.inv = make(map[rune]items.Item)
	i.shield = none
	i.weapon = none
	i.armor = none

	return i
}

// List returns all the items in the inventory
func (i *Inventory) List() []string {
	var ret []string
	for r := range i.inv {
		switch r {
		case i.weapon:
			ret = append(ret, fmt.Sprintf("%s) %s %s", string(r), i.inv[r], "(weapon in hand)"))
		case i.shield:
			fallthrough
		case i.armor:
			ret = append(ret, fmt.Sprintf("%s) %s %s", string(r), i.inv[r], "(being worn)"))
		default:
			ret = append(ret, fmt.Sprintf("%s) %s", string(r), i.inv[r].String()))
		}
	}

	// NOTE: since maps return random order, sort the return slice here
	sort.Slice(ret, func(i, j int) bool {
		return strings.Split(ret[i], " ")[0] < strings.Split(ret[j], " ")[0]
	})

	glog.V(4).Infof("Inventory: %v", ret)

	return ret
}

// AddItem adds a new item the the inventory and returns its assigned rune
func (i *Inventory) AddItem(item items.Item, s *stats.Stats) rune {
	slot := 'a'
	if len(i.unused) == 0 {
		slot += rune(len(i.inv))
	} else {
		slot = i.unused[0]
		i.unused = i.unused[1:]
	}

	i.inv[slot] = item
	item.PickUp(s)

	return slot
}

// Drop an item. Returns the item that was dropped. Caller should call necessary drop func
func (i *Inventory) Drop(r rune, s *stats.Stats) (items.Item, error) {
	item, ok := i.inv[r]
	if !ok {
		return nil, fmt.Errorf("You don't have item %s!", string(r))
	}

	// Remove from wear/wield
	switch r {
	case i.shield:
		fallthrough
	case i.weapon:
		i.Disarm(r, s)
	case i.armor:
		i.TakeOff(r, s)
	}

	delete(i.inv, r)
	i.unused = append(i.unused, r)
	sort.Slice(i.unused, func(a, b int) bool {
		return i.unused[a] < i.unused[b]
	})
	item.Drop(s)

	return item, nil
}

// Wield a weapon. Replaces currently wielded weapon. Caller should call Wield() on returned weapon
func (i *Inventory) Wield(r rune, s *stats.Stats) (items.Weapon, error) {
	item, ok := i.inv[r]
	if !ok {
		return nil, fmt.Errorf("You don't have item %s", string(r))
	}

	w, ok := item.(items.Weapon)
	if !ok {
		return nil, fmt.Errorf("You can't wield item %s", string(r))
	}

	// TODO check if two handed sword and shield

	// Mark this weapn as being wielded
	i.weapon = r
	w.Wield(s)

	return w, nil
}

// Wear to put on armor
func (i *Inventory) Wear(r rune, s *stats.Stats) (items.Armor, error) {
	item, ok := i.inv[r]
	if !ok {
		return nil, fmt.Errorf("You don't have item %s", string(r))
	}

	if i.armor != none {
		return nil, fmt.Errorf("You're already wearing armor")
	}

	a, ok := item.(items.Armor)
	if !ok {
		return nil, fmt.Errorf("You can't wear that!")
	}

	i.armor = r
	a.Wear(s)

	return a, nil
}

// TakeOff armor
func (i *Inventory) TakeOff(_ rune, s *stats.Stats) (items.Armor, error) {
	if i.armor == none {
		return nil, fmt.Errorf("You're not wearing anything")
	}

	a, ok := i.inv[i.armor].(items.Armor)
	if !ok {
		glog.Errorf("not wearing armor: %s", a)
	}
	i.armor = none
	a.TakeOff(s)

	return a, nil
}

// Read a book or scroll. Caller should call Read() on returned item
func (i *Inventory) Read(r rune, s *stats.Stats) (items.Item, error) {
	item, ok := i.inv[r]
	if !ok {
		return nil, fmt.Errorf("You don't have item %s", string(r))
	}

	if _, ok := item.(items.Readable); !ok {
		return nil, fmt.Errorf("You can't read that!")
	}

	delete(i.inv, r)
	return item, nil
}

// Disarm wielded weapon
func (i *Inventory) Disarm(_ rune, s *stats.Stats) (items.Weapon, error) {
	if i.weapon == none {
		return nil, fmt.Errorf("You're not wielding anything")
	}

	w, ok := i.inv[i.weapon].(items.Weapon)
	if !ok {
		glog.Errorf("not wielding weapon: %s", w)
	}
	i.weapon = none
	w.Disarm(s)

	return w, nil
}
