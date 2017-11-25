package game

import "fmt"

// world returns a string representation of the world
func (d *Data) world() [][]rune {
	grid := startZone()
	bar := d.infoBarGrid()
	for i := range bar { // Append the info bar
		grid = append(grid, bar[i])
	}
	return grid
}

// infoBarGrid returns the info bar in display grid format
func (d *Data) infoBarGrid() [][]rune {
	r := make([][]rune, 2)
	r[0] = []rune(fmt.Sprintf("Spells: %v( %v) AC: %v WC: %v Level %v Exp: %v %s", d.spells, d.maxSpells, d.ac, d.wc, d.level, d.exp, d.title))
	r[1] = []rune(fmt.Sprintf("HP: %v( %v) STR=%v INT=%v WIS=%v CON=%v DEX=%v CHA=%v LV: %v Gold: %v", d.hp, d.maxHP, d.str, d.intelligence, d.wisdom, d.con, d.dex, d.cha, d.loc, d.gold))

	return r
}

// startZone is the toplevel of the world
func startZone() [][]rune {
	grid := make([][]rune, borderHeight)
	for i := range grid {
		row := make([]rune, borderWidth)
		for j := range row {
			row[j] = '.'
		}
		grid[i] = row
	} // TODO THOR randomly insert start zone locations

	return grid
}
