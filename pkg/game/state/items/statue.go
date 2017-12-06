package items

const (
	statueRune = '&'
)

type Statue struct {
	DefaultItem
}

// Rune implements the io.Runeable interface
func (s *Statue) Rune() rune {
	if s.Visibility {
		return statueRune
	} else {
		return invisibleRune
	}
}

// Log implements the Displaceable interface
func (s *Statue) Log() string {
	return "You are standing in front of a statue"
}
