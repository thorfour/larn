package maps

import "github.com/thorfour/larn/pkg/io"

const (
	height = 17
	width  = 67
)

const (
	emptyRune = '.'
)

// Empty represents an empty map location
type Empty struct{}

// Rune implements the io.runeable interface
func (e Empty) Rune() rune {
	return emptyRune
}

type Coordinate struct {
	X uint
	Y uint
}

// Maps is the collection of all the levels in the game
type Maps struct {
	volcano [][][]io.Runeable // Gross type alias?
	dungeon [][][]io.Runeable
	home    [][]io.Runeable
}

// New returns a set of maps to represent the game
func New() *Maps {
	m := new(Maps)
	m.dungeon = dungeon()
	m.volcano = volcano()
	m.home = homeLevel()
	return m
}

func homeLevel() [][]io.Runeable {
	home := make([][]io.Runeable, height)
	for i := range home {
		row := make([]io.Runeable, width)
		for j := range row {
			row[j] = Empty{}
		}
		home[i] = row
	}
	return home
}

// TODO
func dungeon() [][][]io.Runeable {
	return nil
}

// TODO
func volcano() [][][]io.Runeable {
	return nil
}

// TODO
func (m *Maps) CurrentMap() [][]io.Runeable {
	return m.home
}
