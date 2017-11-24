package io

import (
	"github.com/golang/glog"
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
)

func RenderWelcome(msg string) error {

	termbox.Clear(defaultColor, defaultColor)
	w, h := termbox.Size()

	glog.V(4).Info("current width %v height: %v", w, h)

	x := 0
	y := 0
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

func RenderBorder(ch rune, width, height int) error {

	termbox.Clear(defaultColor, defaultColor)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i == 0 || j == 0 || i == width-1 || j == height-1 {
				termbox.SetCell(i, j, ch, defaultColor, defaultColor)
			}
		}
	}

	return termbox.Flush()
}
