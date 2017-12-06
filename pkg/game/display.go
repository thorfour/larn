package game

import (
	"fmt"

	"github.com/golang/glog"
	termbox "github.com/nsf/termbox-go"
	"github.com/thorfour/larn/pkg/game/state"
	"github.com/thorfour/larn/pkg/io"
)

const (
	logLength = 5 // Number of lines to display for the status log
)

type Simple rune

func (s Simple) Rune() rune            { return rune(s) }
func (s Simple) Fg() termbox.Attribute { return termbox.ColorDefault }
func (s Simple) Bg() termbox.Attribute { return termbox.ColorDefault }

// display returns a 2d slice representation of the game
func display(s *state.State) [][]io.Runeable {
	return cat(s.CurrentMap(), infoBarGrid(s), statusLog(s))
}

// infoBarGrid returns the info bar in display grid format
func infoBarGrid(s *state.State) [][]io.Runeable {
	r := make([][]io.Runeable, 2)

	info := fmt.Sprintf("Spells: %v( %v) AC: %v WC: %v Level %v Exp: %v %s", s.C.Stats.Spells, s.C.Stats.MaxSpells, s.C.Stats.Ac, s.C.Stats.Wc, s.C.Stats.Level, s.C.Stats.Exp, s.C.Stats.Title)
	for _, c := range info {
		r[0] = append(r[0], Simple(c))
	}

	info = fmt.Sprintf("HP: %v( %v) STR=%v INT=%v WIS=%v CON=%v DEX=%v CHA=%v LV: %v Gold: %v", s.C.Stats.Hp, s.C.Stats.MaxHP, s.C.Stats.Str, s.C.Stats.Intelligence, s.C.Stats.Wisdom, s.C.Stats.Con, s.C.Stats.Dex, s.C.Stats.Cha, s.C.Stats.Loc, s.C.Stats.Gold)
	for _, c := range info {
		r[1] = append(r[1], Simple(c))
	}

	return r
}

// overlay starts overlaying a map at the begging of the original map
func overlay(original, overlay [][]io.Runeable) [][]io.Runeable {
	if len(overlay) > len(original) {
		glog.Errorf("Overlay is longer than original %v > %v", len(overlay), len(original))
		return original
	}
	for i := range overlay {
		original[i] = overlay[i]
	}

	return original
}

// cat concatenats all the maps together
func cat(maps ...[][]io.Runeable) [][]io.Runeable {
	for i := range maps {
		if i == 0 {
			continue
		}
		maps[0] = append(maps[0], maps[i]...)
	}
	return maps[0]
}

// statusLog returns the status log that's displayed on the bottom
func statusLog(s *state.State) [][]io.Runeable {
	// Convert the status log to runes
	return convert(s.StatLog[:logLength])
}

// convert changes a slice of strings to a runeable map
func convert(s []string) [][]io.Runeable {
	r := make([][]io.Runeable, len(s))
	for i, msg := range s {
		for _, c := range msg {
			r[i] = append(r[i], Simple(c))
		}
	}

	return r
}
