package items

const (
	bookRune = 'B'
)

type Book struct {
	Level      uint // the level the book was placed in
	Visibility bool
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
