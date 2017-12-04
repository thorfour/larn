package items

import termbox "github.com/nsf/termbox-go"

const (
	altarRune = 'A'
)

type Altar struct {
	visible bool
}

// Visible implements the visible interface
func (a Altar) Visible(v bool) { a.visible = v }

// Rune implements the io.Runeable interface
func (a Altar) Rune() rune {
	if a.visible {
		return altarRune
	} else {
		return invisibleRune
	}
}

// Fg implements the io.Runeable interface
func (a Altar) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (a Altar) Bg() termbox.Attribute { return termbox.ColorDefault }
