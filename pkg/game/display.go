package game

import "fmt"

// world returns a string representation of the world
func (d *Data) world() string {
	return d.infoBar()
}

// infoBar returns the stringified info bar for terminal display
func (d *Data) infoBar() string {
	return fmt.Sprintf("Spells: %v( %v) AC: %v WC: %v Level %v Exp: %v %s\nHP: %v( %v) STR=%v INT=%v WIS=%v CON=%v DEX=%v CHA=%v LV: %v Gold: %v",
		d.spells, d.maxSpells, d.ac, d.wc, d.level, d.exp, d.title, d.hp, d.maxHP, d.str, d.intelligence, d.wisdom, d.con, d.dex, d.cha, d.loc, d.gold)
}
