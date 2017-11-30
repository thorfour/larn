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
	return addInfoBar(s, s.CurrentMap())
}

// infoBarGrid returns the info bar in display grid format
func infoBarGrid(s *state.State) [][]io.Runeable {
	r := make([][]io.Runeable, 2)

	info := fmt.Sprintf("Spells: %v( %v) AC: %v WC: %v Level %v Exp: %v %s", s.C.Spells, s.C.MaxSpells, s.C.Ac, s.C.Wc, s.C.Level, s.C.Exp, s.C.Title)
	for _, c := range info {
		r[0] = append(r[0], Simple(c))
	}

	info = fmt.Sprintf("HP: %v( %v) STR=%v INT=%v WIS=%v CON=%v DEX=%v CHA=%v LV: %v Gold: %v", s.C.Hp, s.C.MaxHP, s.C.Str, s.C.Intelligence, s.C.Wisdom, s.C.Con, s.C.Dex, s.C.Cha, s.C.Loc, s.C.Gold)
	for _, c := range info {
		r[1] = append(r[1], Simple(c))
	}

	return r
}

// appends the info bar to the current map
func addInfoBar(s *state.State, m [][]io.Runeable) [][]io.Runeable {

	bar := infoBarGrid(s)
	for i := range bar {
		m = append(m, bar[i])
	}
	return m
}
