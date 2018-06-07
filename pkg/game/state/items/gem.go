package items

import "math/rand"

const (
	gemRune = '*'
)

const (
	// Diamond gemstone
	Diamond = iota
	// Ruby gemstone
	Ruby
	// Emerald gemstone
	Emerald
	// Sapphire gemstone
	Sapphire
)

// Gem represents a gemstone
type Gem struct {
	Stone int // indicates the type of gemstone
	Value int // the value of the gemstone
	DefaultItem
	NoStats
}

// Rune implements the io.Runeable interface
func (g *Gem) Rune() rune {
	if g.Visibility {
		return gemRune
	}
	return invisibleRune
}

// Log is the log message for when a user walks over a gemstone
func (g *Gem) Log() string {
	return "You have found " + g.String()
}

// String is the printable description of the gemstone
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

// CreateGem returns a new gemstone
func CreateGem() *Gem {
	// TODO the value is currently the quality, but sale value shoudl be calculated instead
	switch rand.Intn(3) {
	case 0:
		return &Gem{Stone: Diamond, Value: (rand.Intn(50) + 51) / 10}
	case 1:
		return &Gem{Stone: Ruby, Value: (rand.Intn(40) + 41) / 10}
	case 2:
		return &Gem{Stone: Ruby, Value: (rand.Intn(30) + 31) / 10}
	default:
		return &Gem{Stone: Ruby, Value: (rand.Intn(20) + 21) / 10}
	}
}
