package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/thorfour/larn/pkg/game/state/items"
)

type forsale struct {
	index string // character index
	price int    // cost of this item
	stock int    // number of items in stock
	items.Item
}

var store = [][]forsale{
	{ // Page 1
		// Armor
		{"a", 20, 1, &items.ArmorClass{Type: items.Leather}},
		{"b", 100, 1, &items.ArmorClass{Type: items.StuddedLeather}},
		{"c", 400, 1, &items.ArmorClass{Type: items.RingMail}},
		{"d", 850, 1, &items.ArmorClass{Type: items.ChainMail}},
		{"e", 2200, 1, &items.ArmorClass{Type: items.SplintMail}},
		{"f", 4000, 1, &items.ArmorClass{Type: items.PlateMail}},
		{"g", 9000, 1, &items.ArmorClass{Type: items.PlateArmor}},
		{"h", 26000, 1, &items.ArmorClass{Type: items.StainlessPlateArmor}},
		// Weapons
		{"i", 1500, 1, &items.Shield{}},
		{"j", 20, 1, &items.WeaponClass{Type: items.Dagger}},
		{"k", 200, 1, &items.WeaponClass{Type: items.Spear}},
		{"l", 800, 1, &items.WeaponClass{Type: items.Flail}},
		{"m", 1500, 1, &items.WeaponClass{Type: items.BattleAxe}},
		{"n", 4500, 1, &items.WeaponClass{Type: items.LongSword}},
		{"o", 10000, 1, &items.WeaponClass{Type: items.TwoHandedSword}},
		{"p", 50000, 1, &items.WeaponClass{Type: items.SunSword}},
		{"q", 165000, 1, &items.WeaponClass{Type: items.LanceOfDeath}},
		// Rings
		{"t", 1500, 1, &items.Ring{Type: items.Protection}},
		{"u", 850, 1, &items.Ring{Type: items.Strength}},
		{"v", 1200, 1, &items.Ring{Type: items.Dexterity}},
		{"w", 1200, 1, &items.Ring{Type: items.Clever}},
		{"x", 1800, 1, &items.Ring{Type: items.Energy}},
		{"y", 1250, 1, &items.Ring{Type: items.Damage}},
		{"z", 2200, 1, &items.Ring{Type: items.Regen}},
	},
	{ // Page 2
		{"a", 10000, 1, &items.Ring{Type: items.ExtraRegen}},
		{"b", 2800, 1, &items.Belt{}},
		{"c", 4000, 1, &items.Ring{Type: items.Dexterity}}, // TODO amulet of invisibility
	},
}

// dndStoreSplash used to display the dnd store
func dndStoreSplash() string {
	return `
  Welcome to the Larn Thrift Shoppe.  We stock many items explorers find useful
  in their adventures.  Feel free to browse to your hearts content.
  Also be advised, if you break 'em, you pay for 'em.`
}

// dndstorepage renders a given page in the DND store
func dndstorepage(n int) string {
	pg := dndStoreSplash() + "\n"
	buf := bytes.NewBuffer(make([]byte, 100))
	w := tabwriter.NewWriter(buf, 5, 0, 1, ' ', tabwriter.TabIndent)
	for i, item := range store[n%len(store)] {
		switch i % 2 { // 2 items per line
		case 0:
			fmt.Fprintf(w, "  %s) %s\t\t%v\t", item.index, item.Item, item.price)
		case 1:
			fmt.Fprintf(w, "  %s) %s\t\t%v\t\n", item.index, item.Item, item.price)
		}
	}

	w.Flush()

	return pg + "\n" + string(buf.Bytes())
}
