package items

const (
	scrollRune = '?'
)

type Scroll struct {
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (s *Scroll) Rune() rune {
	if s.Visibility {
		return scrollRune
	} else {
		return invisibleRune
	}
}

// Log implements the Disaplceable interface
func (s *Scroll) Log() string {
	return "You have found a magic scroll"
}

// String implements the Item interface
func (s *Scroll) String() string {
	return "a scroll"
}
