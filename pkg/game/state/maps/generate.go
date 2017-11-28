package maps

import (
	"math/rand"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/io"
)

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
	if lvl == homeLevel {
		base = Empty{}
	}

	// Generate full grid
	level := make([][]io.Runeable, height)
	for i := range level {
		row := make([]io.Runeable, width)
		for j := range row {
			row[j] = base
		}
		level[i] = row
	}

	// Carve out the passageways
	if lvl != homeLevel {
		//carve(level)
		eat(Coordinate{1, 1}, level) // original larn
	}

	return level
}

// eat is the way orginal larn ate through the map of walls to create a maze
func eat(c Coordinate, lvl [][]io.Runeable) {
	dir := rand.Intn(4) + 1 // pick a random direction
	try := 2
	for try > 0 { // try all directions twice
		switch dir {
		case 1: // West
			if c.X <= 2 || lvl[c.Y][c.X-1] != (Wall{}) || lvl[c.Y][c.X-2] != (Wall{}) { // Only eat if at least the next 2 are walls
				break
			}
			lvl[c.Y][c.X-1] = Empty{}
			lvl[c.Y][c.X-2] = Empty{}
			eat(Coordinate{c.X - 2, c.Y}, lvl)
		case 2: // East
			if c.X >= width-3 || lvl[c.Y][c.X+1] != (Wall{}) || lvl[c.Y][c.X+2] != (Wall{}) { // Only eat if at least the next 2 are walls
				break
			}
			lvl[c.Y][c.X+1] = Empty{}
			lvl[c.Y][c.X+2] = Empty{}
			eat(Coordinate{c.X + 2, c.Y}, lvl)
		case 3: // South
			if c.Y <= 2 || lvl[c.Y-1][c.X] != (Wall{}) || lvl[c.Y-2][c.X] != (Wall{}) { // Only eat if at least the next 2 are walls
				break
			}
			lvl[c.Y-1][c.X] = Empty{}
			lvl[c.Y-2][c.X] = Empty{}
			eat(Coordinate{c.X, c.Y - 2}, lvl)
		case 4: // North
			if c.Y >= height-3 || lvl[c.Y+1][c.X] != (Wall{}) || lvl[c.Y+2][c.X] != (Wall{}) { // Only eat if at least the next 2 are walls
				break
			}
			lvl[c.Y+1][c.X] = Empty{}
			lvl[c.Y+2][c.X] = Empty{}
			eat(Coordinate{c.X, c.Y + 2}, lvl)
		default:
			panic("Unknown direction")
		}
		if dir++; dir > 4 { // Pick a different direction
			dir = 1
			try--
		}
	}
}

// carve takes a map that consists of walls and carves out walls to create a maze
// it is an implementation of Randomized Prim's algorithm
func carve(lvl [][]io.Runeable) {

	// Pick a random wall and add it to walls list
	walls := []Coordinate{randMapCoord()}

	// Keep carving till the walls list is empty
	for len(walls) != 0 {

		// Get random wall from walls list
		i := rand.Intn(len(walls))
		w := walls[i]

		glog.V(6).Infof("Carve loop. Wall count %v Current Wall %v\n Wall list: %v", len(walls), w, walls)

		// If that wall is adjacent to one open space (less than for initial case)
		if emptyAdjacent(w, lvl) <= 1 {
			// convert the wall to open;
			lvl[w.Y][w.X] = Empty{}
			// and add all of it's neighboring walls to the list
			walls = append(walls, adjacentWalls(w, lvl)...)
		}

		// remove it from the walls list
		walls = append(walls[0:i], walls[i+1:]...)
	}
}

// adjacentWalls returns all adjacent walls in a map
func adjacentWalls(c Coordinate, lvl [][]io.Runeable) []Coordinate {
	adj := adjacent(c)
	var cords []Coordinate
	for _, s := range adj {
		if lvl[s.Y][s.X] == (Wall{}) {
			cords = append(cords, s)
		}
	}

	return cords
}

// getAdjacent returns adjacent coordinates in a map, that are valid coordinates
func adjacent(w Coordinate) []Coordinate {
	var adj []Coordinate
	if w.X+1 < width-1 {
		adj = append(adj, Coordinate{w.X + 1, w.Y})
	}
	if w.X-1 > 0 {
		adj = append(adj, Coordinate{w.X - 1, w.Y})
	}
	if w.Y+1 < height-1 {
		adj = append(adj, Coordinate{w.X, w.Y + 1})
	}
	if w.Y-1 > 0 {
		adj = append(adj, Coordinate{w.X, w.Y - 1})
	}

	return adj
}

// emptyAdjacent returns the number of adjacent spaces that are empty
func emptyAdjacent(c Coordinate, lvl [][]io.Runeable) int {
	adj := adjacent(c)
	count := 0
	for _, s := range adj {
		if lvl[s.Y][s.X] == (Empty{}) {
			count++
		}
	}

	return count
}

func randMapCoord() Coordinate {
	x := uint(rand.Intn(width) + 1)
	y := uint(rand.Intn(height) + 1)
	return Coordinate{x, y}
}