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
		// Filler
		{"r", -1, 0, nil},
		{"s", -1, 0, nil},
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
		// Filler
		{"d", -1, 0, nil},
		{"e", -1, 0, nil},
		{"f", -1, 0, nil},
		{"g", -1, 0, nil},
		// Items
		{"h", 5900, 1, &items.Chest{}},
		{"i", 2000, 1, &items.Book{}},
		{"j", 100, 1, &items.Cookie{}},
		// Potions
		{"k", 200, 1, &items.Potion{ID: items.Sleep, Store: true}},
		{"l", 900, 1, &items.Potion{ID: items.Healing, Store: true}},
		{"m", 5200, 1, &items.Potion{ID: items.RaiseLevel, Store: true}},
		{"n", 1000, 1, &items.Potion{ID: items.IncreaseAbility, Store: true}},
		{"o", 500, 1, &items.Potion{ID: items.GainWisdom, Store: true}},
		{"p", 1500, 1, &items.Potion{ID: items.GainStrength, Store: true}},
		{"q", 700, 1, &items.Potion{ID: items.IncreaseCharisma, Store: true}},
		{"r", 300, 1, &items.Potion{ID: items.Dizziness, Store: true}},
		{"s", 2000, 1, &items.Potion{ID: items.Learning, Store: true}},
		{"t", 500, 1, &items.Potion{ID: items.ObjectDetection, Store: true}},
		{"u", 800, 1, &items.Potion{ID: items.MonsterDetection, Store: true}},
		{"v", 300, 1, &items.Potion{ID: items.Forgetfulness, Store: true}},
		{"w", 200, 1, &items.Potion{ID: items.Water, Store: true}},
		{"x", 400, 1, &items.Potion{ID: items.Blindness, Store: true}},
		{"y", 350, 1, &items.Potion{ID: items.Confusion, Store: true}},
		{"z", 5200, 1, &items.Potion{ID: items.Heroism, Store: true}},
	},
	{ // Page 3
		// Potions
		{"a", 900, 1, &items.Potion{ID: items.Sturdiness, Store: true}},
		{"b", 2000, 1, &items.Potion{ID: items.GiantStrength, Store: true}},
		{"c", 2200, 1, &items.Potion{ID: items.FireResistance, Store: true}},
		{"d", 800, 1, &items.Potion{ID: items.TreasureFinding, Store: true}},
		{"e", 3700, 1, &items.Potion{ID: items.InstantHealing, Store: true}},
		{"f", 500, 1, &items.Potion{ID: items.Poison, Store: true}},
		{"g", 1500, 1, &items.Potion{ID: items.SeeInvisible, Store: true}},
		// Scrolls
		{"h", 1500, 1, &items.Scroll{ID: items.EnchantArmor, Store: true}},
		{"i", 1250, 1, &items.Scroll{ID: items.EnchantWeapon, Store: true}},
		{"j", 600, 1, &items.Scroll{ID: items.Englightenment, Store: true}},
		{"k", 100, 1, &items.Scroll{ID: items.Paper, Store: true}},
		{"l", 1000, 1, &items.Scroll{ID: items.CreateMonster, Store: true}},
		{"m", 2000, 1, &items.Scroll{ID: items.CreateItem, Store: true}},
		{"n", 1100, 1, &items.Scroll{ID: items.Aggravate, Store: true}},
		{"o", 5000, 1, &items.Scroll{ID: items.TimeWarp, Store: true}},
		{"p", 2000, 1, &items.Scroll{ID: items.Teleportation, Store: true}},
		{"q", 2500, 1, &items.Scroll{ID: items.ExpandedAwareness, Store: true}},
		{"r", 200, 1, &items.Scroll{ID: items.HasteMonster, Store: true}},
		{"s", 300, 1, &items.Scroll{ID: items.HealMonster, Store: true}},
		{"t", 3400, 1, &items.Scroll{ID: items.SpiritProtection, Store: true}},
		{"u", 3400, 1, &items.Scroll{ID: items.UndeadProtection, Store: true}},
		{"v", 3000, 1, &items.Scroll{ID: items.Stealth, Store: true}},
		{"w", 4000, 1, &items.Scroll{ID: items.MagicMapping, Store: true}},
		{"x", 5000, 1, &items.Scroll{ID: items.HoldMonster, Store: true}},
		{"y", 10000, 1, &items.Scroll{ID: items.GemPerfection, Store: true}},
		{"z", 5000, 1, &items.Scroll{ID: items.SpellExtension, Store: true}},
	},
	{ // Page 4
		{"a", 3400, 1, &items.Scroll{ID: items.Identify, Store: true}},
		{"b", 2200, 1, &items.Scroll{ID: items.RemoveCurse, Store: true}},
		{"c", -1, 0, nil},
		{"d", 6100, 1, &items.Scroll{ID: items.Pulverization, Store: true}},
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
			if item.stock == 0 {
				fmt.Fprint(w, "      \t\t \t")
			} else {
				fmt.Fprintf(w, "  %s) %s\t\t%v\t", item.index, item.Item, item.price)
			}
		case 1:
			if item.stock == 0 {
				fmt.Fprint(w, "      \t\t \t\n")
			} else {
				fmt.Fprintf(w, "  %s) %s\t\t%v\t\n", item.index, item.Item, item.price)
			}
		}
	}

	w.Flush()

	return pg + "\n" + string(buf.Bytes())
}
