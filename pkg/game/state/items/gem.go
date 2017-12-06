package items

const (
	gemRune = '*'
)

const (
	Diamond = iota
	Ruby
	Emerald
	Sapphire
)

type Gem struct {
	Stone int // indicates the type of gemstone
	Value int // the value of the gemstone
	DefaultItem
	NoStats
}

func (g *Gem) Rune() rune {
	if g.Visibility {
		return gemRune
	} else {
		return invisibleRune
	}
}

func (g *Gem) Log() string {
	return "You have found " + g.String()
}

func (g *Gem) String() string {
	switch g.Stone {
	case Diamond:
		return "a brilliant diamond"
	case Ruby:
		return "a ruby"
	case Emerald:
		return "an enchanting emerald"
	case Sapphire:
		return "a sparkling sapphire"
	default:
		return "a sparkling sapphire"
	}
}
