package maps

import termbox "github.com/nsf/termbox-go"

const (
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

// Coordinate is a map coordinate. (0,0) is the top left corner
type Coordinate struct {
	X uint
	Y uint
}

// Empty represents an empty map location
type Empty struct{}

// Rune implements the io.Runeable interface
func (e Empty) Rune() rune { return emptyRune }

// Fg implements the io.Runeable interface
func (e Empty) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (e Empty) Bg() termbox.Attribute { return termbox.ColorDefault }

// Wall is a maze wall
type Wall struct{}

// Rune implements the io.Runeable interface
func (w Wall) Rune() rune { return wallRune }

// Fg implements the io.Runeable interface
func (w Wall) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (w Wall) Bg() termbox.Attribute { return termbox.ColorDefault }

// Stairs is a staircase
type Stairs struct {
	up bool // indicates if these stairs go up
}

// Rune implements the io.Runeable interface
func (s Stairs) Rune() rune {
	if s.up {
		return stairUpRune
	}
	return stairDownRune
}

// Fg implements the io.Runeable interface
func (s Stairs) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (s Stairs) Bg() termbox.Attribute { return termbox.ColorDefault }

// DungeonEntrance home level to dungeon level 1
type DungeonEntrance struct{}

// Rune implements the io.Runeable interface
func (d DungeonEntrance) Rune() rune { return dungeonERune }

// Fg implements the io.Runeable interface
func (d DungeonEntrance) Fg() termbox.Attribute { return termbox.ColorBlack }

// Bg implements the io.Runeable interface
func (d DungeonEntrance) Bg() termbox.Attribute { return termbox.ColorGreen }
