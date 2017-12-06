package items

const (
	cookieRune = 'c'
)

type Cookie struct {
	DefaultItem
	NoStats
}

func (c *Cookie) Rune() rune {
	if c.Visibility {
		return cookieRune
	} else {
		return invisibleRune
	}
}

func (c *Cookie) Log() string {
	return "You have found a fortune cookie"
}

func (c *Cookie) String() string {
	return "a fortune cookie"
}
