package items

import "github.com/thorfour/larn/pkg/game/state/stats"

const (
	bookRune = 'B'
)

type Book struct {
	Level uint // the level the book was placed in
	DefaultItem
}

// Rune implements the io.Runeable interface
func (b *Book) Rune() rune {
	if b.Visibility {
		return bookRune
	} else {
		return invisibleRune
	}
}

// Log implements the Loggable interface
func (b *Book) Log() string {
	return "You have found a book"
}

// PickUp implements the Item interface
func (b *Book) PickUp(s *stats.Stats) {}

// Drop implements the Item interface
func (b *Book) Drop(s *stats.Stats) {}

// String returns the texutal representation of the item
func (b *Book) String() string { return "a book" }
