package items

import termbox "github.com/nsf/termbox-go"

const (
	bookRune = 'B'
)

type Book struct {
	Level   uint // the level the book was placed in
	visible bool
}

// Visible implementes the visible interface
func (b Book) Visible(v bool) { b.visible = v }

// Rune implements the io.Runeable interface
func (b Book) Rune() rune {
	if b.visible {
		return bookRune
	} else {
		return invisibleRune
	}
}

// Fg implements the io.Runeable interface
func (b Book) Fg() termbox.Attribute { return termbox.ColorDefault }

// Bg implements the io.Runeable interface
func (b Book) Bg() termbox.Attribute { return termbox.ColorDefault }
