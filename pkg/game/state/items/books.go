package items

const (
	bookRune = 'B'
)

type Book struct {
	Level      uint // the level the book was placed in
	Visibility bool
	DisplaceableItem
	DefaultItem
}

// Visible implementes the visible interface
func (b *Book) Visible(v bool) { b.Visibility = v }

// Rune implements the io.Runeable interface
func (b *Book) Rune() rune {
	if b.Visibility {
		return bookRune
	} else {
		return invisibleRune
	}
}
