package io

import (
	"github.com/golang/glog"
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

const (
	DefaultColor = termbox.ColorDefault
	ColorBlack   = termbox.ColorBlack
	ColorRed     = termbox.ColorRed
	ColorGreen   = termbox.ColorGreen
	ColorYellow  = termbox.ColorYellow
	ColorBlue    = termbox.ColorBlue
	ColorMagenta = termbox.ColorMagenta
	ColorCyan    = termbox.ColorCyan
	ColorWhite   = termbox.ColorWhite
)

type Runeable interface {
	Rune() rune
	Fg() termbox.Attribute
	Bg() termbox.Attribute
}

type Cell interface {
	Runeable
	X() int
	Y() int
}

// RenderWelcome renders a welcome string
func RenderWelcome(msg string) error {

	termbox.Clear(DefaultColor, DefaultColor)

	x, y := 0, 0
	for _, c := range msg {
		termbox.SetCell(x, y, c, DefaultColor, DefaultColor)
		switch c {
		case '\n':
			y++
			x = 0
		default:
			x += runewidth.RuneWidth(c)
		}
	}

	return termbox.Flush()
}

// RenderNew will clear the terminal and render a whole new string
func RenderNew(s string) error {

	termbox.Clear(DefaultColor, DefaultColor)

	// Render from left to right, top to bottom
	x, y := 0, 0
	for _, c := range s {
		termbox.SetCell(x, y, c, DefaultColor, DefaultColor)
		switch c {
		case '\n': // Newline; reset to next line
			y++
			x = 0
		default:
			x += runewidth.RuneWidth(c)
		}
	}

	return termbox.Flush()
}

// RenderNewGrid clears the terminal before rendering the grid
func RenderNewGrid(grid [][]Runeable) error {
	termbox.Clear(DefaultColor, DefaultColor)
	return RenderGrid(grid)
}

// RenderGrid renders a 2d grid. where (0,0) is in the top left
// each slice is treated as a row. This is a wrapper around RenderGridOffset
func RenderGrid(grid [][]Runeable) error {
	return RenderGridOffset(0, 0, grid)
}

// RenderGridOffset renders a 2d grid starting at (x,y)
func RenderGridOffset(x, y int, grid [][]Runeable) error {
	xo := x // Set the x offset
	for _, row := range grid {
		for _, c := range row {
			termbox.SetCell(xo, y, c.Rune(), c.Fg(), c.Bg())
			xo += runewidth.RuneWidth(c.Rune())
		}
		xo = x // reset xoffest for next row
		y++    // go to next line
	}
	return termbox.Flush()
}

// RenderNewGridOffset clears the screen before calling render grid offset
func RenderNewGridOffset(x, y int, grid [][]Runeable) error {
	termbox.Clear(DefaultColor, DefaultColor)
	return RenderGridOffset(x, y, grid)
}

// RenderCell renders a sincgle cell
func RenderCell(x, y int, c rune, fg, bg termbox.Attribute) error {
	termbox.SetCell(x, y, c, fg, bg)
	return termbox.Flush()
}

func RenderCells(c []Cell) error {
	for _, ci := range c {
		glog.V(6).Infof("RenderCells: (%v,%v); Char(%v); %v: %v", ci.X(), ci.Y(), ci.Rune(), ci.Fg(), ci.Bg())
		termbox.SetCell(ci.X(), ci.Y(), ci.Rune(), ci.Fg(), ci.Bg())
	}
	return termbox.Flush()
}
