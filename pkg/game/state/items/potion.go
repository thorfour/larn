package items

const (
	potionRune = '!'
)

type Potion struct {
	DefaultItem
	NoStats
}

// rune implements the io.Runeable interface
func (p *Potion) Rune() rune {
	if p.Visibility {
		return potionRune
	} else {
		return invisibleRune
	}
}

// Log implements the Disaplceable interface
func (p *Potion) Log() string {
	return "You have found a magic potion"
}

// String implements the Item interface
func (p *Potion) String() string {
	return "a potion"
}
