package items

import termbox "github.com/nsf/termbox-go"

const (
	altarRune = 'A'
)

type Altar struct {
	Visibility bool
	DisplaceableItem
}

// Visible implements the visible interface
func (a *Altar) Visible(v bool) { a.Visibility = v }

// Rune implements the io.Runeable interface
func (a *Altar) Rune() rune {
	if a.Visibility {
		return altarRune
	} else {
		return invisibleRune
	}
}

// Fg implements the io.Runeable interface
func (a *Altar) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (a *Altar) Bg() termbox.Attribute { return termbox.ColorDefault }
