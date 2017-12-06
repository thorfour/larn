package items

const (
	fountainRune = 'F'
)

type Fountain struct {
	DefaultItem
}

// Rune implements the io.Runeable interface
func (f *Fountain) Rune() rune {
	if f.Visibility {
		return fountainRune
	} else {
		return invisibleRune
	}
}

// Log implements the Displaceable interface
func (f *Fountain) Log() string {
	return "There is a Fountain here"
}
