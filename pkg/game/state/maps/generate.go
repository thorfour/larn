package maps

import (
	"math/rand"
	"time"

	"github.com/golang/glog"
	"github.com/thorfour/larn/pkg/game/state/items"
	"github.com/thorfour/larn/pkg/io"
)

const (
	homeLevel  = 0
	maxDungeon = 10
	maxVolcano = 13 // 3 volcanos. 10 dungeons
)

// newMap is a wrapper of newLevel, it creates the level and places objects in the level.
func newMap(lvl uint) [][]io.Runeable {
	m := newLevel(lvl) // Create the level

	seed := time.Now().UnixNano()
	glog.V(1).Infof("Map Seed: %v", seed)

	// Seed the global rand // FIXME maps shouldn't use global rand
	rand.Seed(seed)

	placeObjects(lvl, m) // Add objects to the level
	return m
}

// newLevel creates a new map for a given level
// It creates a map full of walls and then carves out the pathways
// If level == 0 it returns an empty map for the home level
func newLevel(lvl uint) [][]io.Runeable {

	base := func() io.Runeable {
		return &Wall{DEBUG} // If DEBUG is set, maze will be visible by default
	}
	if lvl == homeLevel {
		base = func() io.Runeable { return Empty{true} }
	}

	// Generate full grid
	level := make([][]io.Runeable, height)
	for i := range level {
		row := make([]io.Runeable, width)
		for j := range row {
			row[j] = base()
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
	dir := rand.Intn(4) + 1  // pick a random direction // TODO THOR use a time seed
	for try := 2; try > 0; { // try all directions twice
		switch dir {
		case 1: // West
			if c.X <= 2 || !isWall(Coordinate{c.X - 1, c.Y}, lvl) || !isWall(Coordinate{c.X - 2, c.Y}, lvl) { // Only eat if at least the next 2 are walls
				break
			}
			lvl[c.Y][c.X-1] = Empty{}
			lvl[c.Y][c.X-2] = Empty{}
			eat(Coordinate{c.X - 2, c.Y}, lvl)
		case 2: // East
			if c.X >= width-3 || !isWall(Coordinate{c.X + 1, c.Y}, lvl) || !isWall(Coordinate{c.X + 2, c.Y}, lvl) { // Only eat if at least the next 2 are walls
				break
			}

			lvl[c.Y][c.X+1] = Empty{}
			lvl[c.Y][c.X+2] = Empty{}
			eat(Coordinate{c.X + 2, c.Y}, lvl)
		case 3: // South
			if c.Y <= 2 || !isWall(Coordinate{c.X, c.Y - 1}, lvl) || !isWall(Coordinate{c.X, c.Y - 2}, lvl) { // Only eat if at least the next 2 are walls
				break
			}
			lvl[c.Y-1][c.X] = Empty{}
			lvl[c.Y-2][c.X] = Empty{}
			eat(Coordinate{c.X, c.Y - 2}, lvl)
		case 4: // North
			if c.Y >= height-3 || !isWall(Coordinate{c.X, c.Y + 1}, lvl) || !isWall(Coordinate{c.X, c.Y + 2}, lvl) { // Only eat if at least the next 2 are walls
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
	adj := adjacent(c, true)
	var cords []Coordinate
	for _, s := range adj {
		switch lvl[s.Y][s.X].(type) {
		case *Wall:
			cords = append(cords, s)
		}
	}

	return cords
}

// getAdjacent returns adjacent coordinates in a map, that are valid coordinates
// mapEdge indicates to include the edge of the map as well
func adjacent(w Coordinate, mapEdge bool) []Coordinate {
	var adj []Coordinate
	min := 0
	maxW := width - 1
	maxH := height - 1
	if mapEdge {
		min = 1
		maxW = width - 2
		maxH = height - 2
	}
	if int(w.X)+1 <= maxW {
		adj = append(adj, Coordinate{w.X + 1, w.Y})
	}
	if int(w.X)-1 >= min {
		adj = append(adj, Coordinate{w.X - 1, w.Y})
	}
	if int(w.Y)+1 <= maxH {
		adj = append(adj, Coordinate{w.X, w.Y + 1})
	}
	if int(w.Y)-1 >= min {
		adj = append(adj, Coordinate{w.X, w.Y - 1})
	}

	return adj
}

// diagonal returns the diagonally adjacent coordinates in a map that are valid coordinates
// mapEdge indicates to include the edge of the map as well
func diagonal(w Coordinate, mapEdge bool) []Coordinate {
	var adj []Coordinate
	min := 0
	maxW := width - 1
	maxH := height - 1
	if mapEdge {
		min = 1
		maxW = width - 2
		maxH = height - 2
	}
	if int(w.X)+1 <= maxW && int(w.Y)+1 <= maxH {
		adj = append(adj, Coordinate{w.X + 1, w.Y + 1})
	}
	if int(w.X)-1 >= min && int(w.Y)-1 >= min {
		adj = append(adj, Coordinate{w.X - 1, w.Y - 1})
	}
	if int(w.Y)+1 <= maxH && int(w.X)-1 >= min {
		adj = append(adj, Coordinate{w.X - 1, w.Y + 1})
	}
	if int(w.Y)-1 >= min && int(w.X)+1 <= maxW {
		adj = append(adj, Coordinate{w.X + 1, w.Y - 1})
	}

	return adj
}

// emptyAdjacent returns the number of adjacent spaces that are empty
func emptyAdjacent(c Coordinate, lvl [][]io.Runeable) int {
	adj := adjacent(c, true)
	count := 0
	for _, s := range adj {
		switch lvl[s.Y][s.X].(type) {
		case Empty:
			count++
		}
	}

	return count
}

func randMapCoord() Coordinate {
	x := uint(rand.Intn(width-1) + 1)
	y := uint(rand.Intn(height-1) + 1)
	return Coordinate{x, y}
}

// walkToEmpty takes coordincate c and randomly walks till it finds an empty location
func walkToEmpty(c Coordinate, lvl [][]io.Runeable) Coordinate {

	// Random walk till and empty room is found
	for {
		switch lvl[c.Y][c.X].(type) {
		case Empty:
			return c
		}
		xadj := uint(rand.Intn(3) - 2) // [-1,1]
		yadj := uint(rand.Intn(3) - 2) // [-1,1]
		if xadj > 0 {
			c.X += xadj
		} else {
			c.X -= xadj
		}

		if yadj > 0 {
			c.Y += yadj
		} else {
			c.Y -= yadj
		}

		if c.X > width-2 {
			c.X = 1
		}
		if c.X < 1 {
			c.X = width - 2
		}
		if c.Y > height-2 {
			c.Y = 1
		}
		if c.Y < 1 {
			c.Y = height - 2
		}
	}

	return c
}

// placeMultipleObjects places [0,N) objects of type o in lvl at random coordinates
func placeMultipleObjects(n int, f func() io.Runeable, lvl [][]io.Runeable) {
	for i := 0; i < rand.Intn(n); i++ {
		// generate the object
		o := f()
		// Set the visibility of the object
		o.(Visible).Visible(DEBUG)

		// Place the object in the map
		placeObject(randMapCoord(), o, lvl)
	}
}

// placeObject places an object in a maze at arandom open location
func placeObject(c Coordinate, o io.Runeable, lvl [][]io.Runeable) (Coordinate, io.Runeable) {

	glog.V(6).Infof("Placing %s", string(o.Rune()))

	c = walkToEmpty(c, lvl)

	// Add the object
	d := lvl[c.Y][c.X]
	lvl[c.Y][c.X] = o

	return c, d
}

// placeObjects places the required objects for a level
// it calls placeObject many times
func placeObjects(lvl uint, m [][]io.Runeable) {

	// Place the stairs
	if lvl == homeLevel {
		placeObject(randMapCoord(), Entrance{dungeonRune, 1, dungeonStr}, m)          // Entrance to the dungeon
		placeObject(randMapCoord(), Entrance{homeRune, homeLvl, homeStr}, m)          // players home
		placeObject(randMapCoord(), Entrance{collegeRune, collegeLvl, collegeStr}, m) // college of larn
		placeObject(randMapCoord(), Entrance{bankRune, bankLvl, bankStr}, m)          // 1st national bank of larn
		placeObject(randMapCoord(), Entrance{volRune, maxDungeon + 1, volStr}, m)     // volano shaft
		placeObject(randMapCoord(), Entrance{dndRune, dndLvl, dndStr}, m)             // the DND store
		placeObject(randMapCoord(), Entrance{tradeRune, tradeLvl, tradeStr}, m)       // the trading post
		placeObject(randMapCoord(), Entrance{lrsRune, lrsLvl, lrsStr}, m)             // the larn revenue service
	} else {
		// Place the stairs
		if lvl != 1 { // Dungeon level 1 has an entrance/exit doesn't need stairs up
			placeObject(randMapCoord(), Stairs{Up, int(lvl - 1), DEBUG}, m)
		}
		if lvl != maxDungeon && lvl != maxVolcano { // Last dungeon/volcano no stairs down
			placeObject(randMapCoord(), Stairs{Down, int(lvl + 1), DEBUG}, m)
		}

		// Place random maze objects
		placeMultipleObjects(3, func() io.Runeable { return &items.Book{Level: lvl} }, m) // up to 2 books a level
		placeMultipleObjects(3, func() io.Runeable { return new(items.Altar) }, m)        // up to 2 altars a level
	}
}

// isWall returns true if the coordinate c is a wall on map lvl
func isWall(c Coordinate, lvl [][]io.Runeable) bool {
	switch lvl[c.Y][c.X].(type) {
	case *Wall:
		return true
	default:
		return false
	}
}
