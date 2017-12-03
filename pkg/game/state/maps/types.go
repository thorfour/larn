package maps

import (
	termbox "github.com/nsf/termbox-go"
)

const (
	invisbleRune  = ' '
	emptyRune     = '.'
	wallRune      = '#'
	stairUpRune   = '>'
	stairDownRune = '<'
	dungeonERune  = 'E'
)

const (
	Up   = true
	Down = false
)

// Loggable returns the string to display on the log status
type Loggable interface {
	Log() string
}

// Visible is the interface to change an objects visibility
type Visible interface {
	Visible(bool)
}

// Enterable indicates a object is enterable
type Enterable interface {
	Enter() int
}

// Displaceable indicates if a character can walk on top of an object
type Displaceable interface {
	Displace() bool
}

// Coordinate is a map coordinate. (0,0) is the top left corner
type Coordinate struct {
	X uint
	Y uint
}

// Empty represents an empty map location
type Empty struct {
	visible bool
}

func (e Empty) Visible(v bool) { e.visible = v }

// Displace implementes the Displaceable interface
func (e Empty) Displace() bool { return true }

// Rune implements the io.Runeable interface
func (e Empty) Rune() rune {
	if e.visible {
		return emptyRune
	} else {
		return invisbleRune
	}
}

// Fg implements the io.Runeable interface
func (e Empty) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (e Empty) Bg() termbox.Attribute { return termbox.ColorDefault }

// Wall is a maze wall
type Wall struct {
	visible bool
}

func (w *Wall) Visible(v bool) { w.visible = v }

// Rune implements the io.Runeable interface
func (w *Wall) Rune() rune {
	if w.visible {
		return wallRune
	} else {
		return invisbleRune
	}
}

// Fg implements the io.Runeable interface
func (w *Wall) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (w *Wall) Bg() termbox.Attribute { return termbox.ColorDefault }

// Stairs is a staircase
type Stairs struct {
	up      bool // indicates if these stairs go up
	level   int  // the level the stairs lead to
	visible bool
}

func (s Stairs) Visible(v bool) { s.visible = v }

// Displace implementes the Displaceable interface
func (s Stairs) Displace() bool { return true }

// Enter implements the Enterable interface
func (s Stairs) Enter() int { return s.level }

// Rune implements the io.Runeable interface
func (s Stairs) Rune() rune {
	if s.visible {
		if s.up {
			return stairUpRune
		}
		return stairDownRune
	} else {
		return invisbleRune
	}
}

// Fg implements the io.Runeable interface
func (s Stairs) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (s Stairs) Bg() termbox.Attribute { return termbox.ColorDefault }

// DungeonEntrance home level to dungeon level 1
type DungeonEntrance struct{}

// Displace implementes the Displaceable interface
func (d DungeonEntrance) Displace() bool { return true }

// Enter implements the Enterable interface
func (d DungeonEntrance) Enter() int { return 1 } // Dungeron entrance always leads to level 1

// Rune implements the io.Runeable interface
func (d DungeonEntrance) Rune() rune { return dungeonERune }

// Fg implements the io.Runeable interface
func (d DungeonEntrance) Fg() termbox.Attribute { return termbox.ColorBlack }

// Bg implements the io.Runeable interface
func (d DungeonEntrance) Bg() termbox.Attribute { return termbox.ColorGreen }

// Log implements the Loggable interface
func (d DungeonEntrance) Log() string { return "You have found the dungeon entrnace" }
