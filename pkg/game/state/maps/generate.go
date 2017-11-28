package maps

import "github.com/thorfour/larn/pkg/io"

const (
	homeLevel  = 0
	maxDungeon = 10
	maxVolcano = 13
)

// newLevel creates a new map for a given level
// It creates a map full of walls and then carves out the pathways
// If level == 0 it returns an empty map for the home level
func newLevel(lvl uint) [][]io.Runeable {

	var base io.Runeable
	base = Wall{}
	if lvl == 0 {
		base = Empty{}
	}

	// Generate full grid
	home := make([][]io.Runeable, height)
	for i := range home {
		row := make([]io.Runeable, width)
		for j := range row {
			row[j] = base
		}
		home[i] = row
	}

	return home
}
