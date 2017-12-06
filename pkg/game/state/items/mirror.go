package items

const (
	mirrorRune = 'M'
)

type Mirror struct {
	DefaultItem
}

// Rune implements the io.Runeable interface
func (m *Mirror) Rune() rune {
	if m.Visibility {
		return mirrorRune
	} else {
		return invisibleRune
	}
}

// Log implementes the Loggable interface
func (m *Mirror) Log() string {
	return "There is a mirror here"
}
