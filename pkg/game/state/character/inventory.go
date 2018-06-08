package character

import (
	"fmt"

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

// List returns all the items in the inventory
func (i *Inventory) List() []string {
	r := 'a'
	var ret []string
	for j := 0; j < len(i.inv); j++ {
		switch r {
		case i.weapon:
			ret = append(ret, fmt.Sprintf("%s %s", i.inv[r], "(weapon in hand)"))
		case i.shield:
			fallthrough
		case i.armor:
			ret = append(ret, fmt.Sprintf("%s %s", i.inv[r], "(being worn)"))
		default:
			ret = append(ret, i.inv[r].String())
		}

		r++
	}

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
		return nil, fmt.Errorf("You don't have item %v!", r)
	}

	delete(i.inv, r)
	i.unused = append(i.unused, r)
	// TODO sort the list everytime a rune is added
	item.Drop(s)

	return item, nil
}

// Wield a weapon. Replaces currently wielded weapon. Caller should call Wield() on returned weapon
func (i *Inventory) Wield(r rune, s *stats.Stats) (items.Weapon, error) {
	item, ok := i.inv[r]
	if !ok {
		return nil, fmt.Errorf("You don't have item %v", r)
	}

	w, ok := item.(items.Weapon)
	if !ok {
		return nil, fmt.Errorf("You can't wield item %v", r)
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
		return nil, fmt.Errorf("You don't have item %v", r)
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
		return nil, fmt.Errorf("You don't have item %v", r)
	}

	if _, ok := item.(items.Readable); !ok {
		return nil, fmt.Errorf("You can't read that!")
	}

	delete(i.inv, r)
	return item, nil
}
