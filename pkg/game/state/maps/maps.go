package maps

const (
	height = 17
	width  = 67
)

const (
	emptyRune = '.'
)

// Empty represents an empty map location
type Empty struct{}

// Rune implements the runeable interface
func (e Empty) Rune() rune {
	return emptyRune
}

// Runeable is an interface that can return a rune representation of itself
type Runeable interface {
	Rune() rune
}

// Maps is the collection of all the levels in the game
type Maps struct {
	volcano []Level
	dungeon []Level
	home    Level
}

// Level is a representation of a single level and all that it contains at the time
type Level [][]Runeable

// New returns a set of maps to represent the game
func New() *Maps {
	m := new(Maps)
	m.dungeon = dungeon()
	m.volcano = volcano()
	m.home = homeLevel()
	return m
}

func homeLevel() Level {
	home := make([][]Runeable, height)
	for i := range home {
		row := make([]Runeable, width)
		for j := range row {
			row[j] = Empty{}
		}
		home[i] = row
	}
	return home
}

// TODO
func dungeon() []Level {
	return nil
}

// TODO
func volcano() []Level {
	return nil
}
