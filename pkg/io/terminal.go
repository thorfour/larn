package io

import (
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
)

// RenderWelcome renders a welcome string
func RenderWelcome(msg string) error {

	termbox.Clear(defaultColor, defaultColor)

	x, y := 0, 0
	for _, c := range msg {
		termbox.SetCell(x, y, c, defaultColor, defaultColor)
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

	termbox.Clear(defaultColor, defaultColor)

	// Render from left to right, top to bottom
	x, y := 0, 0
	for _, c := range s {
		termbox.SetCell(x, y, c, defaultColor, defaultColor)
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
func RenderNewGrid(grid [][]rune) error {
	termbox.Clear(defaultColor, defaultColor)
	return RenderGrid(grid)
}

// RenderGrid renders a 2d grid. where (0,0) is in the top left
// each slice is treated as a row. This is a wrapper around RenderGridOffset
func RenderGrid(grid [][]rune) error {
	return RenderGridOffset(0, 0, grid)
}

// RenderGridOffset renders a 2d grid starting at (x,y)
func RenderGridOffset(x, y int, grid [][]rune) error {
	xo := x // Set the x offset
	for _, row := range grid {
		for _, c := range row {
			termbox.SetCell(xo, y, c, defaultColor, defaultColor)
			xo += runewidth.RuneWidth(c)
		}
		xo = x // reset xoffest for next row
		y++    // go to next line
	}
	return termbox.Flush()
}
