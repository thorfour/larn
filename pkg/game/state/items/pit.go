package items

const (
	pitRune = 'P'
)

type Pit struct {
	DefaultItem
}

// Rune implements the io.Runeable interface
func (p *Pit) Rune() rune {
	if p.Visibility {
		return pitRune
	} else {
		return invisibleRune
	}
}

// Log implements the Displaceable interface
func (p *Pit) Log() string {
	return "You are standing at the top of a pit"
}
