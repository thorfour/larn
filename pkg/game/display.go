package game

import (
	"fmt"

	"github.com/thorfour/larn/pkg/game/state"
)

// world returns a 2d slice representation of the world
func world(s *state.State) [][]rune {
	// TODO this function needs to handle generating world from saved games

	// Generate start zone
	grid := startZone()

	// Generate info bar
	bar := infoBarGrid(s)
	for i := range bar { // Append the info bar
		grid = append(grid, bar[i])
	}

	return grid
}

// infoBarGrid returns the info bar in display grid format
func infoBarGrid(s *state.State) [][]rune {
	r := make([][]rune, 2)
	r[0] = []rune(fmt.Sprintf("Spells: %v( %v) AC: %v WC: %v Level %v Exp: %v %s", s.Spells, s.MaxSpells, s.Ac, s.Wc, s.Level, s.Exp, s.Title))
	r[1] = []rune(fmt.Sprintf("HP: %v( %v) STR=%v INT=%v WIS=%v CON=%v DEX=%v CHA=%v LV: %v Gold: %v", s.Hp, s.MaxHP, s.Str, s.Intelligence, s.Wisdom, s.Con, s.Dex, s.Cha, s.Loc, s.Gold))

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
