package game

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"

	termbox "github.com/nsf/termbox-go"
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
		{"a", 20, 3, &items.ArmorClass{Type: items.Leather}},
		{"b", 100, 2, &items.ArmorClass{Type: items.StuddedLeather}},
		{"c", 400, 2, &items.ArmorClass{Type: items.RingMail}},
		{"d", 850, 2, &items.ArmorClass{Type: items.ChainMail}},
		{"e", 2200, 1, &items.ArmorClass{Type: items.SplintMail}},
		{"f", 4000, 1, &items.ArmorClass{Type: items.PlateMail}},
		{"g", 9000, 1, &items.ArmorClass{Type: items.PlateArmor}},
		{"h", 26000, 1, &items.ArmorClass{Type: items.StainlessPlateArmor}},
		// Weapons
		{"i", 1500, 1, &items.Shield{}},
		{"j", 20, 3, &items.WeaponClass{Type: items.Dagger}},
		{"k", 200, 3, &items.WeaponClass{Type: items.Spear}},
		{"l", 800, 2, &items.WeaponClass{Type: items.Flail}},
		{"m", 1500, 2, &items.WeaponClass{Type: items.BattleAxe}},
		{"n", 4500, 2, &items.WeaponClass{Type: items.LongSword}},
		{"o", 10000, 2, &items.WeaponClass{Type: items.TwoHandedSword}},
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
		{"j", 100, 3, &items.Cookie{}},
		// Potions
		{"k", 200, 6, &items.Potion{ID: items.Sleep, Store: true}},
		{"l", 900, 5, &items.Potion{ID: items.Healing, Store: true}},
		{"m", 5200, 1, &items.Potion{ID: items.RaiseLevel, Store: true}},
		{"n", 1000, 2, &items.Potion{ID: items.IncreaseAbility, Store: true}},
		{"o", 500, 2, &items.Potion{ID: items.GainWisdom, Store: true}},
		{"p", 1500, 2, &items.Potion{ID: items.GainStrength, Store: true}},
		{"q", 700, 1, &items.Potion{ID: items.IncreaseCharisma, Store: true}},
		{"r", 300, 7, &items.Potion{ID: items.Dizziness, Store: true}},
		{"s", 2000, 1, &items.Potion{ID: items.Learning, Store: true}},
		{"t", 500, 1, &items.Potion{ID: items.ObjectDetection, Store: true}},
		{"u", 800, 1, &items.Potion{ID: items.MonsterDetection, Store: true}},
		{"v", 300, 3, &items.Potion{ID: items.Forgetfulness, Store: true}},
		{"w", 200, 5, &items.Potion{ID: items.Water, Store: true}},
		{"x", 400, 3, &items.Potion{ID: items.Blindness, Store: true}},
		{"y", 350, 2, &items.Potion{ID: items.Confusion, Store: true}},
		{"z", 5200, 1, &items.Potion{ID: items.Heroism, Store: true}},
	},
	{ // Page 3
		// Potions
		{"a", 900, 2, &items.Potion{ID: items.Sturdiness, Store: true}},
		{"b", 2000, 2, &items.Potion{ID: items.GiantStrength, Store: true}},
		{"c", 2200, 4, &items.Potion{ID: items.FireResistance, Store: true}},
		{"d", 800, 6, &items.Potion{ID: items.TreasureFinding, Store: true}},
		{"e", 3700, 3, &items.Potion{ID: items.InstantHealing, Store: true}},
		{"f", 500, 1, &items.Potion{ID: items.Poison, Store: true}},
		{"g", 1500, 3, &items.Potion{ID: items.SeeInvisible, Store: true}},
		// Scrolls
		{"h", 1500, 2, &items.Scroll{ID: items.EnchantArmor, Store: true}},
		{"i", 1250, 2, &items.Scroll{ID: items.EnchantWeapon, Store: true}},
		{"j", 600, 4, &items.Scroll{ID: items.Englightenment, Store: true}},
		{"k", 100, 4, &items.Scroll{ID: items.Paper, Store: true}},
		{"l", 1000, 3, &items.Scroll{ID: items.CreateMonster, Store: true}},
		{"m", 2000, 2, &items.Scroll{ID: items.CreateItem, Store: true}},
		{"n", 1100, 1, &items.Scroll{ID: items.Aggravate, Store: true}},
		{"o", 5000, 2, &items.Scroll{ID: items.TimeWarp, Store: true}},
		{"p", 2000, 2, &items.Scroll{ID: items.Teleportation, Store: true}},
		{"q", 2500, 4, &items.Scroll{ID: items.ExpandedAwareness, Store: true}},
		{"r", 200, 5, &items.Scroll{ID: items.HasteMonster, Store: true}},
		{"s", 300, 3, &items.Scroll{ID: items.HealMonster, Store: true}},
		{"t", 3400, 1, &items.Scroll{ID: items.SpiritProtection, Store: true}},
		{"u", 3400, 1, &items.Scroll{ID: items.UndeadProtection, Store: true}},
		{"v", 3000, 2, &items.Scroll{ID: items.Stealth, Store: true}},
		{"w", 4000, 2, &items.Scroll{ID: items.MagicMapping, Store: true}},
		{"x", 5000, 2, &items.Scroll{ID: items.HoldMonster, Store: true}},
		{"y", 10000, 1, &items.Scroll{ID: items.GemPerfection, Store: true}},
		{"z", 5000, 1, &items.Scroll{ID: items.SpellExtension, Store: true}},
	},
	{ // Page 4
		{"a", 3400, 2, &items.Scroll{ID: items.Identify, Store: true}},
		{"b", 2200, 3, &items.Scroll{ID: items.RemoveCurse, Store: true}},
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
func dndstorepage(n int, gold uint) string {
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

	goldline := fmt.Sprintf("\n                                         You have %v gold pieces", gold)
	helpline := "\n\n  Enter your transaction [space for next page, escape to leave]"
	return pg + "\n" + string(buf.Bytes()) + goldline + helpline
}

// purchase an item from the store
func purchase(page int, k rune, gold *uint) (items.Item, error) {
	for i, v := range store[page%len(store)] {
		if v.index == string(k) { // Found the item to purchase
			// Check if item is in stock
			if store[page%len(store)][i].stock == 0 {
				return nil, fmt.Errorf("Sorry, but we are out of that item.")
			}

			if *gold < uint(v.price) { // unable to purchase the item
				return nil, fmt.Errorf("You don't have enough gold to pay for that!")
			}

			// Purchase the item
			store[page%len(store)][i].stock--
			item := store[page%len(store)][i].Item
			return item, nil
		}
	}

	return nil, fmt.Errorf("unable to purchase")
}

// dndStoreHandler inpout handler for the dnd store
func (g *Game) dndStoreHandler() func(termbox.Event) {
	page := 0
	g.renderSplash(dndstorepage(page, g.currentState.C.Stats.Gold))
	return func(e termbox.Event) {
		switch e.Key {
		case termbox.KeyEsc: // Exit
			g.inputHandler = g.defaultHandler
			g.render(display(g.currentState))
		case termbox.KeySpace: // Space key (next page)
			page++
			g.renderSplash(dndstorepage(page, g.currentState.C.Stats.Gold))
		default:
			switch e.Ch {
			case 'a':
				fallthrough
			case 'b':
				fallthrough
			case 'c':
				fallthrough
			case 'd':
				fallthrough
			case 'e':
				fallthrough
			case 'f':
				fallthrough
			case 'g':
				fallthrough
			case 'h':
				fallthrough
			case 'i':
				fallthrough
			case 'j':
				fallthrough
			case 'k':
				fallthrough
			case 'l':
				fallthrough
			case 'm':
				fallthrough
			case 'n':
				fallthrough
			case 'o':
				fallthrough
			case 'p':
				fallthrough
			case 'q':
				fallthrough
			case 'r':
				fallthrough
			case 's':
				fallthrough
			case 't':
				fallthrough
			case 'u':
				fallthrough
			case 'v':
				fallthrough
			case 'w':
				fallthrough
			case 'x':
				fallthrough
			case 'y':
				fallthrough
			case 'z':
				// Attempt to purchase an item
				item, err := purchase(page, e.Ch, &g.currentState.C.Stats.Gold)
				if err != nil {
					g.renderSplash(dndstorepage(page, g.currentState.C.Stats.Gold) + "\n\n  " + err.Error())
					time.Sleep(time.Millisecond * 700) // Quick blink the message
				} else {
					r := g.currentState.C.AddItem(item)
					g.renderSplash(dndstorepage(page, g.currentState.C.Stats.Gold) + "\n\n  " + fmt.Sprintf("You pick up: %s) %s", string(r), item))
					time.Sleep(time.Millisecond * 700) // Quick blink the message
				}
				g.renderSplash(dndstorepage(page, g.currentState.C.Stats.Gold))
			}
		}
	}
}
