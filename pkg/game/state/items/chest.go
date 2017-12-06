package items

const (
	chestRune = 'C'
)

type Chest struct {
	Level uint
	DefaultItem
	NoStats
}

func (c *Chest) Rune() rune {
	if c.Visibility {
		return chestRune
	} else {
		return invisibleRune
	}
}

func (c *Chest) Log() string {
	return "There is a chest here"
}

func (c *Chest) String() string {
	return "a chest"
}
