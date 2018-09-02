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
	dungeonRune   = 'E'
	homeRune      = 'H'
	collegeRune   = 'C'
	lrsRune       = 'L'
	tradeRune     = 'P'
	bankRune      = '$'
	dndRune       = 'D'
	volRune       = 'V'
)

const (
	dungeonStr = "You have found the dungeon entrnace."
	homeStr    = "Your have found your way home."
	collegeStr = "You have found the College of Larn."
	lrsStr     = "There is an LRS office here."
	tradeStr   = "You have found the larn trading post."
	bankStr    = "You have found the bank of Larn."
	dndStr     = "There is a DND store here."
	volStr     = "You have found a volcanic shaft leading downward!"
)

// Special levels for entrances
const (
	HomeLvl    = -1
	BankLvl    = -2
	CollegeLvl = -3
	DndLvl     = -4
	TradeLvl   = -5
	LrsLvl     = -6
)

const (
	// Up is stairs going up
	Up = true
	// Down is stairs goind down
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

// Empty represents an empty map location
type Empty struct {
	visible bool
}

// Visible implements the visibility interface
func (e Empty) Visible(v bool) { e.visible = v }

// Displace implementes the Displaceable interface
func (e Empty) Displace() bool { return true }

// Rune implements the io.Runeable interface
func (e Empty) Rune() rune {
	if e.visible {
		return emptyRune
	}
	return invisbleRune
}

// Fg implements the io.Runeable interface
func (e Empty) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (e Empty) Bg() termbox.Attribute { return termbox.ColorDefault }

// Wall is a maze wall
type Wall struct {
	visible bool
}

// Visible implements the visibility interface
func (w *Wall) Visible(v bool) { w.visible = v }

// Rune implements the io.Runeable interface
func (w *Wall) Rune() rune {
	if w.visible {
		return wallRune
	}
	return invisbleRune
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

// Visible implements the visibility interface
func (s *Stairs) Visible(v bool) { s.visible = v }

// Displace implementes the Displaceable interface
func (s *Stairs) Displace() bool { return true }

// Enter implements the Enterable interface
func (s *Stairs) Enter() int { return s.level }

// Rune implements the io.Runeable interface
func (s *Stairs) Rune() rune {
	if s.visible {
		if s.up {
			return stairUpRune
		}
		return stairDownRune
	}
	return invisbleRune
}

// Fg implements the io.Runeable interface
func (s *Stairs) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (s *Stairs) Bg() termbox.Attribute { return termbox.ColorDefault }

// Entrance type are the entrances that are on the home level
type Entrance struct {
	r         rune // entrnace rune to displace
	enterCode int  // code that is returned upon entering
	log       string
}

// Displace implementes the Displaceable interface
func (e Entrance) Displace() bool { return true }

// Enter implements the Enterable interface
func (e Entrance) Enter() int { return e.enterCode }

// Rune implements the io.Runeable interface
func (e Entrance) Rune() rune { return e.r }

// Fg implements the io.Runeable interface
func (e Entrance) Fg() termbox.Attribute { return termbox.ColorBlack }

// Log implements the Loggable interface
func (e Entrance) Log() string { return e.log }

// Bg implements the io.Runeable interface
func (e Entrance) Bg() termbox.Attribute { return termbox.ColorGreen }
