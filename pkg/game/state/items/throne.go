package items

const (
	throneRune     = 'T'
	deadThroneRune = 't'
)

type Throne struct {
	DefaultItem
}

// Rune implements the io.Runeable interface
func (t *Throne) Rune() rune {
	if t.Visibility {
		return throneRune
	} else {
		return invisibleRune
	}
}

// Log implements the Displaceable interface
func (t *Throne) Log() string {
	return "There is a handsome jewel encrusted throne"
}
