package items

const (
	cookieRune = 'c'
)

// Cookie implementes the item interface
type Cookie struct {
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (c *Cookie) Rune() rune {
	if c.Visibility {
		return cookieRune
	}
	return invisibleRune
}

// Log implements the displaceable interface
func (c *Cookie) Log() string {
	return "You have found a fortune cookie"
}

// String implementes the item interface
func (c *Cookie) String() string {
	return "a fortune cookie"
}
