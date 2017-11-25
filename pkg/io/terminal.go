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
