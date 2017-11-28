package maps

import termbox "github.com/nsf/termbox-go"

const (
	emptyRune = '.'
	wallRune  = '#'
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
