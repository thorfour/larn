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
		x += runewidth.RuneWidth(c)
	}

	return termbox.Flush()
}
