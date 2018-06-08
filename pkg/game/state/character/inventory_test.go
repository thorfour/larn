package character

import (
	"strings"
	"testing"

	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/game/state/stats"
)

func TestAddItems(t *testing.T) {
	i := NewInventory()
	s := new(stats.Stats)

	w := items.GetNewWeapon(items.Dagger, 0)
	if r := i.AddItem(w, s); r != 'a' {
		t.Error("unexpected rune returned from first item")
	}

	if r := i.AddItem(w, s); r != 'b' {
		t.Error("unexpected rune returned from second item")
	}

	if r := i.AddItem(w, s); r != 'c' {
		t.Error("unexpected rune returned from third item")
	}
}

func TestDropItems(t *testing.T) {
	i := NewInventory()
	s := new(stats.Stats)

	i.AddItem(items.GetNewWeapon(items.Dagger, 0), s)
	i.AddItem(items.GetNewWeapon(items.Dagger, 0), s)
	i.Drop('a', s)

	// Ensure the remaining item is item B
	r := i.List()
	if len(r) != 1 {
		t.Error("unexpected number of results")
	}

	if !strings.HasPrefix(r[0], "b)") {
		t.Error("unexpected index of item")
	}
}

func TestItemOrder(t *testing.T) {
	i := NewInventory()
	s := new(stats.Stats)

	i.AddItem(items.GetNewWeapon(items.Dagger, 0), s)
	i.AddItem(items.GetNewWeapon(items.Dagger, 0), s)
	i.AddItem(items.GetNewWeapon(items.Dagger, 0), s)
	i.AddItem(items.GetNewWeapon(items.Dagger, 0), s)
	i.Drop('c', s)

	// Ensure the remaining item is item B
	r := i.List()
	for i := range r {
		switch getIndex(r[i]) {
		case "a)":
		case "b)":
		case "d)":
		default:
			t.Error("unexpected index returned")
		}
	}
}

func getIndex(r string) string {
	return strings.Split(r, " ")[0]
}

// TestInit is a unit test to ensure the character Init function wont ever fail
func TestInit(t *testing.T) {
	i := NewInventory()
	s := new(stats.Stats)

	if _, err := i.Wield(i.AddItem(items.GetNewWeapon(items.Dagger, 0), s), s); err != nil {
		t.Error("failed to wield", err)
	}
	if _, err := i.Wear(i.AddItem(items.NewArmor(items.Leather, 0), s), s); err != nil {
		t.Error("failed to wear", err)
	}
}
