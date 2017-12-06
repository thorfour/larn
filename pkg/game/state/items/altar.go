package items

const (
	altarRune = 'A'
)

type Altar struct {
	Visibility bool
	DefaultItem
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

// Log implementes the Loggable interface
func (a *Altar) Log() string {
	return "There is an altar here"
}
