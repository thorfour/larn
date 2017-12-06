package items

const (
	altarRune = 'A'
)

// Altar is the altar the player can pray at
type Altar struct {
	DefaultItem
}

// Rune implements the io.Runeable interface
func (a *Altar) Rune() rune {
	if a.Visibility {
		return altarRune
	}
	return invisibleRune
}

// Log implementes the Loggable interface
func (a *Altar) Log() string {
	return "There is a Holy Altar here!"
}
