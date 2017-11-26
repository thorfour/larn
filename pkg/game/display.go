package game

import (
	"fmt"

	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state"
	"github.com/thorfour/larn/pkg/io"
)

type Simple rune

func (s Simple) Rune() rune            { return rune(s) }
func (s Simple) Fg() termbox.Attribute { return termbox.ColorDefault }
func (s Simple) Bg() termbox.Attribute { return termbox.ColorDefault }

// display returns a 2d slice representation of the game
func display(s *state.State) [][]io.Runeable {
	// Generate start zone
	grid := s.CurrentMap()

	// Generate info bar
	bar := infoBarGrid(s)
	for i := range bar { // Append the info bar
		grid = append(grid, bar[i])
	}

	return grid
}

// infoBarGrid returns the info bar in display grid format
func infoBarGrid(s *state.State) [][]io.Runeable {
	r := make([][]io.Runeable, 2)

	info := fmt.Sprintf("Spells: %v( %v) AC: %v WC: %v Level %v Exp: %v %s", s.Spells, s.MaxSpells, s.Ac, s.Wc, s.Level, s.Exp, s.Title)
	for _, c := range info {
		r[0] = append(r[0], Simple(c))
	}

	info = fmt.Sprintf("HP: %v( %v) STR=%v INT=%v WIS=%v CON=%v DEX=%v CHA=%v LV: %v Gold: %v", s.Hp, s.MaxHP, s.Str, s.Intelligence, s.Wisdom, s.Con, s.Dex, s.Cha, s.Loc, s.Gold)
	for _, c := range info {
		r[1] = append(r[1], Simple(c))
	}

	return r
}
